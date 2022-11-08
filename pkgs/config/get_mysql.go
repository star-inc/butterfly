// Butterfly - The web crawler base on Apache Solr for StarStart!
// Copyright(c) 2020 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func GetMySQL() *sql.DB {
	mysql, err := sql.Open("mysql", MysqlDSN)
	if err != nil {
		log.Panicln(err)
	}
	mysql.SetConnMaxLifetime(time.Minute * 3)
	mysql.SetMaxOpenConns(10)
	mysql.SetMaxIdleConns(10)
	return mysql
}
