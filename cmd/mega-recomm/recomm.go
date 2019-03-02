package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/twfx7758/megatron/cmd/mega-recomm/sqlcommon"
)

func main() {
	config := sqlcommon.NewConfig()
	defer config.Engine.Close()

	results, err := config.Engine.Query(sqlcommon.Detail)
	if err != nil {
		panic(err)
	}

	for _, row := range results {
		fmt.Println(row)
	}
}
