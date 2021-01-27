package db

import (
	"context"
	"time"

	"github.com/KenethSandoval/goTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil de usuario*/
func ModificoRegistro(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("clonetwitter")
	col := db.Collection("users")

	registro := make(map[string]interface{})

	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		registro["lasname"] = u.LastName
	}

	registro["birthdate"] = u.Birthdate

	if len(u.Email) > 0 {
		registro["email"] = u.Email
	}

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		registro["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		registro["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		registro["website"] = u.WebSite
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
