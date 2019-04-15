package main

import (
	"fmt"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	TITLE     string `json:"title"`
	COMPLETED bool   `json:"completed"`
	SELECTED  bool   `json:"selected"`
}

type IDIndex struct {
	ID int `json:"id"`
}

type NewTodo struct {
	Title string `json:"title"`
}

// Step-8. [Frontend + Backend — Create] Add endpoint Fetch API

func main() {

	// fmt.Println(todos)
	// fmt.Println(todos[0].TITLE)
	// fmt.Println(todos[1].TITLE)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CROS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// API Endpoint
	e.GET("/fetchTodos", fetchTodos)

	e.POST("/addTodos", addTodos)

	e.Logger.Fatal(e.Start(":1323"))

}

func addTodos(c echo.Context) (err error) {
	u := new(Todo)

	// ถ้า User เพิ่ง Add เค้ามาเรายังจะไม่รู้ ID เพราะฉะนั้นจึงต้องไป Fetch เพื่อเอา ID ล่าสุดของ ตารางมาก่อน
	var idIndexs []IDIndex

	db, err := sql.Open("mysql", "yourusername:yourpassword@/testing")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("OK Connect DB - for Add new todo")
	results, err := db.Query("SELECT id FROM todos ORDER BY id DESC LIMIT 1")

	// ปิดแบบที 1
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	// เหมือนจะใช้ Over ไปหน่อยนะ เพราะจริงๆ มันมีแค่ ID เดียวเอง
	for results.Next() {
		var idIndex IDIndex

		err = results.Scan(&idIndex.ID)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		idIndexs = append(idIndexs, idIndex)
	}

	// เอา ID ตัวสุดท้ายมาแล้วเพิ่มไปอีก 1
	//fmt.Println(idIndexs[0].ID + 1)
	// Check ให้ชัวร์ว่า Size = 1 จริงๆ
	//fmt.Println(len(idIndexs))

	// Assign ID ใหม่ให้กับ USER
	u.ID = idIndexs[0].ID + 1

	// เอา content ที่ถูกส่งมาจาก JSON มา bind เข้ากับ i ที่ประกาศมาใหม่
	i := new(NewTodo)

	if err = c.Bind(i); err != nil {
		return
	}

	// สุดท้ายเอารายละเอียดใน Title มายัดเข้าไปใน u ที่ประกาศขึ้นมา
	u.TITLE = i.Title
	// สร้าง todo ใหม่ผมคิดว่า ว่ามันต้องยังไม่ completed และ ก็ยังไม่ถูกเลือก จึงบังคับให้เป็น false
	u.COMPLETED = false
	u.SELECTED = false

	results2, err := db.Query("INSERT INTO `testing`.`todos` (`id`, `title`, `completed`, `selected`) VALUES (?, ?, '0', '0')", u.ID, u.TITLE)

	if err != nil {
		panic(err.Error())
	}

	// ปิดแบบนี้ก็ได้เนอะ แปลก
	defer results2.Close()

	return c.JSON(http.StatusOK, u)
}

func fetchTodos(c echo.Context) error {
	var todos []Todo

	db, err := sql.Open("mysql", "yourusername:yourpassword@/testing")

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("OK Connect DB - for fetch list todo")
	results, err := db.Query("SELECT * FROM todos")

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	for results.Next() {
		var todo Todo

		err = results.Scan(&todo.ID, &todo.TITLE, &todo.COMPLETED, &todo.SELECTED)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		todos = append(todos, todo)

	}

	return c.JSON(http.StatusOK, todos)
}
