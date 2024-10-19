package main

func main() {
	server := Server{Port: 3000}
	server.GET("/", func(req Request, res *Response) {
		res.Write("Hola Mundo!")
	})
	server.GET("/chau", func(req Request, res *Response) {
		res.Write("Chau Mundo!")
	})
	server.Start()
}
