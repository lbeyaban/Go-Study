package restful

import (
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {

	response, err := http.Get("http://localhost:3000/users")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	resBodyBytes, _ := io.ReadAll(response.Body)

	resBodyString := string(resBodyBytes)

	fmt.Println("Body : ", resBodyString)

}
