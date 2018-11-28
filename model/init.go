package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// MySQL driver.
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var DB *Database
var RC *RedisConn

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

type RedisConn struct {
	Self redis.Conn
}

func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func (rc *RedisConn) Init() error {
	getRc, err := GetSelfRedis()
	if err != nil {
		return err
	}
	RC = &RedisConn{
		Self: getRc,
	}
	return nil
}

func (rc *RedisConn) Close() {
	RC.Self.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// 建立数据库连接
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// 用于设置最大打开的连接数，默认为0不限制，设置最大的连接数可以避免并发太高
	// 导致连接mysql出现 too many connections 的错误
	//db.DB().SetMaxOpenConns(20000)
	// 用于设置闲置的连接数，设置闲置连接数则当开启的一个连接使用完成后可以放在连接池里等待下次使用
	db.DB().SetMaxIdleConns(0)
}

// 使用了cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func InitSelfRedis() (redis.Conn, error) {
	rc, err := redis.Dial(viper.GetString("redis.network"), viper.GetString("redis.addr"))
	return rc, err
}

func GetSelfRedis() (redis.Conn, error) {
	return InitSelfRedis()
}

func (rc *RedisConn) SetKeyInRc(key, timeout string, value interface{}) error {
	// 对本次连接进行set操作
	// EX单位为秒
	_, setErr := RC.Self.Do("set", key, value, "EX", timeout)
	return setErr
}

func (rc *RedisConn) GetKeyInRc(key string) (interface{}, error) {
	// 使用redis的string类型获取set的k/v信息
	val, getErr := redis.String(RC.Self.Do("get", key))
	return val, getErr
}

func (rc *RedisConn) DelKeyInRc(key string) error {
	_, delerr := RC.Self.Do("del", key)
	return delerr
}
