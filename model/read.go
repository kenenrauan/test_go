package model

import (
	"fmt"
	"project/views"
)

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM employee")

	if err != nil {
		fmt.Println("Some error")
		return nil, err
	}
	employees := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Id, &data.Fullname, &data.Phone, &data.City)
		employees = append(employees, data)
	}

	return employees, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM employee WHERE fullname = $1", name)

	if err != nil {
		fmt.Println("Some error")
		return nil, err
	}

	employees := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Id, &data.Fullname, &data.Phone, &data.City)
		employees = append(employees, data)
	}

	return employees, nil

}
