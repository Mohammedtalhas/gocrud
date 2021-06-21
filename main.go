//main.go
package main

import (
	"fmt"
	"os"

	"github.com/Mohammedtalhas/gocrud/Config"
	"github.com/Mohammedtalhas/gocrud/Models"
	"github.com/Mohammedtalhas/gocrud/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run(os.Getenv("PORT"))
}
