package etctv3conf
//package main

import (
	"log"
	"testing"
)


func TestGetEtcdv3Conf(t *testing.T){

	log.Println("start")
	e:=Etcdv3ConfManager{}
	res,err:=e.GetEtcdv3Conf("127.0.0.1:2379","/micro/config")
	if err!=nil{
		log.Printf("== %+v",err)
	}
	
	log.Println(res)

}


/*
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


func main(){
	fmt.Println("start")

	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.StripPrefix(true),
		etcd.WithPrefix("/micro/config"),
		//etcd.WithPrefix("/a"),
	)

	conf, err := config.NewConfig()
	if err!=nil{
		log.Printf("err:%+v",err)
	}
	conf.Load(etcdSource)

	//var redisConfig RedisConfig
	//res:=conf.Get("micro", "config", "database")
	//fmt.Println("conf.res:", res)
	//fmt.Println("conf.res:", res.Scan(redisConfig))

	//if err := conf.Scan(&redisConfig); err != nil {
	//	fmt.Println("conf.Scan:", err)
	//}
	//fmt.Println("conf.Scan:", redisConfig)


	//var redisConfig2 RedisConfig
	//res:=conf.Get("address").Bytes()
	//
	//log.Println(string(res))


	var c Etcdv3Conf
	//err=config.Get("address").Scan(&redisConfig2)
	err=conf.Scan(&c)
	if err!=nil{
		log.Println(err)
	}
	log.Printf("===== %+v",c)

	// 10.0.0.1 3306
	//fmt.Println(redisConfig2.Address, redisConfig2.Port)


}

*/

