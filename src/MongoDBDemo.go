package main
import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println()
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	var p1 Person
	p1.Name = "shenyan"
	p1.Phone = "123456787899999"
//	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	err = c.Insert(&p1)
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "shenyan"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}
