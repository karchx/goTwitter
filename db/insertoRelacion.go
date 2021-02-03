package db

import (
	"context"
	"time"

	"github.com/KenethSandoval/goTwitter/models"
)

/* InsertoRelacion inserta la relacion entre los usuarios */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("clonetwitter")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
