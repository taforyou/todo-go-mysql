import React, { Component } from "react";
import "./App.css";
import { Input, Icon, Button, Card, List, Checkbox } from "antd";

// Step 3 - Design component and state

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      inputText: "",
      todos: [
        {
          id: "1",
          title: "Title1",
          completed: true,
          selected: true
        },
        {
          id: "2",
          title: "Title2",
          completed: false,
          selected: false
        }
      ],
      isLoading: false
    };
  }

  render() {
    return (
      <div className="App">
        <Card>
          <h1>To-do-list</h1>

          <div style={{ marginBottom: "10px" }}>
            <Input.Search
              placeholder="input search text"
              enterButton="Add"
              size="large"
            />
          </div>

          <List
            bordered
            dataSource={this.state.todos}
            style={{ height: 300, overflow: "auto" }}
            loading={this.state.isLoading}
            header={
              <div style={{ position: "relative", height: 20 }}>
                <Checkbox
                  style={{ marginRight: 10, position: "absolute", left: 0 }}
                >
                  Select All
                </Checkbox>
                <Button.Group style={{ position: "absolute", right: 0 }}>
                  <Button type="dashed" size="small" style={{ marginRight: 5 }}>
                    All
                  </Button>
                  <Button type="dashed" size="small" style={{ marginRight: 5 }}>
                    Active
                  </Button>
                  <Button
                    type="danger"
                    size="small"
                    style={{
                      display: this.state.isLoading ? "none" : " ",
                      marginLeft: 10
                    }}
                  >
                    Clear completed
                  </Button>
                </Button.Group>
              </div>
            }
            renderItem={todo => (
              <List.Item
                actions={[
                  <Icon
                    type="close-circle"
                    style={{ fontSize: 16, color: "rgb(255, 145, 0)" }}
                  />
                ]}
                style={{
                  textDecorationLine: todo.completed ? "line-through" : "none"
                }}
              >
                <Checkbox checked={todo.selected} style={{ marginRight: 10 }} />
                <h4>{todo.title}</h4>
              </List.Item>
            )}
          />
        </Card>
      </div>
    );
  }
}

export default App;
