package service

import (
	"testing"
	"fmt"
	"iax.bidtimes/dao"
	"bytes"
	"strconv"
	"beego/cache"
)

func Test_Point(t *testing.T) {
	fmt.Println("haha---")
	array := pointArray()
	fmt.Printf("%T", array)
	i := pointInt()
	fmt.Println(i)
	array2 := pointArray2()
	fmt.Printf("%T", array2)

	saveData()

}

func saveData() {
	cache.FileCache{}
	redisConn := dao.RedisClient.Get()
	defer redisConn.Close()
	for i := 0; i< 1000000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString("BIDTIME[iax]1[iax]2016-05-31-10-00[iax]")
		buffer.WriteString(strconv.Itoa(i))
		redisConn.Do("set", buffer.String(), i)
	}
}

func pointArray() (*[]string) {
	var result []string
	for i := 0; i<10; i++ {
		result = append(result, "abcdefghijklmnopqrstuvwxyz")
	}
	return &result
}

func pointArray2() ([]string) {
	var result []string
	for i := 0; i<10; i++ {
		result = append(result, "abcdefghijklmnopqrstuvwxyz")
	}
	return result
}

func pointInt() *int {
	var a int = 1
	return &a
}