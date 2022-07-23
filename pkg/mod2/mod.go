package mod2

// func Hi(name string) {
// 	fmt.Println(name)
// }

type Todo struct {
	UserId    int    `example:"UserId"`
	Id        int    `example:"Id"`
	Title     string `example:"Title"`
	Completed bool   `example:"Completed"`
}

/*
func Print1(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("1 вариант")

	var result1 map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result1)

	fmt.Printf("userId = %v \n", result1["UserId"])
	fmt.Printf("Id = %v \n", result1["Id"])
	fmt.Printf("title = %v \n", result1["Title"])
	fmt.Printf("completed = %v \n", result1["Completed"])
}
*/

/*
func Print2(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("2 вариант")
	data, err := ioutil.ReadAll(resp.Body)

	var result2 Todo
	json.Unmarshal(data, &result2)

	fmt.Printf("userId = %v \n", result2.UserId)
	fmt.Printf("Id = %v \n", result2.Id)
	fmt.Printf("title = %v \n", result2.Title)
	fmt.Printf("completed = %v \n", result2.Completed)

}
*/
