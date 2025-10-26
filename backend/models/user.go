package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username"`
	Email          string             `bson:"email"`
	HashedPassword string             `bson:"hashedPassword"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	SessionToken   string             `bson:"sessionToken,omitempty"`
	CSRFToken      string             `bson:"csrfToken,omitempty"`
}

func (u *User) Save(db *mongo.Database) error {
	collection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if u.ID.IsZero() {
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		result, err := collection.InsertOne(ctx, u)
		if err != nil {
			return err
		}
		u.ID = result.InsertedID.(primitive.ObjectID)
		return nil
	}

	u.UpdatedAt = time.Now()
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": u.ID},
		bson.M{"$set": u},
	)
	return err
}

func FindUserByEmail(email string, db *mongo.Database) (*User, error) {
	collection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserSession(userID primitive.ObjectID, sessionToken, csrfToken string, db *mongo.Database) error {
	collection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"sessionToken": sessionToken,
			"csrfToken":    csrfToken,
			"updatedAt":    time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": userID}, update)
	return err
}

func FindUserByUsername(username string, db *mongo.Database) (*User, error) {
	collection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
