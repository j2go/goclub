package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Story struct {
	Id               int       `orm:"column(id);auto"`
	CreateTime       time.Time `orm:"column(create_time);type(datetime);null"`
	Creator          string    `orm:"column(creator);size(255);null"`
	LastModifiedTime time.Time `orm:"column(last_modified_time);type(datetime);null"`
	Content          string    `orm:"column(content);size(2000)"`
	IsDeleted        uint64    `orm:"column(is_deleted);size(1)"`
	LastReplyTime    time.Time `orm:"column(last_reply_time);type(date);null"`
	LastReplyUserId  int64     `orm:"column(last_reply_user_id);null"`
	LikeTimes        int       `orm:"column(like_times)"`
	PageViews        int       `orm:"column(page_views)"`
	ReplyTimes       int       `orm:"column(reply_times)"`
	Title            string    `orm:"column(title);size(100)"`
	TabId            int       `orm:"column(tab_id)"`
}

func (t *Story) TableName() string {
	return "story"
}

func init() {
	orm.RegisterModel(new(Story))
}

// AddStory insert a new Story into database and returns
// last inserted Id on success.
func AddStory(m *Story) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStoryById retrieves Story by Id. Returns error if
// Id doesn't exist
func GetStoryById(id int) (v *Story, err error) {
	o := orm.NewOrm()
	v = &Story{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllStory retrieves all Story matches certain condition. Returns empty list if
// no records exist
func GetAllStory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Story))
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

	var l []Story
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

// UpdateStory updates Story by Id and returns error if
// the record to be updated doesn't exist
func UpdateStoryById(m *Story) (err error) {
	o := orm.NewOrm()
	v := Story{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStory deletes Story by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStory(id int) (err error) {
	o := orm.NewOrm()
	v := Story{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Story{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
