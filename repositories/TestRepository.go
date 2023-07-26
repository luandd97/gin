package repositories

import "go.mongodb.org/mongo-driver/mongo"

type TestRepository interface{}

type testConnection struct{
	mongodb *mongo.Database
}

func NewTestRepository(mongodb *mongo.Database) TestRepository {
	return &testConnection{
		mongodb: mongodb,
	}
}
