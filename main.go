package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Tweet struct {
	gorm.Model
	Content string
}

// DB接続の処理
func dbInit() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&Tweet{})
}

//　DBにインサート
func dbInsert(content string) error {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(&Tweet{Content: content})
	return nil
}

//　情報を取得
func GetAll() ([]Tweet, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tweets []Tweet
	db.Order("created_at desc").Find(&tweets)
	return tweets, nil

}

func main() {
	e := echo.New()

	dbInit()

	e.GET("/", func(c echo.Context) error {
		tweets, err := GetAll()
		if err != nil {
			return err
		}
		c.HTML(200, tweets[0].Content)
		return nil
	})

	e.POST("/new", func(c echo.Context) error {
		content := c.FormValue("content")
		dbInsert(content)
		c.Redirect(200, "/")
		return nil
	})

	e.Start(":8080")

}
