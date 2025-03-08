package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"mongo-basic/internal/model"
	"mongo-basic/internal/view"
	"mongo-basic/pkg/util"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// Methods have to be capitalized to be exported
// GetUser handles fetching a user by ID
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Convert id string to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get collection
	collection := uc.client.Database("go-test").Collection("users")

	// Fetch user
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
		} else {
			util.InternalServerErr(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		util.BadRequestErr(w, err)
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.BadRequestErr(w, err)
		return
	}

	// Assign a new ObjectID
	user.ID = primitive.NewObjectID()

	// Get collection
	collection := uc.client.Database("go-test").Collection("users")

	// Insert user into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		util.InternalServerErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		util.InternalServerErr(w, err)
	}
}

// DeleteUser handles deleting a user by ID
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Convert id string to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get collection
	collection := uc.client.Database("go-test").Collection("users")

	// Delete user
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		util.InternalServerErr(w, err)
		return
	}

	// If no document was deleted, return 404
	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user %v\n", objID)
}

func (uc *UserController) GetHome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	view.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
