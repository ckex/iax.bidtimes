package service

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"iax.bidtimes/dao"
	"iax.bidtimes/models"
	"time"
)

func Start() {

	key := "BIDTIME*"
	keys, err := listAllKey(key)
	if err != nil {
		fmt.Printf("listAllKey error: %v \n", err)
		return
	}

	validKey := filterKeys(keys)

	fmt.Printf("result size %d, after filter size %d \n", len(keys), len(validKey))

	group := make(map[groupKey][]bidtimes)
	for _, value := range validKey {
		if v, ok := group[value.groupKey]; ok {
			v = append(v, value)
			group[value.groupKey] = v
		} else {
			var gvalue []bidtimes
			group[value.groupKey] = append(gvalue, value)
		}
	}

	for key, value := range group {

		count, useTimes := caclGroup(value)

		countfloat := float64(count)
		percent99 := int(countfloat * float64(0.99)) // 99%
		percent85 := int(countfloat * float64(0.85)) // 85%
		percent55 := int(countfloat * float64(0.55)) // 55%

		var cacl = count
		var index99 = -1
		var index85 = -1
		var index55 = -1

		for i := len(useTimes) - 1; i >= 0; i-- {
			cacl = cacl - useTimes[i]
			if index55 < 0 && cacl < percent55 {
				index55 = i + 1
				if index85 < 0 {
					index85 = index55
				}
				if index99 < 0 {
					index99 = index55
				}
				break
			} else if index85 < 0 && cacl < percent85 {
				index85 = i + 1
				if index99 < 0 {
					index99 = index85
				}
			} else if index99 < 0 && cacl < percent99 {
				index99 = i + 1
			}
		}

		if index55 > 0 && index85 > 0 && index99 > 0 {
			var id int64
			if key.isTest {
				iaxTestResponseTimes := models.IaxTestResponseTimes{USERID: int64(value[0].dspID), TIME: value[0].theTime, FIFTYTH: index55, EIGHTYFIVETH: index85, NINETYNINETH: index99, DATETIMECREATE: time.Now(), DATETIMEMODIFIED: time.Now()}
				id, err = models.AddIaxTestResponseTimes(&iaxTestResponseTimes)
			} else {
				iaxResponseTimes := models.IaxResponseTimes{USERID: int64(value[0].dspID), TIME: value[0].theTime, FIFTYTH: index55, EIGHTYFIVETH: index85, NINETYNINETH: index99, DATETIMECREATE: time.Now(), DATETIMEMODIFIED: time.Now()}
				id, err = models.AddIaxResponseTimes(&iaxResponseTimes)
				// if err != nil {
				// 	fmt.Println("insert iax_response_times error ", err)
				// 	continue
				// }
				// fmt.Println("insert iax_response_times success ,id=", id)
			}

			if err != nil {
				fmt.Println("insert ", key.isTest, " table  error ", err)
				continue
			}
			fmt.Println("insert ", key.isTest, " table success ,id=", id)

		}

		delkey(value)
		fmt.Println(count, percent99, percent85, percent55, index99, index85, index55)
	}
}

func delkey(value []bidtimes) {
	for _, key := range value {
		delcache(key.sourceKey)
	}
}

func caclGroup(value []bidtimes) (int, []int) {
	count := 0
	var useTimes = make([]int, 500)
	for _, timesvalue := range value {
		count = count + timesvalue.count
		index := timesvalue.useTimes - 1
		if index > 500 {
			index = 500
		}
		useTimes[index] = timesvalue.count
	}
	return count, useTimes
}

func listAllKey(keyPre string) ([]string, error) {
	redisConn := dao.RedisClient.Get()
	defer redisConn.Close()

	ret, err := redisConn.Do("keys", keyPre)
	if err != nil {
		return nil, err
	}
	if arr, ok := ret.([]interface{}); ok {
		result := make([]string, len(arr))
		for index, value := range arr {
			if strval, valOk := value.([]byte); valOk {
				result[index] = string(strval)
			}
		}
		return result, nil
	}
	return nil, errors.New("Empty result")
}

func getValue(key string) int {
	redisConn := dao.RedisClient.Get()
	defer redisConn.Close()
	ret, err := redisConn.Do("get", key)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if value, ok := ret.([]byte); ok {
		ret, err := strconv.Atoi(string(value))
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return ret
	}
	fmt.Println(" error. . . ", key)
	return 0
}

func delcache(key string) {
	redisConn := dao.RedisClient.Get()
	defer redisConn.Close()
	_, err := redisConn.Do("del", key)
	if err != nil {
		fmt.Println("del ", key, " error,", err)
	}
}

func filterKeys(keys []string) (vaildKey []bidtimes) {

	var result []bidtimes
	for _, value := range keys {

		fields := strings.Split(value, "[iax]")
		theTime, err := time.ParseInLocation("2006-01-02-15-04-05", fields[2]+"-00", time.UTC)
		if err != nil {
			fmt.Println("time parse error", fields[2], err)
			continue
		}

		stopTime := theTime.Add(time.Minute * 5)

		// if stoptime < now  ok  else skip .
		if stopTime.After(time.Now().UTC()) {
			// thetime > stoptime
			// logger.Warn("skip stoptime>now,stop=%s(%v) %v %v", stopTime.Format("2006-01-02 15:04:05"), stopTime.Unix(), theTime, time.Now())
			continue
		}

		times := getValue(value)
		if times < 1 {
			fmt.Println(" < 0 ? ", value, times)
			continue
		}
		dspID, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println(fields[1], dspID)
			continue
		}
		useTimes, err := strconv.Atoi(fields[3])
		if err != nil {
			fmt.Println(fields[3], useTimes)
			continue
		}
		isTest := strings.HasSuffix(fields[0], "-TEST")
		bidtime := bidtimes{groupKey{dspID, isTest, theTime}, value, useTimes, times}
		result = append(result, bidtime)
	}
	return result
}

type groupKey struct {
	dspID   int
	isTest  bool
	theTime time.Time
}

type bidtimes struct {
	groupKey
	sourceKey string
	useTimes  int
	count     int
}

func (times *bidtimes) toString() string {
	var buffer bytes.Buffer
	buffer.WriteString("bidtimes{sourceKey=")
	buffer.WriteString(times.sourceKey)
	buffer.WriteString(",dspID=")
	buffer.WriteString(strconv.Itoa(times.groupKey.dspID))
	buffer.WriteString(",theTime=")
	buffer.WriteString(times.groupKey.theTime.Format("2006-01-02 15:04:05"))
	buffer.WriteString(",useTime=")
	buffer.WriteString(strconv.Itoa(times.useTimes))
	buffer.WriteString(",count=")
	buffer.WriteString(strconv.Itoa(times.count))
	buffer.WriteString("}")
	return buffer.String()
}
