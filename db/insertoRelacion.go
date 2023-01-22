package db

import (
	"context"
	"time"

	"github.com/NicolasCardenas/Posty-master/models"
)

/*InsertoRelacion graba la relación en la BD */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("posty")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
