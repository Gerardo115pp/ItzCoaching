package database

import (
	"database/sql"
	"fmt"
	app_config "itz_JD_service/Config"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func createDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", app_config.MYSQL_USER, app_config.MYSQL_PASSWORD, app_config.MYSQL_HOST, app_config.MYSQL_PORT, app_config.MYSQL_DB)
}
