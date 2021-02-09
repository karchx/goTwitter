package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/KenethSandoval/goTwitter/models"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("clonetwitter")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "usuarioid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cusor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cusor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
