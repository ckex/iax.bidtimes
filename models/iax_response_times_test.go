package models

import (
	"fmt"
	"testing"
	"time"

	_ "iax.bidtimes/dao"
	"iax.bidtimes/models"
)

func Test_Insert(t *testing.T) {
	// loc, _ := time.LoadLocation("CTS")
	// fmt.Printf("time local=%v\n", loc)
	iaxResponseTimes := models.IaxResponseTimes{
		USERID:           int64(1),
		TIME:             time.Now(),
		FIFTYTH:          105,
		EIGHTYFIVETH:     185,
		NINETYNINETH:     199,
		DATETIMECREATE:   time.Now(),
		DATETIMEMODIFIED: time.Now(),
	}
	id, err := models.AddIaxResponseTimes(&iaxResponseTimes)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("insert success %v\n", id)
}

// func Test_Get(t *testing.T) {
// 	fmt.Println("-----------test.")
// 	times , err := GetIaxResponseTimesById(int64(1))
// 	fmt.Println("error==> ",err)
// 	fmt.Println("times==> ",times)
// }
