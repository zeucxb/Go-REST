package models

import mgo "gopkg.in/mgo.v2"

// GetSession return a mongo session
func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo runnig?
	if err != nil {
		panic(err)
	}

	return s
}
