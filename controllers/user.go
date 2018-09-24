package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DaveSalazar/ed-comments/commons"
	"github.com/DaveSalazar/ed-comments/configuration"

	"github.com/DaveSalazar/ed-comments/models"
)

//Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error, %s\n", err)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)

	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})

		if err != nil {
			log.Fatalf("error al convertir el token a json %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "usuario o clave incorrecto",
			Code:    http.StatusUnauthorized,
		}

		commons.DisplayMessage(w, m)
	}

}

//UserCreate permite registrar un usuario
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		m.Message = fmt.Sprintf("error al leer el usuario a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = "Las contrasenas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	user.Picture = "https://gravatar.com/avatar/" + picstr + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error

	if err != nil {
		m.Message = fmt.Sprintf("error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Usuario creado con exito"
	m.Code = http.StatusCreated

	commons.DisplayMessage(w, m)
}
