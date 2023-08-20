package ent

import (
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"todo/ent"
)

func Open() (*ent.Client, error) {
	entOptions := []ent.Option{}

	// 発行されるSQLをロギング
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST") + ":" + "3306",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
		return nil, err
	}
	return client, nil
}
