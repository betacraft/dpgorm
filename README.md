# Defer Panic gorm package

[![wercker status](https://app.wercker.com/status/36c6dbc709cd3f8b9318debc3b28f2f2/s "wercker status")](https://app.wercker.com/project/bykey/36c6dbc709cd3f8b9318debc3b28f2f2)

[![GoDoc](https://godoc.org/github.com/deferpanic/dpgorm?status.svg)](https://godoc.org/github.com/deferpanic/dpgorm)

### Defer Panic gorm add-on pkg

Many [deferpanic](https://deferpanic.com "deferpanic") users use [gorm](https://github.com/jinzhu/gorm "gorm").
This is the first attempt to start adding support. Much is left to do and this will probably change
dramatically in the near future. 

```
package main

import (
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/deferpanic/dpgorm"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name string
}

func main() {

	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")

	_db, err := gorm.Open("postgres", "dbname=gorm sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	db := dpgorm.NewDB(&_db)

	// set our threshold to log everything
	db.SetThreshold(-1)

	go dps.CaptureStats()

	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu"}
	blah := db.Create(&user)
	log.Println(blah)
	log.Println(db.First(&user))

	var count int
	db.Model(User{}).Where("name = ?", "Jinzhu").Count(&count)
	log.Println(count)

	time.Sleep(120 * time.Second)
}
```
