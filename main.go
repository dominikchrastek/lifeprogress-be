package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lifeprogress/api"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func main() {
	app := gin.Default()

	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	var c config
	json.Unmarshal(raw, &c)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port,
		c.User, c.Password, c.DbName)
	dbc, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbc.Close()

	// settings := cors.Config{
	// 	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	AllowAllOrigins: true,
	// }

	app.Use(cors.Default())

	api.Register(app, dbc)

	app.Run(":3000")
}
