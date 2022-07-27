package main

import (
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal"
	"github.com/alandtsang/go-leaderelection/datasource/mysql/dal/model"
	"gorm.io/gen"
)

func init() {
	dsn := "root:11223344@tcp(127.0.0.1:3306)/oort?charset=utf8&parseTime=True"
	dal.DB = dal.ConnectDB(dsn).Debug()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./datasource/mysql/dal/query",
		Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(dal.DB)
	g.ApplyBasic(model.LeaderInfo{})
	g.Execute()
}
