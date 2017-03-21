package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type PersonName struct {
	till time.Time
	parts []string

}

type Person struct {
	id bson.ObjectId
	cn string  // common name
	sn string  // short name
	names []PersonName
}

func (p *Person) GetBSON() (interface{}, error) {
	return bson.M{
		"_id": p.id,
		"cn": p.cn,
		"sn": p.sn,
	}, nil
}
