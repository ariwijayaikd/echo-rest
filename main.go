package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"net/http"
	// "strconv"
	// "log"

  // "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/labstack/echo/v4"
	// _ "github.com/go-sql-driver/mysql"
)

// var db *gorm.DB
// var err error

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Emails   string `"json:email"`
	Roles    int `json:"role"`
}

type Users []User

var users Users

// func generateId() int {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	return r.Intn(10000)
// }

func listUser(c echo.Context) error {
	// User ID from path `users/:id`
	return c.JSON(http.StatusOK, users)
}

func newUser(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	users = append(users, user)
	return c.JSON(http.StatusCreated, users)
}

func main() {
	fmt.Println("Running...")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, welcome to echo!")
	})

	// db, err := gorm.Open("mysql", "root:Gebangsari1@/go_rest_api_crud?charset=utf8&parseTime=True")
	// if err != nil {
	// 	log.Println("Connection failed", err)
	// } else {
	// 	log.Println("Connection enstabilshed")
	// }
	//
	// db.AutoMigrate(&Product{})
	// handleRequests()

	e.POST("/user", newUser)
	e.GET("/user/:id", listUser)
	// e.GET("/user", detailUser)
	// e.PUT("/user/:id", updateUser)
	// e.DELETE("/user/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
