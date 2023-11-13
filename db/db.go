package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func connectDB() {
	settings := fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=utf8mb4,utf8&parseTime=true&loc=%s",
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.ServerAddr,
		config.Cfg.DB.Port,
		config.Cfg.DB.DatabaseName,
		"Asia/Shanghai")

	db, err := sql.Open(strings.ToLower(config.Cfg.DB.SqlType), settings)
	if err != nil {
		log.Log.Errorf("open database %s failed \n", config.Cfg.DB.DatabaseName)
		panic(err.Error())
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetConnMaxIdleTime(time.Minute)

	err = db.Ping()
	if err != nil {
		log.Log.Errorf("connect to database %s failed \n", config.Cfg.DB.DatabaseName)
		db.Close()
		panic(err.Error())
	}

	log.Log.Infof("open database %s successful \n", config.Cfg.DB.DatabaseName)
	DB = db
}

func init() {
	connectDB()
}
