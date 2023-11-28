package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "bob Smith",
		Bio:  "hello",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
