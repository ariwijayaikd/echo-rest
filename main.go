package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func getlistUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func createnewUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	roles := c.FormValue("roles")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email + ", roles:" + roles)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/user", createnewUser)
	e.GET("/user/:id", getlistUser)
	e.GET("/user", getdetailUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	// Create REST API:
	// Create new user: POST /user/
	// Update user: PUT /user/{id}
	// Delete user: DELETE /user/{id}
	// Get list users: GET /user
	// Get Detail user: GET /user/{id}

	e.Logger.Fatal(e.Start(":1323"))
}