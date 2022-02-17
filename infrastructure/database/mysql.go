package database

import (
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"kingsmith.com.cn/ks-horizon/config"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDb() error {
	conf := config.Conf.Mysql
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port, conf.DbName, conf.Charset)
	// db init
	db, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	boil.SetDB(db)
	return nil
}
