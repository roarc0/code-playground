package database

import (
	"context"
	"log"

	"github.com/roarc0/jobs/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	jobDatabaseName = "jobs"
)

type DB struct {
	client *mongo.Client
}

func Connect(ctx context.Context, connectionURI string) (*DB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &DB{
		client: client,
	}, nil
}

func (db *DB) GetJob(ctx context.Context, id string) (*model.JobListing, error) {
	jobCollec := db.client.Database(jobDatabaseName).Collection("jobs")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err = jobCollec.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		return nil, err
	}
	return &jobListing, nil
}

func (db *DB) GetJobs(ctx context.Context) ([]*model.JobListing, error) {
	jobCollec := db.client.Database(jobDatabaseName).Collection("jobs")
	var jobListings []*model.JobListing
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &jobListings); err != nil {
		return nil, err
	}

	return jobListings, nil
}

func (db *DB) CreateJobListing(ctx context.Context, jobInfo model.CreateJobListingInput) (*model.JobListing, error) {
	jobCollec := db.client.Database(jobDatabaseName).Collection("jobs")
	res, err := jobCollec.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{ID: insertedID, Title: jobInfo.Title, Company: jobInfo.Company, Description: jobInfo.Description, URL: jobInfo.URL}
	return &returnJobListing, nil
}

func (db *DB) UpdateJobListing(ctx context.Context, jobId string, jobInfo model.UpdateJobListingInput) (*model.JobListing, error) {
	jobCollec := db.client.Database(jobDatabaseName).Collection("jobs")
	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}

	_id, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing, nil
}

func (db *DB) DeleteJobListing(ctx context.Context, jobId string) (*model.DeleteJobResponse, error) {
	jobCollec := db.client.Database(jobDatabaseName).Collection("jobs")

	_id, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	_, err = jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.DeleteJobResponse{DeletedJobID: jobId}, nil
}
