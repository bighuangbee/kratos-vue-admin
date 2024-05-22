package main

import (
	"flag"
	"github.com/bighuangbee/kratos-vue-admin/app/admin/internal/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
)

var flagconf1 string

func init() {
	flag.StringVar(&flagconf1, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func main() {

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf1),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "../internal/data/dal/query",
		ModelPkgPath:      "../internal/data/dal/model",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable:     false,                                                              // generate pointer when field is nullable
		FieldCoverable:    false,                                                              // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		FieldSignable:     true,                                                               // detect integer field's unsigned type, adjust generated data type
		FieldWithIndexTag: false,                                                              // generate with gorm index tag
		FieldWithTypeTag:  true,
	})

	gormdb := NewDB(bc.Data, log.NewStdLogger(os.Stdout))
	g.UseDB(gormdb)

	g.GenerateAllTable()

	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}

func NewDB(config *conf.Data, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "receive-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(config.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		logs.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logs.Fatalf("failed get sql.DB : %v", err)
	}
	sqlDB.SetMaxIdleConns(int(config.Database.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(config.Database.MaxOpenConns))

	return db
}
