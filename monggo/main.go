package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person struct
type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println("learn interacting with mongoDB in Go")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("monggo").C("people")
	err = c.Insert(
		&Person{"kholidfu", "+62 81 2168435 51"},
		&Person{"mamah", "+62 81 1234 4567"})
	if err != nil {
		log.Fatal(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name": "kholidfu"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
