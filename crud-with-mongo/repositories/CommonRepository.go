package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/models"
	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/utilities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var collection *mongo.Collection

// runs only once in app runtime
func init() {
	//create connection
	clientOptions:= options.Client().ApplyURI(utilities.CONNECTION_STRING)
	connection, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("DB connection success...")
	
	//get collection ref
	collection = connection.Database(utilities.DB_NAME).Collection(utilities.COLLECTION_NAME)
	fmt.Println("collection ref is ready...")
}

// repo - movie
func GetAllMovies() ([]primitive.M) {
	//find
	cursor, err:=collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	
	//loop cursor
	var movies []primitive.M
	
	for cursor.Next(context.Background()) {
		var movie bson.M
		err:=cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		//push
		movies = append(movies, movie)
	}

	//return
	return movies
}

func GetMovieById(movieId string) (primitive.M) {
	//find
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result :=collection.FindOne(context.Background(), filter)
	
	var movie primitive.M
	result.Decode(&movie)
	//return
	return movie
}

func CreateMovie(movie models.Movie) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("insert success, insertId:", result.InsertedID)
}

func UpdateMovieById(movieId string) {
	//setup filter
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	//setup update data
	update:=bson.M{"$set": bson.M{"watched": true}}

	//actual update ops
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update success, count:",result.ModifiedCount)
}

func DeleteMovieById(movieId string){
	//setup filter
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	// actual delete execution
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete success, count:",result.DeletedCount)
}

func DeleteAllMovies() (int64){
	//setup filter
	filter := bson.D{{}}

	// actual delete execution
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete success, count:",result.DeletedCount)
	return result.DeletedCount
}
