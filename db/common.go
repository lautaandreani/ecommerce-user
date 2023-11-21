package db

import (
	"database/sql"
	"fmt"
	"os"

	"home/lautaro/dev/ecommerce-user/models"
	"home/lautaro/dev/ecommerce-user/secretm"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecrets(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", getConnectionString(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connection successfully")
	return nil
}

func getConnectionString(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "ecommercedb"

	formatString := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

	return formatString
}
