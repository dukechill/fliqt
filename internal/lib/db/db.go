package db

import (
	"database/sql"
	"fmt"
	"time"

	"fliqt/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func Init(cfg *config.Config) {
	var err error
	if cfg.Adapter == "mysql" {
		dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBProtocol,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
			cfg.DBParams,
		)

		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		DB.SetMaxOpenConns(cfg.DBMaxOpenConns)
		DB.SetMaxIdleConns(cfg.DBMaxOpenConns - 5)

		DB.SetConnMaxLifetime(30 * time.Second)

		readerDsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBProtocol,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
			cfg.DBParams)

		DBReader, err = sql.Open("mysql", readerDsn)
		// if replica connection failed, read master config
		if err != nil {
			DBReader, err = sql.Open("mysql", dsn)

			DBReader.SetMaxOpenConns(cfg.DBMaxOpenConns)
			DBReader.SetMaxIdleConns(cfg.DBMaxOpenConns - 5)
		} else {
			DBReader.SetMaxOpenConns(cfg.DBMaxOpenConns)
			DBReader.SetMaxIdleConns(cfg.DBMaxOpenConns - 5)
		}
		DBReader.SetConnMaxLifetime(30 * time.Second)
	}
	if err != nil {
		panic(err)
	}

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

	DBGorm, err = gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}))
	if err != nil {
		panic(err)
	}

	DBReaderGorm, err = gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}))
	if err != nil {
		panic(err)
	}
}
