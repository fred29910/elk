/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func main() {
	host := "127.0.0.1:27017/paidui"
	m_session, err := mgo.Dial(host)
	if err != nil {
		err = fmt.Errorf("Dial dns %v failed %v", host, err)
		panic(err)
	}
	m_session.SetMode(mgo.Monotonic, true)
	m_session.SetPoolLimit(10000)
	MakeCollectionIndex(m_session, "xasdafsda", []string{"uid"}, false, false)
}

func MakeCollectionIndex(m_session *mgo.Session, col string, keys []string, unique, dropdup bool) {
	collection := m_session.DB("xasdafsda").C(col)

	index := mgo.Index{
		Key:        keys,
		Unique:     unique,
		DropDups:   dropdup,
		Background: true,
		Sparse:     true,
	}
	if err := collection.EnsureIndex(index); err != nil {
		log.Errorf("Create index error: %v, for collection: %v, key: %v", err, col, keys)
	} else {
		log.Infof("Create index success for collection: %v, key: %v", col, keys)
	}

}
