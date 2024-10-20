package main

import (
	"encoding/json"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []User{{1, "juan"}, {2, "pedro"}, {3, "pablo"}}

func main() {
	server := Server{Port: 3000}
	server.GET("/", func(req Request, res *Response) {
		res.Json(users)
	})

	server.POST("/new", func(req Request, res *Response) {
		body := req.Body
		newUser := User{}
		err := json.Unmarshal(body, &newUser)
		if err != nil {
			log.Println(err.Error())
		}
		newUser.Id = len(users) + 1
		users = append(users, newUser)
		res.SetStatus(Created)
		res.Send()
	})

	server.GET("/chau", func(req Request, res *Response) {
		res.Send("Chau Mundo!")
	})
	server.Start()
}
