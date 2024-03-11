package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"watchlistAPI/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://$USER:$PASSWORD@cluster0.drjwtyf.mongodb.net/?retryWrites=true&w=majority"
const dbName = "moviecollection"
const colName = "watchlist"

var collection *mongo.Collection

// Connect to database
func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection success!")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready.")
}

// func checkNil(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// helpers - to file
// insert a record
func addMovie(movie model.Watchlist) {
	added, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Added to collection, ID is", added.InsertedID)
}

// update a record
func updateMovieId(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie updated:", updateResult.ModifiedCount)
}

// delete a record
func deleteMovieId(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted:", deleteResult.DeletedCount)
}

// delete all records
func deleteAllMovies() int64 {
	deleteAllResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies deleted:", deleteAllResult.DeletedCount)
	return deleteAllResult.DeletedCount
}

// get all records
func getAllMovies() []primitive.M {
	getAllResult, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for getAllResult.Next(context.Background()) {
		var movie bson.M
		err := getAllResult.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer getAllResult.Close(context.Background())
	return movies
}

// actual controllers - to file
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Watchlist
	_ = json.NewDecoder(r.Body).Decode(&movie)
	addMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateIdAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateMovieId(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovieId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovieId(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	delcount := deleteAllMovies()
	json.NewEncoder(w).Encode(delcount)
}
