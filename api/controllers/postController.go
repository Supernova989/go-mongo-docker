package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mongo-api/database"
	m "go-mongo-api/models"
	"io/ioutil"
	"net/http"
)

var FindPosts = func(db database.PostInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter interface{}
		q := r.URL.Query().Get("q")
		if q != "" {
			err := json.Unmarshal([]byte(q), &filter)
			if err != nil {
				WriteJsonResponse(w, nil, http.StatusBadRequest, nil, "")
				return
			}
		}
		res, err := db.Find(filter)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, err.Error())
			return
		}
		WriteJsonResponse(w, res, http.StatusOK, nil, "")
	}
}

var GetPost = func(db database.PostInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		res, err := db.Get(id)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, err.Error())
			return
		}
		WriteJsonResponse(w, res, http.StatusOK, nil, "")

	}
}

var CreatePost = func(db database.PostInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := m.Post{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, "")
			return
		}
		err = json.Unmarshal(body, &post)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, "")
			return
		}
		res, err := db.Insert(post)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, err.Error())
			return
		}
		WriteJsonResponse(w, res, http.StatusCreated, nil, "")
	}
}

var PatchPost = func(db database.PostInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, "")
			return
		}
		var post interface{}
		err = json.Unmarshal(body, &post)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, "")
			return
		}
		res, err := db.Update(id, post)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, err.Error())
			return
		}
		WriteJsonResponse(w, res, http.StatusOK, nil, "")
	}
}

var DeletePost = func(db database.PostInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		res, err := db.Delete(id)
		if err != nil {
			WriteJsonResponse(w, nil, http.StatusBadRequest, nil, err.Error())
			return
		}
		WriteJsonResponse(w, res, http.StatusOK, nil, "")
	}
}
