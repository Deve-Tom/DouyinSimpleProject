package config

import (
	"DouyinSimpleProject/dao"
	"DouyinSimpleProject/entity"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB          *gorm.DB
	SERVER_HOST string
	SERVER_PORT string

	STATIC_ROOT_PATH       string
	DEFAULT_COVER_FILENAME string

	TIME_FORMAT = "2006-01-02 15:04:05"

	VIDEO_LIMIT int
)

func SetupViper() {
	// viper setting
	viper.SetConfigFile("conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to load config file")
	}
}

func LoadDB() *gorm.DB {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return db
}

func SetupDB() {
	db := LoadDB()
	db.AutoMigrate(&entity.User{}, &entity.Video{})
	// Combine `dao.Query` with db
	dao.SetDefault(db)
}

func Setup() {
	SetupViper()
	SetupDB()

	// get all *GLOBAL VARIABLES*
	SERVER_HOST = viper.GetString("server.host")
	SERVER_PORT = viper.GetString("server.port")

	DEFAULT_COVER_FILENAME = viper.GetString("static.default_cover_path")
	STATIC_ROOT_PATH = viper.GetString("static.root_path")

	VIDEO_LIMIT = viper.GetInt("video_limit")
}
