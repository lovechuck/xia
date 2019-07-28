package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"time"
)

var client *mongo.Client

func get() *mongo.Client {
	if client != nil {
		return client
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}

func Insert(database string,collection string,doc interface{}) error {
	conn := get().Database(database).Collection(collection)
	_, err := conn.InsertOne(context.TODO(), doc)
	return err
}

func GetAll(database string,collection string,result interface{}) error {
	conn := get().Database(database).Collection(collection)
	ctx := context.TODO()
	cur, err := conn.Find(ctx, bson.D{})
	defer cur.Close(ctx)
	if err == mongo.ErrNilCursor {
		return nil
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	v := reflect.ValueOf(result)
	v.Set(reflect.MakeSlice(v.Type(), 100, 100))
	i:=0
	for cur.Next(ctx) {
		tmp:=reflect.New(v.Type().Elem())
		err := cur.Decode(&tmp)
		if err != nil{
			return err
		}
		v.Index(i).Set(tmp)
		i++
	}

	return nil
}

