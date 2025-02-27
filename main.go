package main

import (
	"github.com/gin-gonic/gin" // Web framework for handling HTTP requests
	"net/http" // Standard package in Go for sending responses
)

// Task struct represents a task
// (similar to a class in other languages)
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// In-memory task list
// dynamic array that stores tasks
// each task is represented as a struct instance
var tasks = []Task{
	{ID: "1", Title: "Learn Go", Status: "pending"},
	{ID: "2", Title: "Build a REST API", Status: "in progress"},
}


// Entry Point
func main() {
	router := gin.Default() // Creates a new Gin router with default middleware (logging & recovery)

	// Define API endpoints
	router.GET("/tasks", getTasks)          // Get all tasks
	router.GET("/tasks/:id", getTaskByID)   // Get a specific task by ID
	router.POST("/tasks", createTask)       // Create a new task
	router.PUT("/tasks/:id", updateTask)    // Update a task by ID
	router.DELETE("/tasks/:id", deleteTask) // Delete a task by ID

	router.Run(":8080") // Start the web server on port 8080
}

// -------- CRUD Operations ------------------

// Get all tasks
// curl http://localhost:8080/tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// Get a single task by ID
// curl http://localhost:8080/tasks/1
func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Create a new task
//curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"id": "3", "title": "Learn cURL", "status": "pending"}'
// -H: This tells the server that the request body contains JSON data.
// -d: This sends a JSON object containing task details in the request body.
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// Update a task
// curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"id": "1", "title": "Master Go", "status": "completed"}'
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Delete a task
// curl -X DELETE http://localhost:8080/tasks/2
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
