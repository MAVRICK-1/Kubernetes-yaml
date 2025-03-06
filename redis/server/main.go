package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin framework
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET endpoint at "/ping"
	r.GET("/ping", func(c *gin.Context) {
		// Respond with a JSON object containing a message
		c.JSON(200, gin.H{
			"message": "diya nigaa",
		})
	})

	// Define another GET endpoint at "/fuck"
	r.GET("/fuck", func(c *gin.Context) {
		// First JSON response (this will be sent successfully)
		c.JSON(200, gin.H{
			"hi": "byee",
		})

		// Second JSON response (this will NOT be sent because the response is already written)
		c.JSON(400, gin.H{
			"something wnet": "ein",
		})
		// ❌ ERROR: In Gin (or any HTTP server), you can only send **one response** per request.
		// The second `c.JSON(400, gin.H{...})` is ignored because the first response was already sent.
	})

	// Define a POST endpoint at "/post"
	r.POST("/post", func(ctx *gin.Context) {
		var user any // Declare a variable `user` of type `any` (which can store any type of data)

		// Bind JSON request body to the `user` variable
		if err := ctx.ShouldBindJSON(&user); err != nil {
			// If there's an error (e.g., invalid JSON), return HTTP 400 with the error message
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return // Stop further execution
		}

		// Print the user data to the console (for debugging)
		fmt.Printf("User data: %+v\n", user)

		// Send a success response with the received JSON data
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"data":    user,
		})
	})

	// Start the HTTP server on port 6060
	r.Run(":6060") // This will run the server at http://localhost:6060
}
