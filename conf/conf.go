package conf

import (
	"io/ioutil"
	"realgo/lib/toml"
	_"sync"
	"sync"
	"fmt"
	"reflect"
)

var confCache sync.Map


// 读取配置文件
func ReadFile(confPath string, config interface{})(err error){

	var bs []byte
	var realFilePath string = FileAbsPath(confPath)
	var cacheKey string = getConfCacheKey(realFilePath)
	vs, ok := confCache.Load(getConfCacheKey(cacheKey))

	if !ok {
		bs, err = ioutil.ReadFile(realFilePath)
		if err != nil{
			return  err
		}
		var confStr = string(bs)
		_, err = toml.Decode(confStr, config)
		if err != nil{
			return  err
		}
		confCache.Store(cacheKey, reflect.ValueOf(config).Elem().Interface())
		return
	}

	reflect.ValueOf(config).Elem().Set(reflect.ValueOf(vs))
	return nil

}

//读取app配置文件
func ReadAppConfFile(confPath string, config interface{})(err error)  {
	confPath = "./conf/" + confPath
	return ReadFile(confPath, config)
}

func getConfCacheKey(confPath string) string {
	return fmt.Sprintf("%s", confPath)
}