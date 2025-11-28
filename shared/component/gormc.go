package sharedcomponent

import (
	"flag"

	"github.com/pkg/errors"
	sctx "github.com/viettranx/service-context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultMaxConnToDB     = 50
	DefaultMaxIdleConnToDB = 10
	DefaultMaxLifetimeToDB = 60 * 10 // in seconds
)

type GormComp struct {
	id              string
	dsn             string
	maxConnToDB     int
	maxIdleConnToDB int
	maxLifetimeToDB int // in seconds

	db *gorm.DB
}

func NewGormComp(id string) *GormComp {
	return &GormComp{id: id}
}

func (g *GormComp) ID() string {
	return g.id
}

func (g *GormComp) InitFlags() {
	flag.StringVar(&g.dsn, "db-dsn", "", "Main MySQLDatabase DSN")
	flag.IntVar(&g.maxConnToDB, "max-conn-to-db", DefaultMaxConnToDB, "Maximum number of connections to the database")
	flag.IntVar(&g.maxIdleConnToDB, "max-idle-conn-to-db", DefaultMaxIdleConnToDB, "Maximum number of idle connections to the database")
	flag.IntVar(&g.maxLifetimeToDB, "max-lifetime-to-db", DefaultMaxLifetimeToDB, "Maximum lifetime of a connection to the database in seconds")
}

func (g *GormComp) Activate(sctx sctx.ServiceContext) error {
	db, err := gorm.Open(mysql.Open(g.dsn), &gorm.Config{})

	if err != nil {
		return errors.WithStack(err)
	}

	if sctx.EnvName() != "production" {
		db = db.Debug()
	}

	g.db = db
	return nil
}

func (g *GormComp) Stop() error {
	// if g.db != nil {
	// 	sqlDB, _ := g.db.DB()
	// 	sqlDB.Close()
	// }

	return nil
}

func (g *GormComp) DB() *gorm.DB {
	return g.db.Session(&gorm.Session{NewDB: true})
}

type IGormComp interface {
	DB() *gorm.DB
}