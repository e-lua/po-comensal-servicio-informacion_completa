package registro

import (
	"github.com/labstack/echo/v4"

	//MDOELS
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

var RegisterFrom_CS001 *registerFromCS001

type registerFromCS001 struct {
}

func (cr *registerFromCS001) RegisterBasicData(c echo.Context) error {

	//Instanciamos una variable del modelo Code
	var comensal models.Pg_Comensal

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&comensal)
	if err != nil {
		results := Response_WithString{Error: true, DataError: "Se debe enviar los datos del pais, nombre, apellido y contraseña del comensal, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Validamos los valores enviados
	if comensal.Phone < 999999 && len(comensal.Name) < 1 && len(comensal.LastName) < 1 {
		results := Response_WithString{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := RegisterBasicData_Service(comensal)
	results := Response_WithString{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
