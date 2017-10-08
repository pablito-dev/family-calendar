package databaseutils

import "gopkg.in/mgo.v2"

func ConnectToDatabase()(*mgo.Session) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func PrepareDatabase(s *mgo.Session) {
	session := s.Copy()

	defer session.Close()

	collection := SetCollectionInNewSession(session)

	index := mgo.Index{
		Key: []string{"id"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func SetCollectionInNewSession(s *mgo.Session)(c *mgo.Collection) {
	collection := s.DB("calendar").C("events")

	return collection
}

