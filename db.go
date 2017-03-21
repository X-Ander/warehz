package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	dbVersion = "0.1"
)

type MetaData struct {
	Version string
}

var (
	db *mgo.Database
	oc *mgo.Collection  // Object collection
	mc *mgo.Collection  // Metadata collection
	meta MetaData
)

func prepareDB() {
	db = dbSes.DB("")
	oc = db.C("objects")
	mc = db.C("meta")
	checkMeta()
}

func checkMeta() {
	err := mc.Find(nil).One(&meta)
	if err == mgo.ErrNotFound {
		meta.Version = dbVersion
		err = mc.Insert(&meta)
	}
	if err != nil {
		log.Fatal(err)
	}
	if meta.Version != dbVersion {
		log.Fatalf("Database version (%s) is not supported\n", meta.Version)
	}
}
