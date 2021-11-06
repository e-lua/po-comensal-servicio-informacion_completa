package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-informacion_completa/models"
)

func Pg_Add_BasicData(comensalpg models.Pg_Comensal) error {

	db := models.Conectar_Pg_DB()
	q := "INSERT INTO Comensal_BasicData(idcountry,name,lastname, phone,password,createddate,updateddate) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	add_checker, err_add := db.Prepare(q)

	if err_add != nil {
		return err_add
	}

	add_checker.Exec(comensalpg.IdCountry, comensalpg.Name, comensalpg.LastName, comensalpg.Phone, comensalpg.Password, comensalpg.UpdatedDate, comensalpg.UpdatedDate)
	return nil
}
