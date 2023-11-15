package db

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"xorm.io/xorm/names"

	"github.com/Pingye007/godoing/config"
	"github.com/Pingye007/godoing/log"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

const (
	TableUser   = "gd_user"
	TableDoing  = "gd_doing"
	TableResult = "gd_result"
)

func connectDB() {
	settings := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&loc=%s",
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.ServerAddr,
		config.Cfg.DB.Port,
		config.Cfg.DB.DatabaseName,
		url.QueryEscape("Asia/Shanghai"))

	// Create a new engine which seals database/sql and all related functions
	eg, err := xorm.NewEngine(strings.ToLower(config.Cfg.DB.SqlType), settings)
	if err != nil {
		log.Log.Errorf("open database %s failed \n", config.Cfg.DB.DatabaseName)
		panic(err.Error())
	}

	// Set some configurations
	eg.SetMaxOpenConns(50)
	eg.SetMaxIdleConns(5)
	eg.SetConnMaxLifetime(time.Minute * 2)
	eg.SetConnMaxIdleTime(time.Minute)

	// Test connection of database
	err = eg.Ping()
	if err != nil {
		log.Log.Errorf("connect to database %s failed \n", config.Cfg.DB.DatabaseName)
		eg.Close()
		panic(err.Error())
	}

	eg.SetMapper(names.SnakeMapper{})

	log.Log.Infof("open database %s successful \n", config.Cfg.DB.DatabaseName)
	Engine = eg
}

func initTables(table ...string) {
	tableNum := len(table)
	if tableNum == 0 {
		return
	}
	tables := make([]any, 0)
	for _, t := range table {
		log.Log.Infoln("table:", t)
		if exist, err := Engine.IsTableExist(t); err != nil {
			log.Log.Errorf("check table %s existence failed \n", t)
			panic(err.Error())
		} else if !exist {
			switch t {
			case TableUser:
				tables = append(tables, new(User))
			case TableDoing:
				tables = append(tables, new(Doing))
			case TableResult:
				tables = append(tables, new(Result))
			}
		}
	}

	if len(tables) > 0 {
		log.Log.Infoln("need create tables")
		err := Engine.CreateTables(tables)
		if err != nil {
			log.Log.Errorln("create tables failed")
			panic(err.Error())
		}
	}
}

func init() {
	connectDB()
	initTables(TableUser, TableDoing, TableResult)
}
