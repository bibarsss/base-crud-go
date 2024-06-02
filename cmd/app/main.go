package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Title: "Clean Room", Completed: false},
	{ID: "2", Title: "Read book", Completed: false},
	{ID: "3", Title: "Record Video", Completed: false},
}

func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "This is the index")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
