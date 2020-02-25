package etctv3conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
)

type RedisConfig struct {
	Address string
	Port int
}

type Etcdv3Conf struct {
	Appname       string `json:"appname"`
	Httpaddr      string `json:"httpaddr"`
	Httpport      int    `json:"httpport"`
	Imgaddr       string `json:"imgaddr"`
	Mysqladdr     string `json:"mysqladdr"`
	Mysqldbname   string `json:"mysqldbname"`
	Mysqlport     int    `json:"mysqlport"`
	Redisaddr     string `json:"redisaddr"`
	Redisdbnum    int    `json:"redisdbnum"`
	Redispassword string `json:"redispassword"`
	Redisport     int    `json:"redisport"`
}

type Etcdv3ConfManager struct {

}

func (et *Etcdv3ConfManager)GetEtcdv3Conf(address string,prefix string)(res Etcdv3Conf,err error){
	etcdSource := etcd.NewSource(
		// "127.0.0.1:2379"
		etcd.WithAddress(address),
		etcd.StripPrefix(true),
		//etcd.WithPrefix("/a"),
		etcd.WithPrefix(prefix),
	)
	conf, err := config.NewConfig()
	if err!=nil{
		return res,err
	}
	conf.Load(etcdSource)

	//err=config.Get("address").Scan(&redisConfig2)
	err=conf.Scan(&res)
	if err!=nil{
		return res,err
	}
	return res,nil
}

