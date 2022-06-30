package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Susan Smith",
		Age:  11,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	// user := struct {
	// 	Name string
	// 	Age int
	// 	}{
	// 		Name : "Susan Smith",
	// 		Age: 11,
	// 	}

	fmt.Println(user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}