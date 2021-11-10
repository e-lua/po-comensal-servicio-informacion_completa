package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

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
		log.Fatal(error_conection)
	}

	defer ch.Close()

	msgs, err_consume := ch.Consume("comensal/basicdata", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal(err_consume)
	}

	noStop := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	<-noStop

}
