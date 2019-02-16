package main

import "github.com/go-xorm/xorm"

func main() {
	engine, err := xorm.NewEngine("mysql", "")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

}