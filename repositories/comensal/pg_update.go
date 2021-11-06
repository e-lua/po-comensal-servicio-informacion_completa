package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

func Pg_Update_BasicData(comensal_basicdata models.Pg_Comensal) error {

	db := models.Conectar_Pg_DB()
	q := "UPDATE Comensal SET idcountry=$1,password=$2,updateddate=$3 WHERE phone=$4"
	updateBusiness_Name, error_update := db.Prepare(q)

	//Instanciamos una variable del modelo R_Country
	if error_update != nil {
		defer db.Close()
		return error_update
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	updateBusiness_Name.Exec(comensal_basicdata.IdCountry, comensal_basicdata.Password, comensal_basicdata.Password, comensal_basicdata.UpdatedDate, comensal_basicdata.Phone)

	defer db.Close()

	//Si todo esta bien
	return nil

}
