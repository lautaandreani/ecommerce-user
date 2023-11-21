package db

import (
	"fmt"
	"home/lautaro/dev/ecommerce-user/models"
	"home/lautaro/dev/ecommerce-user/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Starting registry in DB")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "', '" + tools.MySQLDate() + "')"
	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Sign Up successfully")
	return nil
}
