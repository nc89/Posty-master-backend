package models

/*Post captura del Body, el mensaje que nos llega */
type Post struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
