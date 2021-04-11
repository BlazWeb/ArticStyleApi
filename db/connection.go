package db

import (
	"database/sql"
	"fmt"
	"log"

	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Variables Entorno
var (
	db_name     = os.Getenv("DB_TABLE")
	db_user     = os.Getenv("DB_USER")
	db_password = os.Getenv("DB_PASSWORD")
	db_addr     = os.Getenv("DB_ADDR")
)

func checkParams() {
	if db_name == "" {
		fmt.Println("ArticDev -> Obligatorio establecer DB_TABLE, para table de la base de datos")
		time.Sleep(4 * time.Second)
		log.Fatal("ArticDev -> Cerrando programa....")
	}
	time.Sleep(1 * time.Second)
	if db_addr == "" {
		fmt.Println("ArticDev -> Se requiere que pongas una IP, por defecto le asigno: 127.0.0.1 añada una variable de entorno DB_ADDR")
	}
	time.Sleep(2 * time.Second)
	if db_user == "" {
		db_user = "root"
		fmt.Println("ArticDev -> Se recomienda que establezca un usuario para acceder a la conexión, variable de entrono DB_USER, por defecto le asigno: " + db_user)
	}
	time.Sleep(2 * time.Second)
	if db_password == "" {
		fmt.Println("ArticDev -> Se recomienda que establezca una contraseña para acceder a la conexión, variable de entrono DB_PASS")
	}
	time.Sleep(2 * time.Second)
}

func FirtConnection() (*sql.DB, error) {
	checkParams()
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	fmt.Println("ArticDev -> " + s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ArticDev -> Conexión exitosa con la base de datos")
	return db, nil
}

func GetDB() (*sql.DB, error) {
	if db_addr == "" {
		db_addr = "127.0.0.1"
	}
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_name)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)

	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
