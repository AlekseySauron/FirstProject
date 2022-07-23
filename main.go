package main

import (
	"fmt"
	// "io"
	// "net/http"
	// "os"
	myFirstModule "github.com/AlekseySauron/myFirstModule"
	//"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	myFirstModule.Hi("ввв")

	//fmt.Println(testmod.Hi("roberto"))

	// message := quote.Hello()
	// fmt.Printf(message)

	//return
	/*
		fmt.Println("dsds")

		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		io.Copy(os.Stdout, resp.Body)
	*/

	/*
	   for true {

	       bs := make([]byte, 1014)
	       n, err := resp.Body.Read(bs)
	       fmt.Println(string(bs[:n]))

	       if n == 0 || err != nil{
	           break
	       }
	   }
	*/

}
