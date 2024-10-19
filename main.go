package main

type User struct {
	Id       int
	Username string
}

var users = []User{{1, "juan"}, {2, "pedro"}, {3, "pablo"}}

func main() {
	server := Server{Port: 3000}
	server.GET("/", func(req Request, res *Response) {
		res.Json(users)
	})
	server.GET("/chau", func(req Request, res *Response) {
		res.Send("Chau Mundo!")
	})
	server.Start()
}
