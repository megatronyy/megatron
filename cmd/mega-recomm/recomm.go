package main

import "github.com/go-xorm/xorm"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	conn = "root:QU90()op@tcp(192.168.0.106:3306)/pholcus?charset=utf8"
)

func main() {
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	results, err := engine.Query("SELECT * FROM business__house_detail ORDER BY 1 DESC LIMIT 10")
	if err != nil {
		panic(err)
	}

	for _, row := range results {
		fmt.Println(row)
	}
}
