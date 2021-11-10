package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
	"github.com/streadway/amqp"
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

	conn, error_connec_mqtt := amqp.Dial("amqp://edwardlopez:servermqtt@165.227.69.88:8888/")
	if error_connec_mqtt != nil {
		defer conn.Close()
		log.Fatal(error_connec_mqtt)
	}

	ch, error_conection := conn.Channel()
	if error_conection != nil {
		defer ch.Close()
		log.Fatal(error_conection)
	}

	msgs, err_consume := ch.Consume("comensal/basicdata", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal(err_consume)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

}
