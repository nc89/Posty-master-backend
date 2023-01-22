package db

import (
	"context"
	"time"

	"github.com/NicolasCardenas/Posty-master/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoPosts lee los posts de un perfil */
func LeoPosts(ID string, pagina int64) ([]*models.DevuelvoPosts, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("posty")
	col := db.Collection("post")

	var resultados []*models.DevuelvoPosts

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)

	//los ordena por fecha en orden descendente
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoPosts
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
