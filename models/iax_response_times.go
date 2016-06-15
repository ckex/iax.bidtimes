package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type IaxResponseTimes struct {
	Id               int64     `orm:"column(ID);auto"`
	USERID           int64     `orm:"column(USER_ID);"`
	TIME             time.Time `orm:"column(TIME);type(datetime)"`
	FIFTYTH          int       `orm:"column(FIFTY_TH)"`
	EIGHTYFIVETH     int       `orm:"column(EIGHTY_FIVE_TH)"`
	NINETYNINETH     int       `orm:"column(NINETY_NINE_TH)"`
	DATETIMEMODIFIED time.Time `orm:"column(DATETIME_MODIFIED);type(datetime)"`
	DATETIMECREATE   time.Time `orm:"column(DATETIME_CREATE);type(datetime)"`
}

func (t *IaxResponseTimes) TableName() string {
	return "iax_response_times"
}

func init() {
	fmt.Println("init times. ")
	orm.RegisterModel(new(IaxResponseTimes))
	fmt.Println("init times finished. ")
}

// AddIaxResponseTimes insert a new IaxResponseTimes into database and returns
// last inserted Id on success.
func AddIaxResponseTimes(m *IaxResponseTimes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIaxResponseTimesById retrieves IaxResponseTimes by Id. Returns error if
// Id doesn't exist
func GetIaxResponseTimesById(id int64) (v *IaxResponseTimes, err error) {
	o := orm.NewOrm()
	v = &IaxResponseTimes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllIaxResponseTimes retrieves all IaxResponseTimes matches certain condition. Returns empty list if
// no records exist
func GetAllIaxResponseTimes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(IaxResponseTimes))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []IaxResponseTimes
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateIaxResponseTimes updates IaxResponseTimes by Id and returns error if
// the record to be updated doesn't exist
func UpdateIaxResponseTimesById(m *IaxResponseTimes) (err error) {
	o := orm.NewOrm()
	v := IaxResponseTimes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIaxResponseTimes deletes IaxResponseTimes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIaxResponseTimes(id int64) (err error) {
	o := orm.NewOrm()
	v := IaxResponseTimes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&IaxResponseTimes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
