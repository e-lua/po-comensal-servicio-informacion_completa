package registro

import (

	//MDOELS
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"

	//REPOSITORIES
	comensal_repository "github.com/Aphofisis/po-comensal-servicio-informacion_completa/repositories/comensal"
)

func RegisterBasicData_Service(input_comensal models.Pg_Comensal) (int, bool, string, string) {

	//Validamos si esta registrado en el modelo Comensal
	idcomensal_founded, _ := comensal_repository.Pg_Find_By_Id(input_comensal.Phone)

	if idcomensal_founded > 0 {
		error_update := comensal_repository.Pg_Update_BasicData(input_comensal)
		if error_update != nil {
			return 500, true, "Error en el servidor interno al intentar actualizar la informacion del comensal, detalle: " + error_update.Error(), ""
		}
		return 200, false, "Información básica actualizada correctamente", ""
	}

	error_add := comensal_repository.Pg_Add_BasicData(input_comensal)

	if error_add != nil {
		return 500, true, "Error en el servidor interno al agregar la información del comensal comensal, detalle: " + error_add.Error(), ""
	}

	return 200, false, "", "Información básica actualizada correctamente"
}
