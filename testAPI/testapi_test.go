package main

import ("testing"

)

func TestAddingTest(t *testing.T){
	if AddingTest() != "Person added"{
		t.Error("Expected response from server: Person added")
	}
}

func TestDeleteTest(t *testing.T){
	if DeleteTest() != "user deleted"{
		t.Error("Expected response from server: user deleted")
	}
}

func TestUpdateTest(t *testing.T){
	if UpdateTest() != "update succeseful"{
		t.Error("Expected response from server: update succeseful")
	}
}

func TestGetTest(t *testing.T){
	if GetTest() == "no results for this query"{
		t.Error("Expected response from server: list of users")
	}
}
