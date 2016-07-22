package main

import (
	"fmt"
	"time"

	logger "github.com/alecthomas/log4go"
	"github.com/astaxie/beego"
	_ "iax.bidtimes/dao"
	_ "iax.bidtimes/routers"
	"iax.bidtimes/service"
)

//func init() {
//	//输出到控制台,级别为DEBUG
//	logger.AddFilter("stdout", logger.DEBUG, logger.NewConsoleLogWriter())
//	//输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
//	logger.AddFilter("file", logger.DEBUG, logger.NewFileLogWriter("bidtimes.log", false))
//	//logger.LoadConfiguration("log.xml")//使用加载配置文件,类似与java的log4j.propertites
//	logger.Debug("start bidtimes server  .... ")
//	// defer logger.Close() //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件
//}

func init() {
	logger.LoadConfiguration("./conf/log4go.xml")
}


func main() {
	startTask()
	beego.Run()
}

func startTask() {
	deftime := 5 * 60
	intervalTime := beego.AppConfig.DefaultInt("intervalTime", deftime)
	ticker := time.NewTicker(time.Second * time.Duration(intervalTime))
	go func() {
		for _ = range ticker.C {
			ret := call()
			logger.Info("bid response time task %v %v ", ret, time.Now())
		}
	}()
}

func call() (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error: %v \n", r)
			ret = false
		} else {
			ret = true
		}
	}()
	service.Start()
	return
}
