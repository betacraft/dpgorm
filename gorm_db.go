package dpgorm

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/jinzhu/gorm"
	"time"
)

type DbGorm struct {
	Other     *gorm.DB
	threshold int
}

func NewDB(m *gorm.DB) *DbGorm {

	return &DbGorm{
		Other:     m,
		threshold: 500,
	}
}

func (db *DbGorm) SetThreshold(ms int) {
	db.threshold = ms
}

func (db *DbGorm) logQuery(startTime time.Time, query string) {

	endTime := time.Now()
	t := int(((endTime.Sub(startTime)).Nanoseconds() / 1000000))

	ddb := deferstats.DeferDB{
		Query: query,
		Time:  t,
	}

	if t >= db.threshold {
		deferstats.Querylist.Add(ddb)
	}
}

func (db *DbGorm) Model(value interface{}) *gorm.DB {
	return db.Other.Model(value)
}

func (db *DbGorm) AutoMigrate(values ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", values))
	return db.Other.AutoMigrate(values...)
}

func (db *DbGorm) Create(value interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Create(value)
}

func (db *DbGorm) Save(value interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Save(value)
}

func (db *DbGorm) Delete(value interface{}, where ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Delete(value)
}

func (db *DbGorm) First(out interface{}, where ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", out))
	return db.Other.First(out, where...)
}

func (db *DbGorm) Last(out interface{}, where ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", out))
	return db.Other.Last(out, where...)
}

func (db *DbGorm) Count(value interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Count(value)
}

func (db *DbGorm) FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", out))
	return db.Other.FirstOrCreate(out, where...)
}

func (db *DbGorm) Find(out interface{}, where ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", out))
	return db.Other.Find(out, where...)
}

func (db *DbGorm) Where(query interface{}, args ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", query))
	return db.Other.Where(query, args...)
}

func (db *DbGorm) Or(query interface{}, args ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", query))
	return db.Other.Or(query, args...)
}

func (db *DbGorm) Not(query interface{}, args ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", query))
	return db.Other.Not(query, args...)
}

func (db *DbGorm) Limit(value interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Limit(value)
}

func (db *DbGorm) Offset(value interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", value))
	return db.Other.Offset(value)
}

func (db *DbGorm) Order(value string, reorder ...bool) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, value)
	return db.Other.Order(value, reorder...)
}

func (db *DbGorm) Select(query interface{}, args ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", query))
	return db.Other.Select(query, args...)
}

func (db *DbGorm) Omit(columns ...string) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, fmt.Sprintf("%#v", columns))
	return db.Other.Omit(columns...)
}

func (db *DbGorm) Group(query string) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, query)
	return db.Other.Group(query)
}

func (db *DbGorm) Having(query string, values ...interface{}) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, query)
	return db.Other.Having(query, values...)
}

func (db *DbGorm) Joins(query string) *gorm.DB {
	startTime := time.Now()
	defer db.logQuery(startTime, query)
	return db.Other.Joins(query)
}
