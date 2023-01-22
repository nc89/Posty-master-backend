package models

import "time"

/*GraboPost es el formato o estructura que tendrá nuestro post en la BD */
type GraboPost struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
