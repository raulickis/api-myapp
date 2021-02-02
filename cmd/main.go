package main

import (
	"fmt"
	"github.com/autorei/api-myapp/api"
	"github.com/autorei/api-myapp/config"
	"github.com/autorei/api-myapp/internal/database"
	"github.com/autorei/api-myapp/internal/database/migrations"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	printEnvVars()

	db := &database.Repository{}
	err := db.Start()

	if err != nil {
		fmt.Printf("ERRO DURANTE INICIALIZAÇÃO DO BANCO DE DADOS \n")
		panic(err.Error())
	}

	migrations.RunMigrations()

	api.Run()

}

func printEnvVars() {
	fmt.Printf("VARIAVEIS DE AMBIENTE \n")
	fmt.Printf("ENV              : %s \n", config.ENV)
	fmt.Printf("APP_PORT         : %s \n", config.APP_PORT)
	fmt.Printf("PORT_EXTERNAL    : %s \n", config.PORT_EXTERNAL)
	fmt.Printf("PORT_INTERNAL    : %s \n", config.PORT_INTERNAL)
	fmt.Printf("DOCKER_HOST_IP   : %s \n", config.DOCKER_HOST_IP)
	fmt.Printf("DB_HOST          : %s \n", config.DB_HOST )
	fmt.Printf("DB_NAME          : %s \n", config.DB_NAME )
	fmt.Printf("DB_USER          : %s \n", config.DB_USER )
	fmt.Printf("DB_PASS          : %s \n", config.DB_PASS )
	fmt.Printf("DB_PORT          : %v \n", config.DB_PORT )
}