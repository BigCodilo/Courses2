package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequests(t *testing.T){
	Convey("Gives some response from server", t, func(){
		expectingResponseAdd := "Person added"
		unExpectingResponseAdd := "uncorrect format"
		expectingResponseUpdate := "update succeseful"
		unExpectingResponseUpdate := "Something wrong"
		expectingResponseDelete := "user deleted"
		unExpectingResponseDelete := "uncorrect format"
		Convey("When send POST request", func(){
			Convey("Send correctly request", func(){
				So(expectingResponseAdd, ShouldEqual, AddingTest())
			})
			Convey("Send uncorrectly request", func(){
				So(unExpectingResponseAdd, ShouldEqual, AddingUncTest())
			})
		})
		Convey("When send PUT request", func(){
			Convey("Send correctly request", func(){
				So(expectingResponseUpdate, ShouldEqual, UpdateTest())
			})
			Convey("Send uncorrectly request", func(){
				So(unExpectingResponseUpdate, ShouldEqual, UpdateUncTest())
			})
		})
		Convey("When send DELETE request", func(){
			Convey("Send correctly request", func(){
				So(expectingResponseDelete, ShouldEqual, DeleteTest())
			})
			Convey("Send uncorrectly request", func(){
				So(unExpectingResponseDelete, ShouldEqual, DeleteUncTest())
			})
		})
	})

}

//func TestDeleteTest(t *testing.T){
//	if DeleteTest() != "user deleted"{
//		t.Error("Expected response from server: user deleted")
//	}
//}
//
//func TestUpdateTest(t *testing.T){
//	if UpdateTest() != "update succeseful"{
//		t.Error("Expected response from server: update succeseful")
//	}
//}
//
//func TestGetTest(t *testing.T){
//	if GetTest() == "no results for this query"{
//		t.Error("Expected response from server: list of users")
//	}
//}
