package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongourl string
	Database string
)

type MyData struct {
	Name string
}

func insert() {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("MyData")
	err := c.Insert(MyData{"Joseph"})
	if err != nil {
		log.Fatal(err)
	}

}

var callDb = func(name string) MyData {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("MyData")

	var result MyData
	c.Find(bson.M{"name": name}).One(&result)

	return result
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial(mongourl)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func init() {

	if mongourl = os.Getenv("MONGO_URL"); mongourl == "" {
		mongourl = "mongodb://localhost"
	}

	Database = "gweb"
}
