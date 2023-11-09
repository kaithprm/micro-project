package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var db *sql.DB
var dns = "root:123456@tcp(127.0.0.1:3306)/test"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("D:\\Develop\\gopath\\micro\\login\\html\\login.tmpl") //模板解析
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": "登录页面"}) //模板渲染
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	Mysql()
}
func Mysql() (err error) {
	db, err = sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
