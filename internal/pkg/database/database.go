package database

import (
	"time"

	"moul.io/zapgorm2"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Options is  configuration of database
type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

//NewOptions 获取配置信息
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

//New Init 初始化数据库
func New(o *Options, logger *zap.Logger) (*gorm.DB, error) {
	var err error

	gormLog := zapgorm2.New(logger)
	gormLog.SetAsDefault()

	dbConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: gormLog,
		NowFunc: func() time.Time {
			return time.Unix(time.Now().UTC().Unix(), 0)
		},
	}

	//db, err := gorm.Open(mysql.Open(o.URL), dbConfig)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: o.URL, DisableDatetimePrecision: true}), dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}

	if o.Debug {
		db = db.Debug()
	}

	//连接池设置
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)           //SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(50)           //SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) //SetConnMaxLifetime 设置了连接可复用的最大时间

	return db, nil
}

//ProviderSet wire注册器
var ProviderSet = wire.NewSet(New, NewOptions)
