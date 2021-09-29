// Create New Simple CRUD
// Entity details:
// Table user:
// id: int
// username string
// email string
// roles: int

// Create REST API:
// Create new user: POST /user/
// Update user: PUT /user/{id}
// Delete user: DELETE /user/{id}
// Get list users: GET /user
// Get Detail user: GET /user/{id}

package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func listUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func newUser(c echo.Context) error {
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

	e.POST("/user", newUser)
	e.GET("/user/:id", listUser)
	e.GET("/user", detailUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
