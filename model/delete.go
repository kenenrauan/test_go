package model

import "fmt"

func DeleteEmployee(name string) error {
	query := "DELETE FROM employee WHERE fullname = $1"
	_, err := con.Exec(query, name)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
