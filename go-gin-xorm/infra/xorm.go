package infra

import (
	"log"

	"github.com/daisuke23bubu/go-gin-xorm/model"
	"github.com/go-xorm/xorm"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp([127.0.0.1]:3306)/go_gin_xorm_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// xorm で利用した SQL をログに出力
	engine.ShowSQL(true)

	exist, err := engine.IsTableExist("users")
	if err != nil {
		log.Fatal(err)
	}

	if !exist {
		engine.CreateTables(&model.Users{})
	}

	return engine
}
