package db

import (
	"time"
	"fmt"
	"math/rand"
	_"github.com/go-sql-driver/mysql"
	"github.com/greatming/realgo/lib/logger"
	"github.com/jinzhu/gorm"
)


type DBPoolConf struct {
	MaxOpenConn  int `toml:"MaxOpenConn"`
	MaxIdleConn  int `toml:"MaxIdleConn"`
	MaxLifeTime  int `toml:"MaxLifeTime"`
	ReadTimeout  int `toml:"ReadTimeout"`
	WriteTimeout int `toml:"WriteTimeout"`
}

type DBInfoConf struct {
	Host   []string `toml:"host"`
	User   string   `toml:"user"`
	Pwd    string   `toml:"pwd"`
	DBName string   `toml:"dbname"`
}

type DBItemConf struct {
	Pool DBPoolConf `toml:"pool"`
	Info DBInfoConf `toml:"info"`
}

type DBHandler struct {
	Info    DBInfoConf
	PoolCfg DBPoolConf
	Handler []*DBConnInfo
	Logger *logger.Logger
}




type DBConnInfo struct {
	DB          *gorm.DB
	Logger  *logger.Logger
}

func (db *DBHandler) getHandler(user, pwd, host, dbname string) (*DBConnInfo, error) {
	str := fmt.Sprintf("%s:%s@tcp(%s)/%s?readTimeout=%dms&writeTimeout=%dms&parseTime=True&charset=utf8",
		user, pwd, host, dbname, db.PoolCfg.ReadTimeout, db.PoolCfg.WriteTimeout)
	handler, err := gorm.Open("mysql", str)
	if err != nil {
		db.Logger.Warn("Open " + str + " failed, err: " + err.Error())
		fmt.Println("Open " + str + " failed, err: " + err.Error())
		return nil, err
	}
	dbhandler := handler.DB()
	dbhandler.SetMaxOpenConns(db.PoolCfg.MaxOpenConn)
	dbhandler.SetMaxIdleConns(db.PoolCfg.MaxIdleConn)
	dbhandler.SetConnMaxLifetime(time.Duration(db.PoolCfg.MaxLifeTime) * time.Second)


	db.Logger.Info4("Open DB: %s with pool [MaxOpenConns:%d] [MaxIdleConns:%d] [MaxLifeTime:%d]",
		str, db.PoolCfg.MaxOpenConn, db.PoolCfg.MaxIdleConn, db.PoolCfg.MaxLifeTime)
	//测试是否能连上DB
	if err := dbhandler.Ping(); err != nil {
		return nil, err
	}
	return &DBConnInfo{DB:handler, Logger:db.Logger}, err
}
func (db *DBHandler) GetInstance() (*DBConnInfo, error) {
	if len(db.Handler) == 0 {
		return nil, fmt.Errorf("no db handler, maybe forgot init DBHandler")
	}
	size := len(db.Handler)
	index := rand.Intn(size)
	var err error
	for i := 0; i < size; i++ {
		if db.Handler[index] != nil {
			return db.Handler[index], nil
		}
		//db Handle init failed at first, init Handle
		db.Handler[index], err = db.getHandler(db.Info.User, db.Info.Pwd, db.Info.Host[index], db.Info.DBName)
		if err != nil {
			index = (index + 1) % size
		} else {
			return db.Handler[index], nil
		}
	}
	return nil, fmt.Errorf("can not find any host connect success")
}
func NewDBHandler(poolCfg DBPoolConf, info DBInfoConf, log *logger.Logger) *DBHandler {
	db := new(DBHandler)
	db.Info = info
	db.PoolCfg = poolCfg
	db.Logger = log
	db.Handler = make([]*DBConnInfo, len(db.Info.Host))
	for i := range db.Handler {
		db.Handler[i], _ = db.getHandler(db.Info.User, db.Info.Pwd, db.Info.Host[i], db.Info.DBName)
	}
	return db
}
