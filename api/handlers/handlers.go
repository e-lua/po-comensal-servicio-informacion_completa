package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	register "github.com/Aphofisis/po-comensal-servicio-informacion_completa/services/flujo_de_informacion/registro_basico_de_CS001"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	//V1 TO ENTITY-CODE
	router_comensal_v1 := version_1.Group("/comensal/basicdata")
	router_comensal_v1.PUT("", register.RegisterFrom_CS001.RegisterBasicData)

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
