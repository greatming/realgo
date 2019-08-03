package db

import (
	"sync"
	"realgo/lib/logger"
	"fmt"
)

type RealDB struct {
	m sync.Map
}


var RDB = RealDB{}

func (r *RealDB)GetCluster(dbname string) (*DBConnInfo, error) {
	tdb, ok :=  r.m.Load(dbname)
	if ok{
		db,err := tdb.(*DBHandler).GetInstance()
		if err != nil{
			return nil, err
		}
		return db,nil
	}
	return nil, fmt.Errorf("get cluster fail")
}

func (r *RealDB)SetCluster(dbname string, DBCfg *DBItemConf) bool {
	_, ok :=  r.m.Load(dbname)
	if ok{
		return true
	}

	log := logger.New()
	tdb := NewDBHandler(DBCfg.Pool, DBCfg.Info, log)
	fmt.Println(tdb.GetInstance())
	r.m.Store(dbname, tdb)
	return true
}





