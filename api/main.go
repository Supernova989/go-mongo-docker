package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-api/config"
	"go-mongo-api/controllers"
	"go-mongo-api/database"
	"log"
	"net/http"
)

func main() {
	postClient := &database.PostClient{
		Col: database.GetDB().Collection(config.GetConfig().Mongo.Collections.Posts),
		Ctx: database.GetGlobalContext(),
	}

	r := mux.NewRouter()

	r.HandleFunc("/posts", controllers.FindPosts(postClient)).Methods("GET")
	r.HandleFunc("/posts/{id}", controllers.GetPost(postClient)).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost(postClient)).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.PatchPost(postClient)).Methods("PATCH")
	r.HandleFunc("/posts/{id}", controllers.DeletePost(postClient)).Methods("DELETE")

	log.Println(fmt.Sprintf("App's running on port: 8080"))
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
