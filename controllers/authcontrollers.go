package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/models"
	"server/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userCollection = config.GetCollection("users")
	var foundUser models.User
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	error := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	fmt.Println(error)

	if error == nil {
		http.Error(w, "user already exists", http.StatusFound)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	user.ID = primitive.NewObjectID().Hex()
	defer cancel()
	_, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		http.Error(w, "Error regitering User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userCollection = config.GetCollection("users")
	var user models.User
	var foundUser models.User

	json.NewDecoder(r.Body).Decode(&user)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(foundUser.Email)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
