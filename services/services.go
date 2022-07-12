package services

import (
	"log"

	"github.com/bimonugraraga/user-log-golang/collection"
	"github.com/bimonugraraga/user-log-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateNewUser(params map[string]interface{}) bool {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=coba_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                 // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	user := models.User{Username: params["username"].(string), Password: params["password"].(string)}
	result := db.Create(&user)

	if result.Error != nil {
		return false
	}
	return true
}

func GetAllUser() ([]models.User, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=coba_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                 // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	var user []models.User
	result := db.Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func Login(user collection.LoggedUser) (interface{}, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=coba_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                 // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	var targerUser models.User
	result := db.Where(models.User{Username: user.Username}).Find(&targerUser)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	if targerUser.Password != user.Password {
		return nil, err
	}

	var loggedIn collection.AToken
	var accesToken = targerUser.Username + targerUser.Password
	for i := len(accesToken) - 1; i >= 0; i-- {
		loggedIn.JWT += string(accesToken[i])
	}
	loggedIn.Username = targerUser.Username

	return loggedIn, nil
}
