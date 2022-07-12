package main

import (
	"net/http"

	"github.com/bimonugraraga/user-log-golang/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  "host=localhost user=postgres password=postgres dbname=coba_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
	// 	PreferSimpleProtocol: true,                                                                                                                 // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	// }), &gorm.Config{})

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	router := httprouter.New()
	router.POST("/register", controllers.Register)
	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}
	router.GET("/users", controllers.GetUsers)
	router.POST("/login", controllers.LoginUser)
	server.ListenAndServe()
}
