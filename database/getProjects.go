package database

import (
	"context"
	"errors"
	"os"
	"stratus-core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProjects() ([]models.Project, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return nil, errors.New("Database name not present")
	}

	if MongoDBClient == nil {
		return nil, errors.New("MongoDB client is not initialized")
	}

	projectsCollection := MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	cursor, err := projectsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var projects []models.Project
	for cursor.Next(context.TODO()) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
