package registro

import (

	//MODELS
	"log"

	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

var RegisterFrom_CS001 *registerFromCS001

type registerFromCS001 struct {
}

func (cr *registerFromCS001) RegisterBasicData(comensal models.Deserialized) {

	//Enviamos los datos al servicio
	error_r := RegisterBasicData_Service(comensal)
	if error_r != nil {
		log.Fatal(error_r)
	}
}
