package controllers

import (
	"todo-mvc/todo"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/websocket"
)

type TodoController struct {
	Service todo.Service

	Session *sessions.Session

	NS *websocket.NSConn
}

func (c *TodoController) BeforeActivation(b mvc.BeforeActivation) {

	b.Dependencies().Register(func(ctx iris.Context) (items []todo.Item) {
		ctx.ReadJSON(&items)
		return
	})
}

func (c *TodoController) Get() []todo.Item {
	return c.Service.Get(c.Session.ID())
}

type PostItemResponse struct {
	Success bool `json:"success"`
}

var emptyResponse = PostItemResponse{Success: false}

func (c *TodoController) Post(newItems []todo.Item) PostItemResponse {
	if err := c.Service.Save(c.Session.ID(), newItems); err != nil {
		return emptyResponse
	}

	return PostItemResponse{Success: true}
}

func (c *TodoController) Save(msg websocket.Message) error {
	id := c.Session.ID()
	c.NS.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "saved",
		To:        id,
		Body:      websocket.Marshal(c.Service.Get(id)),
	})

	return nil
}
