package main

import "github.com/go-martini/martini"

func main() {

	Api()
}

func Api() {

	m := martini.Classic()
	m.Get("/greet", func() string {
		return "Hello world! \n"
	})
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + PrintName(params["name"]) + "\n"
	})
	m.RunOnAddr(":8080")
}

func PrintName(name string) string {
	return name
}
