package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	router := gin.Default()
	router.GET("/exec", runHandler)
	err := router.Run(":4000")
	if err != nil {
		return
	}
}

func runHandler(c *gin.Context) {
	cmd := exec.Command("duf", "-all")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.String(http.StatusOK, string(output))
}
