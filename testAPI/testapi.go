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
	//fmt.Println(len(AddingUncTest()))
	fmt.Println(UpdateUncTest())
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
	resp, err := http.Post("http://localhost:1234/persons", "application/json", bytes.NewReader(personJSON))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func AddingUncTest() string{
	personJSON, _ := json.Marshal("duhweojheoewhoi")
	resp, err := http.Post("http://localhost:1234/persons", "application/json", bytes.NewReader(personJSON))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func DeleteTest() string{
	//id := "117"
	//resp, err := http.Post("http://localhost:1234/delete", "application/json", bytes.NewReader([]byte(id)))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//body, _ := ioutil.ReadAll(resp.Body)
	//return string(body)
	//idJSON, _ := json.Marshal(id)
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE", "http://localhost:1234/persons?id=117", nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
	return "q"
}

func DeleteUncTest() string{
	//id := "117"
	//resp, err := http.Post("http://localhost:1234/delete", "application/json", bytes.NewReader([]byte(id)))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//body, _ := ioutil.ReadAll(resp.Body)
	//return string(body)
	//idJSON, _ := json.Marshal(id)
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE", "http://localhost:1234/persons?ighjd=117", nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
	return "q"
}

func UpdateUncTest() string{
	type IDEmail struct {
		ID     string    `json:"id"`
		Jopa string
	}
	idEmail := IDEmail{
		ID: "j113k",
		Jopa: "BigAss",
	}
	idPersonJSON, _ := json.Marshal(idEmail)
	client := &http.Client{}
	req, err := http.NewRequest(
		"PUT", "http://localhost:1234/persons", bytes.NewReader(idPersonJSON),
	)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func UpdateTest() string{
	type IDEmail struct {
		ID     int    `json:"id"`
		Email string `json:"email"`
	}
	idEmail := IDEmail{
		ID: 113,
		Email: "qweqweqwewqeqweqeqwewqewqewqeqwe",
	}
	idPersonJSON, _ := json.Marshal(idEmail)
	client := &http.Client{}
	req, err := http.NewRequest(
		"PUT", "http://localhost:1234/persons", bytes.NewReader(idPersonJSON),
	)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetTest() string{
	resp, err := http.Get("http://localhost:1234/persons?name=Zak")
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetUncTest() string{
	resp, err := http.Get("http://localhost:1234/persons?ass=ZakZakMotherFucker&toDate=fucker")
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}