package db

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password recibida  y la encripta 2^8 veces*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
