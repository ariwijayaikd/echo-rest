package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"strconv"
	// "log"

	"github.com/labstack/echo/v4"
	// _ "github.com/go-sql-driver/mysql"
)

// var db *gorm.DB
// var err error

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Emails   string `json:"email"`
	// Roles    int `json:"role"`
}

type TipeRole struct {
	Roles[]Role `json:"role"`
}

type Users []User

var users = Users{User{generateId,"Ari Wijaya","@gmail.com"/*,generateRole*/}}

func generateId() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}

// func generateRole() int {
//
//
// }
//lanjut disini

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
	user.Id = generateId()
	users = append(users, user)
	return c.JSON(http.StatusCreated, users)
}

func detailUser(c echo.Context) error {
	id, _ := strconv.Atot(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atot(c.Param("id"))
	for i, _ := range users {
		if users[i].Id == id {
			users[i].Online = false
			return c.JSON(http.StatusOK, users)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atot(c.Param("id"))
	for i, _ := range users {
		if users[i].Id == id {
			users = append(users[:i],users[i+1]...)
			return c.JSON(http.StatusOK, users)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
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
	e.GET("/user/", listUser)
	e.GET("/user/:id", detailUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
