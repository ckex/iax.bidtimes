package service

import (
	"fmt"
	"testing"
	"time"

	logger "github.com/alecthomas/log4go"
)

func init() {
	//输出到控制台,级别为DEBUG
	logger.AddFilter("stdout", logger.DEBUG, logger.NewConsoleLogWriter())
	//输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	logger.AddFilter("file", logger.DEBUG, logger.NewFileLogWriter("bidtimes.log", false))
	//logger.LoadConfiguration("log.xml")//使用加载配置文件,类似与java的log4j.propertites
	logger.Debug("start bidtimes server  .... ")
}

func Test_TimeAfter(t *testing.T) {
	defer logger.Close() //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件
	t1 := time.Now().Add(time.Hour * -20)
	t2 := time.Now()
	t3 := time.Now().Add(time.Hour * 20)
	logger.Info("\nt1= %s(%v) \nt2= %s(%v) \nt3= %s(%v)\n", format(t1), t1.Unix(), format(t2), t2.Unix(), format(t3), t3.Unix())

	b1 := t2.Unix() > t1.Unix()
	b2 := t2.Unix() > t3.Unix()
	fmt.Printf("t2 after t1(t2>t1)[%s>%s] = %v \n", format(t2), format(t1), b1) // t2 > t1
	fmt.Printf("t2 after t3(t2>t3)[%s>%s] = %v \n", format(t2), format(t3), b2)
	fmt.Printf("hahaha   %v",t2.Sub(t1))
}

func format(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
