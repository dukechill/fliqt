package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DB writer db instance
	DB *sql.DB
	// DBReader only reader db instance
	DBReader *sql.DB
	// DBGorm only gorm use
	DBGorm *gorm.DB
	// DBReaderGorm only reader gorm use
	DBReaderGorm *gorm.DB
)

func Init() {
	if DB == nil {
		panic("DB nil")
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	initGorm()
}

func initGorm() {
	var err error

	DBGorm, err = gorm.Open("mysql", DB)
	if err != nil {
		panic(err)
	}

	DBReaderGorm, err = gorm.Open("mysql", DBReader)
	if err != nil {
		panic(err)
	}
}
