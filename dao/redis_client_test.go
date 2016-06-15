package dao

import (
	"testing"
	"fmt"
)

// go test -test.bench="Benchmark_Client" -test.benchmem -test.count 5 -test.cpu 4
func Benchmark_Client(b *testing.B) {
	for i := 0; i<b.N; i++ {
		conn := RedisClient.Get()
		defer conn.Close()
		_, err := conn.Do("get", "abc")
		if err != nil {
			continue
		}
	}
}

func Benchmark_Client2(b *testing.B) {
	for i := 0; i<b.N; i++ {
		_, err := RedisClient.Get().Do("get", "abc")
		if err != nil {
			b.Error(err)
		}
	}
}

// go test -v
func Test_Client(t *testing.T) {
	conn := RedisClient.Get()
	defer conn.Close()
	ret, err := conn.Do("keys", "*")
	if err !=nil {
		t.Error(err)
	}

	if instence, ok := ret.([]interface{}); ok {
		for _, value := range instence {
			if b, bok := value.([]uint8); bok {
				key := string(b)
				times, err := conn.Do("get", key)
				if err != nil {
					t.Error(err)
				}
				if timesvalue, timeok := times.([]byte); timeok {
					fmt.Println(key, string(timesvalue))
				}
			}
		}
	}
	switch ret.(type) {      //多选语句switch
	case string: //是字符时做的事情
	case int://是整数时做的事情
	}

}