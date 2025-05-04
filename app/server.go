package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBuser     string
	DBPassword string
	DBname     string
	DbPort     string
	DBDriver   string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfing DBConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	server.initializeDB(dbConfing)
	server.InitializeRoutes()
}

func (server *Server) initializeDB(dbConfing DBConfig) {
	var err error
	if dbConfing.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfing.DBuser, dbConfing.DBPassword, dbConfing.DBHost, dbConfing.DbPort, dbConfing.DBname)

		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfing.DBHost, dbConfing.DBuser, dbConfing.DBPassword, dbConfing.DBname, dbConfing.DbPort)

		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("failed on connecting to the database server")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully.")

}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf(("Error On File Env"))
	}

	appConfig.AppName = getEnv("APP_NAME", "GoTokoApp")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBuser = getEnv("DB_USER", "andy")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBname = getEnv("DB_NAME", "db_toko")
	dbConfig.DbPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
