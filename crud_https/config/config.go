package config

import "gopkg.in/mgo.v2"

func GetMongoDB() (*mgo.Database, error) {

	host := "mongodb://127.0.0.1:27017"
	dbName := "mongo-golang"

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil

}