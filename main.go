package main

import "github.com/go-martini/martini"

func main() {
	Api()
}

func Api() {
	m := martini.Classic()

	// Root endpoint
	m.Get("/", func() string {
		return "Welcome to the GO-CI-CD application!\n"
	})

	// Greet endpoint
	m.Get("/greet", func() string {
		return "Hello world! \n"
	})

	// Dynamic hello endpoint
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + PrintName(params["name"]) + "\n"
	})

	m.RunOnAddr(":8080")
}

// Function to return the name
func PrintName(name string) string {
	return name
}
