package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	FirstName string  `json:"name"`
	LastName  string  `json:"surname"`
	Email     string  `json:"email"`
	Gender    string  `json:"gender"`
	Loan      float64 `json:"loan"`
}

func main() {
	fmt.Println(AddingTest())
	fmt.Println(DeleteTest())
	fmt.Println(UpdateTest())
	fmt.Println(GetTest())
}

func AddingTest() string{
	person := Person{
		FirstName: "oleg1111",
		LastName:  "osyka",
		Email:     "olegosyka@gmail.com",
		Gender:    "Male",
		Loan:      254.3,
	}
	personJSON, _ := json.Marshal(person)
	resp, err := http.Post("http://localhost:1234/add", "application/json", bytes.NewReader(personJSON))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func DeleteTest() string{
	resp, err := http.Post("http://localhost:1234/delete", "application/json", bytes.NewReader([]byte("121")))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func UpdateTest() string{
	type IDPerson struct {
		ID     int    `json:"id"`
		Person Person `json:"person"`
	}
	idPerson := IDPerson{
		ID: 112,
		Person: Person{
			FirstName: "jeka",
			Email:     "olegosyka@gmail.comcomcom",
			Gender:    "Male",
			Loan:      254.3,
		},
	}
	idPersonJSON, _ := json.Marshal(idPerson)
	resp, err := http.Post("http://localhost:1234/update", "application/json", bytes.NewReader(idPersonJSON))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetTest() string{
	resp, err := http.Get("http://localhost:1234/persons?name=o5leg")
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}