package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mmuflih/envgo/conf"
	"gitlab.com/aksestani-lib/core/app"
	"gopkg.in/mgo.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-01-31 09:25
**/

type LocationDB struct {
	*mgo.Database
}

func createConnection(cfg conf.Config, conn string) (*gorm.DB, error) {
	dbUser := cfg.GetString(conn + `.user`)
	dbPass := cfg.GetString(conn + `.pass`)
	dbName := cfg.GetString(conn + `.database`)
	dbHost := cfg.GetString(conn + `.address`)
	dbPort := cfg.GetString(conn + `.port`)
	query := cfg.GetString(conn + `.query`)

	if query == "" {
		query = "parseTime=true&loc=Asia%2FJakarta"
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?" + query
	app.Logger(dsn)

	/** create aksestani v2 database */
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		app.Logger("GORM Conn", err)
		return nil, err
	}

	if cfg.GetString("env") == "dev" || cfg.GetString("env") == "staging" {
		db = db.Debug()
	}
	return db, nil
}

func createMongoConnection(cfg conf.Config, conn string) (*mgo.Database, error) {
	address := cfg.GetString(conn + `.address`)
	user := cfg.GetString(conn + `.user`)
	pass := cfg.GetString(conn + `.pass`)
	database := cfg.GetString(conn + `.database`)
	port := cfg.GetString(conn + `.port`)
	auth := cfg.GetBool(conn + `.auth`)

	app.Logger(address, port, database)

	session, err := mgo.Dial(address + ":" + port)
	if err != nil {
		app.Logger("Mongodb Conn", err)
		return nil, err
	}
	dbSessionAksestani := session.DB(database)
	app.Logger("Mongo auth", auth)
	if auth {
		err := dbSessionAksestani.Login(user, pass)
		if err != nil {
			app.Logger("Mongodb Conn", err)
			return nil, err
		}
	}
	return dbSessionAksestani, nil
}

func NewLocationDatabase(cfg conf.Config, loc string) *LocationDB {
	var dbLoc *mgo.Database
	var err error

	dbLoc, err = createMongoConnection(cfg, loc)
	if err != nil {
		fmt.Println("Error creating connection [Location]", err)
	}
	return &LocationDB{dbLoc}
}
