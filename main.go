package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "io"
	// "net/http"
	// "os"
	//"rsc.io/quote"
)

func main() {
	/*
		fmt.Println("Hello, World!")

		res := myFirstModule.Hi("ввв")
		fmt.Println(res)

		mod2.Hi("вывывы")
	*/

	//logrus

	//fmt.Println(testmod.Hi("roberto"))

	// message := quote.Hello()
	// fmt.Printf(message)

	//return

	//fmt.Println("START")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	/*
		fmt.Println("1 вариант")

		var result1 map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result1)

		fmt.Printf("userId = %v \n", result1["userId"])
		fmt.Printf("Id = %v \n", result1["id"])
		fmt.Printf("title = %v \n", result1["title"])
		fmt.Printf("completed = %v \n", result1["completed"])
	*/

	fmt.Println("2 вариант")
	data, err := ioutil.ReadAll(resp.Body)

	type StructResult struct {
		userId    int
		id        int
		title     string
		completed bool
	}

	var result2 StructResult
	json.Unmarshal(data, &result2)

	fmt.Println("Struct is:", result2)

	fmt.Printf("userId = %v \n", result2.userId)
	fmt.Printf("Id = %v \n", result2.id)
	fmt.Printf("title = %v \n", result2.title)
	fmt.Printf("completed = %v \n", result2.completed)

}
