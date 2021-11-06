package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

func Pg_Add_BasicData(comensalpg models.Pg_Comensal) error {

	db := models.Conectar_Pg_DB()
	q := "INSERT INTO Comensal_BasicData(idcountry,phone,password,updateddate) VALUES ($1,$2,$3,$4)"
	add_checker, err_add := db.Prepare(q)

	if err_add != nil {
		return err_add
	}

	add_checker.Exec(comensalpg.IdCountry, comensalpg.Phone, comensalpg.Password, comensalpg.UpdatedDate)
	return nil
}
