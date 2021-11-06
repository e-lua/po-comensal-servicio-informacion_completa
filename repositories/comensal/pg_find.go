package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

func Pg_Find_By_Id(input_phone int) (int, error) {

	var output_phone int

	db := models.Conectar_Pg_DB()
	q := "SELECT phone FROM Comensal WHERE phone=$1"
	error_showphone := db.QueryRow(q, input_phone).Scan(&output_phone)

	if error_showphone != nil {
		defer db.Close()
		return 0, error_showphone
	}

	defer db.Close()

	//Si todo esta bien
	return output_phone, nil
}
