package model

import "fmt"

func Create(fullname, phone, city string) error {
	query := "INSERT INTO employee (fullname, phone, city) VALUES ($1, $2, $3)"

	_, err := con.Exec(query, fullname, phone, city)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Данные успешно добавлены в таблицу employee!")
	return nil

}
