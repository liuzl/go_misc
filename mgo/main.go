package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Record struct {
	Id   string `json:"_id" bson:"-"`
	Name string `json:"name", bson:"name"`
}

func main() {
	conn, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	conn.SetMode(mgo.Monotonic, true)
	db := conn.DB("test")
	c := db.C("table")
	rec := &Record{"1", "liang"}
	info, err := c.UpsertId(rec.Id, rec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)
}
