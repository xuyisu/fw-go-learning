package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type UserMongo struct {
	Name  string
	Phone string
	Email string
}

var (
	client     *mongo.Client
	err        error
	collection *mongo.Collection
)

// mongo 连接
func main() {

	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:123456@localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//选择数据库 my_db里的某个表 ，写入的时候会默认创建这个库和表
	collection = client.Database("fw_mongo").Collection("fw_col")

	// insert
	//insertMongo(err, collection)

	//批量insert
	//insertManyMongo(err, collection)

	//更新
	//updateOneMongo()

	//更新多条
	//updateManyMongo()

	//查询数据
	//bson.D 是一组key-value的集合  有序
	//bson.M是一个map 无序的key-value集合，在不要求顺序的情况下可以替代bson.D
	//bson.A 是一个数组
	//bson.E是只能包含一个key-value的map
	//var user UserMongo

	//查询一个
	findMongo()

	//查询多个
	//findManyMongo()

	//删除单个
	//deleteOneMongo()

	//删除多条数据
	//deleteManyMongo()

}

func deleteOneMongo() {
	filter := bson.D{{"name", "貂蝉"}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func deleteManyMongo() {
	filter := bson.D{{"name", "貂蝉"}}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func findMongo() {
	var result UserMongo
	filter := bson.D{{"name", "貂蝉"}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("查询单个 document: %+v\n", result)
}

func findManyMongo() {
	findOptions := options.Find()
	//限制两个
	findOptions.SetLimit(2)

	var results []*UserMongo
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//close cursor
	defer func() {
		if err = cur.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem UserMongo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(*result)
	}
}

func updateManyMongo() {
	update := bson.D{
		{"$set", bson.D{
			{"name", "貂蝉"},
		}},
	}
	filter := bson.D{{"name", "关羽"}}
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("匹配 %v documents and 更新了 %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func updateOneMongo() {
	update := bson.D{
		{"$set", bson.D{
			{"name", "关羽"},
		}},
	}
	filter := bson.D{{"name", "刘备1"}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("匹配 %v documents and 更新了 %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func insertManyMongo(err error, collection *mongo.Collection) {
	insertMany, err := collection.InsertMany(context.TODO(), []interface{}{
		UserMongo{
			Name:  "刘备" + strconv.Itoa(rand.Intn(11)),
			Phone: strconv.Itoa(rand.Intn(11)),
			Email: strconv.Itoa(rand.Intn(5)) + "@qq.com",
		},
		UserMongo{
			Name:  "刘备" + strconv.Itoa(rand.Intn(11)),
			Phone: strconv.Itoa(rand.Intn(11)),
			Email: strconv.Itoa(rand.Intn(5)) + "@qq.com",
		},
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	for index, insertedID := range insertMany.InsertedIDs {
		objectID := insertedID.(primitive.ObjectID)
		fmt.Println("自增ID", objectID.Hex())
		//索引
		fmt.Println("index=", index)
	}
}

func insertMongo(err error, collection *mongo.Collection) {
	userMongo := UserMongo{
		Name:  "刘备" + strconv.Itoa(rand.Intn(11)),
		Phone: strconv.Itoa(rand.Intn(11)),
		Email: strconv.Itoa(rand.Intn(5)) + "@qq.com",
	}
	insertOne, err := collection.InsertOne(context.TODO(), userMongo)
	if err != nil {
		fmt.Print(err)
	} else {
		//_id:默认生成一个全局唯一ID
		var id primitive.ObjectID
		id = insertOne.InsertedID.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}
