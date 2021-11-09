package registro

import (

	//MDOELS
	"log"

	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"

	//REPOSITORIES
	comensal_repository "github.com/Aphofisis/po-comensal-servicio-informacion_completa/repositories/comensal"
)

func RegisterBasicData_Service(input_comensal models.Deserialized) error {

	var comensal models.Pg_Comensal

	//Validamos si esta registrado en el modelo Comensal
	/*idcomensal_founded, _ := comensal_repository.Pg_Find_By_Id(input_comensal.Phone)

	if idcomensal_founded > 0 {
		error_update := comensal_repository.Pg_Update_BasicData(input_comensal)
		if error_update != nil {
			return 500, true, "Error en el servidor interno al intentar actualizar la informacion del comensal, detalle: " + error_update.Error(), ""
		}
		return 200, false, "Información básica actualizada correctamente", ""
	}*/

	comensal.IdCountry = input_comensal.IdCountry
	comensal.Name = input_comensal.Name
	comensal.LastName = comensal.LastName
	comensal.CreatedDate = input_comensal.UpdatedDate
	comensal.UpdatedDate = input_comensal.UpdatedDate

	error_add := comensal_repository.Pg_Add_BasicData(comensal)

	if error_add != nil {
		log.Fatal(error_add)
	}

	return nil
}
