package db

import (
	"context"
	"fmt"
	"simple_api/env"
	"simple_api/module/rowdata"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo() error {
	credential := options.Credential{
		Username:   "root",
		Password:   "zxc12345",
		AuthSource: "admin",
	}
	opts := options.Client().SetHosts(strings.Split("example-mongo:27017", ",")).SetConnectTimeout(10 * time.Second).SetAuth(credential).SetReadPreference(readpref.SecondaryPreferred())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	env.Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}
	err = env.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	return nil
}

func CreateRowData(d *rowdata.RowData) error {
	var optsS = options.Session()
	session, err := env.Client.StartSession(optsS.SetDefaultReadPreference(readpref.Secondary()))
	if err != nil {
		return err
	}
	// New a collection of Record.
	col := session.Client().Database("climate").Collection("rowdata")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer session.EndSession(ctx)
	//doc := bson.D{{"title", "Invisible Cities"}, {"author", "Italo Calvino"}, {"year_published", 1974}}
	_, err = col.InsertOne(ctx, &d)
	if err != nil {
		return err
	}
	return nil
}
func FindOneRowData(uuid string) (err error, d *rowdata.RowData) {
	var optsS = options.Session()
	session, err := env.Client.StartSession(optsS.SetDefaultReadPreference(readpref.Secondary()))
	if err != nil {
		d = nil
		return err, d
	}
	// New a collection of Record.
	col := session.Client().Database("climate").Collection("rowdata")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer session.EndSession(ctx)
	d = new(rowdata.RowData)
	doc := bson.D{{"uuid", uuid}}
	err = col.FindOne(ctx, doc).Decode(&d)
	if err != nil {
		d = nil
		return err, d
	}
	return nil, d
}
func UpdateData(uuid string, updatedata *rowdata.RowData) (err error, res *mongo.UpdateResult) {
	var optsS = options.Session()
	session, err := env.Client.StartSession(optsS.SetDefaultReadPreference(readpref.Secondary()))
	if err != nil {
		res = nil
		return err, res
	}
	// New a collection of Record.
	col := session.Client().Database("climate").Collection("rowdata")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer session.EndSession(ctx)
	filter := bson.D{{"uuid", uuid}}
	doc := bson.D{
		{"uuid", updatedata.Uuid},
		{"parentid", updatedata.Parentid},
		{"comment", updatedata.Comment},
		{"author", updatedata.Author},
		{"update", updatedata.Update},
		{"favorite", updatedata.Favorite},
	}
	update := bson.D{{"$set", doc}}
	res, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		res = nil
		return err, res
	}
	return nil, res
}
func DeleteRowData(uuid string) (err error, res *mongo.DeleteResult) {
	var optsS = options.Session()
	session, err := env.Client.StartSession(optsS.SetDefaultReadPreference(readpref.Secondary()))
	if err != nil {
		res = nil
		return err, res
	}
	// New a collection of Record.
	col := session.Client().Database("climate").Collection("rowdata")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer session.EndSession(ctx)
	filter := bson.D{{"uuid", uuid}}
	res, err = col.DeleteOne(ctx, filter)
	if err != nil {
		res = nil
		return err, res
	}
	return nil, res
}
