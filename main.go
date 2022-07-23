package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myFirstGoApp/pkg/mod2"
	"net/http"
)

// "io"
// "net/http"
// "os"
//"rsc.io/quote"

func main() {
	/*
		fmt.Println("Hello, World!")

		mod2.Hi("вывывы")
	*/

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	// fmt.Println("1 вариант")

	// var result1 map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&result1)

	// fmt.Printf("userId = %v \n", result1["userId"])
	// fmt.Printf("Id = %v \n", result1["id"])
	// fmt.Printf("title = %v \n", result1["title"])
	// fmt.Printf("completed = %v \n", result1["completed"])
	//mod2.Print1("https://jsonplaceholder.typicode.com/todos/1")

	// fmt.Println("2 вариант")
	// data, err := ioutil.ReadAll(resp.Body)

	// type StructResult struct {
	// 	userId    int
	// 	id        int
	// 	title     string
	// 	completed bool
	// }

	// var result2 StructResult
	// json.Unmarshal(data, &result2)

	// fmt.Println("Struct is:", result2)

	// fmt.Printf("userId = %v \n", result2.userId)
	// fmt.Printf("Id = %v \n", result2.id)
	// fmt.Printf("title = %v \n", result2.title)
	// fmt.Printf("completed = %v \n", result2.completed)
	//mod2.Print2("https://jsonplaceholder.typicode.com/todos/1")

	path := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("2 вариант")
	data, err := ioutil.ReadAll(resp.Body)

	var result2 mod2.Todo
	json.Unmarshal(data, &result2)

	fmt.Printf("userId = %v \n", result2.UserId)
	fmt.Printf("Id = %v \n", result2.Id)
	fmt.Printf("title = %v \n", result2.Title)
	fmt.Printf("completed = %v \n", result2.Completed)

}
