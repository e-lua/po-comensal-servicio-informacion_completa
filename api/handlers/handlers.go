package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	go Consumer()

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "3100"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consumer() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		defer ch.Close()
		log.Fatal("EEEEEEEEEEEEEEEEEEEEEERRRRRRRRRRRRRRRRRRRRRRR11111111111111111")
	}

	msgs, err_consume := ch.Consume("comensal/basicdata", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("EEEEEEEEEEEEEEEEEEEEEERRRRRRRRRRRRRRRRRRRRRRR2222222222222222222222")
	}

	noStop := make(chan bool)
	go func() {
		for d := range msgs {
			var comensal models.Deserialized
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&comensal)
			if err_consume != nil {
				log.Fatal("EEEEEEEEEEEEEEEEEEEEEERRRRRRRRRRRRRRRRRRRRRRR3333333333333")
			}
			log.Println("NNNNNNNNNNAAAAAAAMMMMME" + comensal.Name)

		}
	}()

	<-noStop
}
