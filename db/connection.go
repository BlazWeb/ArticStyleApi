package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Variables Entorno
var _ = godotenv.Load(".env")
var (
	db_name     = os.Getenv("DB_TABLE")
	db_user     = os.Getenv("DB_USER")
	db_password = os.Getenv("DB_PASSWORD")
	db_addr     = os.Getenv("DB_ADDR")
)

func checkParams() {
	if db_name == "" {
		fmt.Println("ArticDev -> Es necesario introducir la variable de entorno DB_TABLE:")
		var dato_db_name string
		fmt.Scanln(&dato_db_name)

		if dato_db_name != "" {
			db_name = dato_db_name
			os.Setenv("DB_TABLE", dato_db_name)
		}
	}
	time.Sleep(1 * time.Second)
	if db_addr == "" {
		db_addr = "127.0.0.1"
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

func FirtConnection() error {
	checkParams()
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_addr, db_name)
	db, err := sql.Open("mysql", s)
	fmt.Println("ArticDev -> " + s)
	if err != nil {
		return err
	}
	_, err = db.Query("SELECT * FROM information_schema.tables WHERE table_schema = ?", db_name)
	if err != nil {
		return err
	}
	fmt.Println("ArticDev -> Conexion exitosa con la base de datos")
	return nil
}

func GetDB() (*sql.DB, error) {
	if db_addr == "" {
		db_addr = "127.0.0.1"
	}
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?&parseTime=true", db_user, db_password, db_addr, db_name)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)

	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
