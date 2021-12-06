package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo represents data about a todo list.
type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

// todolist are slice to seed todo data.
var todo_list = []todo{
	{ID: "1", Title: "Cat1", Task: "do serialization in rust", Status: false},
	{ID: "2", Title: "Cat2", Task: "implement NFT", Status: false},
	{ID: "3", Title: "cat1", Task: "Task 3", Status: false},
}

func main() {
	router := gin.Default()
	router.GET("/todolists", getTodoList)
	router.POST("/todolists", postTask)
	router.GET("/todolists/:id", getTaskByID)

	router.Run("localhost:8000")
}

// gettodolist responds with the list of all task as JSON.
func getTodoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo_list)
}

//posttask adds an task from JSON received in the request body
func postTask(c *gin.Context) {
	var newTodo todo

	//call bindJSON to bind the received JSON to
	//new newTodo.

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	//Add the new album to the slice
	todo_list = append(todo_list, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

//getTaskByID locates the task whose ID value matches the id
// parameter send by the client, then returns that task as a reponse.
func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	//loop over the list of todolist. looking for an task whose ID value matches the parameter.
	for _, a := range todo_list {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
