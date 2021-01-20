package db

import (
	"context"
	"time"

	"github.com/KenethSandoval/goTwitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertRegistry es la parada final con la DB para insertar los datos del usuario */
func InsertRegistry(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("clonetwitter")
	col := db.Collection("users")

	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
