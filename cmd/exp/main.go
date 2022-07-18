package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Susan Smith",
		Bio:  `<script>alert("Haha, you have been hexxed!");<script>`,
		Age:  11,
	}

	// user := struct {
	// 	Name string
	// 	Age int
	// 	}{
	// 		Name : "Susan Smith",
	// 		Age: 11,
	// 	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
