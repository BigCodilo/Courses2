package main

import (
	"encoding/json"
	"errors"
	"github.com/BigCodilo/Courses2/logger"
	"github.com/BigCodilo/Courses2/logic"
	"net/http"
	"strconv"
	"time"
)


func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Print("GET for", r.RequestURI, "\n User agent: ", r.UserAgent(),  "\n Cookies: ", r.Cookies())
	logger.Info.Println("GET for", r.URL)
	name := r.URL.Query().Get("name")
	fromDate := r.URL.Query().Get("fromDate")
	toDate := r.URL.Query().Get("toDate")
	gender := r.URL.Query().Get("gender")
	persons, err := DB.GetAllPersons()
	if err != nil {
		http.Error(w, "something wrong with database", http.StatusNotFound)
		logger.Debug.Println( "unsuccessfully.\n")
		logger.Info.Println( "unsuccessfully.\n")
		logger.Error.Println("GET for", r.RequestURI, " ---", err, "---")
		return
	}
	err = nil
	validPersons := logic.Persons{}
	//Если в урле есть запрос по полу или мени, то в этом цикле это отфильтрует
	for _, v := range persons {
		nameFlag := true
		genderFlag := true
		if len(gender) > 0 && v.Gender != gender {
			genderFlag = false
		}
		if len(name) > 0 && v.FirstName != name {
			nameFlag = false
		}
		if nameFlag && genderFlag {
			validPersons = append(validPersons, v)
		}
	}
	//Если есть запрос по дате регитсрации  то вот здесь его обработает
	if len(fromDate) > 0 && len(toDate) > 0 {
		validPersons, err = validPersons.GetInRegisterRange(fromDate, toDate)
	}
	if len(fromDate) == 0 && len(toDate) > 0 {
		validPersons, err = validPersons.GetInRegisterRange("01/01/2001", toDate)
	}
	if len(fromDate) > 0 && len(toDate) == 0 {
		validPersons, err = validPersons.GetInRegisterRange(fromDate, "12/31/3000")
	}
	if err != nil {
		//http.Error(w, "uncorrect date form", http.StatusNotFound)
		w.Write([]byte("uncorrect date form111"))
		logger.Debug.Println( "unsuccessfully.\n")
		logger.Info.Println( "unsuccessfully.\n")
		logger.Error.Println("GET for", r.RequestURI,  " ---", err, "---")
		return
	}
	//Если ненайдено людей с удовлетворяющими требованиями то напишет об этом
	if len(validPersons) == 0 {
		http.Error(w, "no results for this query", http.StatusNotFound)
		logger.Debug.Println( "unsuccessfully.\n")
		logger.Info.Println( "unsuccessfully.\n")
		logger.Error.Println("GET for", r.RequestURI, " ---", errors.New("no valid persons"), "---")
		return
	}
	validPersonsJSON, err := json.Marshal(validPersons)
	logger.Debug.Println( "successfully.\n")
	logger.Info.Println( "successfully.\n")
	w.Write(validPersonsJSON)
}

func AddPersonHandler(w http.ResponseWriter, r *http.Request) {
	//personJSON := r.FormValue("person")
	person := logic.Person{}
	//err := json.Unmarshal([]byte(personJSON), &person)
	err := json.NewDecoder(r.Body).Decode(&person)
	logger.Debug.Print("POST for", r.RequestURI, "\n User agent: ", r.UserAgent(), "\n Body: ", person, "\n Cookies: ", r.Cookies())
	logger.Info.Println("POST for", r.URL)
	if err != nil {
		//http.Error(w, "uncorrect format", http.StatusBadRequest)
		w.Write([]byte("uncorrect format"))
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("POST with body ", person, "to", r.RequestURI, " ---", err, "---")
		return
	}
	person.RegisterDate = time.Now()
	logic.SetIotaGender(person)
	err = DB.Add(person)
	if err != nil {
		http.Error(w, "Problem with database", 418)
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("POST with body ", person, "to", r.RequestURI, " ---", err, "---")
		return
	}
	logger.Debug.Println("successfully\n")
	logger.Info.Println("successfully\n")
	w.Write([]byte("Person added"))
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	idS := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idS)
	logger.Debug.Print("DELETE for", r.RequestURI, "\n User agent: ", r.UserAgent(), "\n Body: ", id,  "\n Cookies: ", r.Cookies())
	logger.Info.Println("DELETE for", r.URL)
	if err != nil {
		//http.Error(w, "Uncorrect format -", http.StatusBadRequest)
		w.Write([]byte("uncorrect format"))
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---")
		return
	}
	//id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, "Uncorrect format", http.StatusBadRequest)
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---")
		return
	}
	err = DB.Delete(id)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("POST with body ", id, "to", r.RequestURI, "---", err, "---")
		return
	}
	logger.Debug.Println("successfully.\n")
	logger.Info.Println("successfully.\n")
	w.Write([]byte("user deleted"))
}

func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	type IDEmail struct {
		ID     int          `json:"id"`
		Email string `json:"email"`
	}
	idEmail := IDEmail{}

	err := json.NewDecoder(r.Body).Decode(&idEmail)
	logger.Debug.Print("PUT for", r.RequestURI, "\n User agent: ", r.UserAgent(), "\n Body: ", idEmail,  "\n Cookies: ", r.Cookies())
	logger.Info.Println("PUT for", r.URL)
	if err != nil {
		//http.Error(w, "Something wrong", 418)
		w.Write([]byte("Something wrong"))
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("PUT with body ", idEmail, "to", r.RequestURI, "---", err, "---")
		return
	}
	//err = DB.Update(idEmail.ID, idEmail.Email)
	err = DB.Update(idEmail.ID, idEmail.Email)
	if err != nil {
		http.Error(w, "Something wrong", 418)
		logger.Debug.Println("unsuccessfully.\n")
		logger.Info.Println("unsuccessfully.\n")
		logger.Error.Println("PUT with body ", idEmail, "to", r.RequestURI, "---", err, "---")
		return
	}
	logger.Debug.Println("successfully.\n")
	logger.Info.Println("successfully.\n")
	w.Write([]byte("update succeseful"))
}
