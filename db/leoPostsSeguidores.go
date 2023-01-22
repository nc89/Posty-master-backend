package db

import (
	"context"
	"time"

	"github.com/NicolasCardenas/Posty-master/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoPostsSeguidores lee los tweets de mis seguidores */
func LeoPostsSeguidores(ID string, pagina int) ([]models.DevuelvoPostsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("posty")
	col := db.Collection("relation")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	//match para buscar el id de la relacion
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	//lookup para unir dos tablas
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "post",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "post",
		}})
	//unwind permite que todos los documentos lleguen igual, para omitir dado el caso de haber informacion repetida
	condiciones = append(condiciones, bson.M{"$unwind": "$post"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"post.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	//funcion aggregate de mongodb
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoPostsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
