package dpgorm

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	selectThreshold int
)

type DbGorm struct {
	Other *gorm.DB
}

func NewDB(m *gorm.DB) *DbGorm {
	selectThreshold = 500

	return &DbGorm{
		m,
	}
}

func (db *DbGorm) logQuery(startTime time.Time, query string) {

	endTime := time.Now()
	t := int(((endTime.Sub(startTime)).Nanoseconds() / 1000000))

	ddb := deferstats.DeferDB{
		Query: query,
		Time:  t,
	}

	if t >= selectThreshold {
		deferstats.Querylist.Add(ddb)
	}
}

func (db *DbGorm) AutoMigrate(values ...interface{}) *DbGorm {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", values))
	db.Other.AutoMigrate(values...)
	return db
}

func (db *DbGorm) Create(value interface{}) *DbGorm {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	db.Other.Create(value)
	return db
}

func (db *DbGorm) First(out interface{}, where ...interface{}) *DbGorm {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", out))
	db.Other.First(out, where...)
	return db
}
