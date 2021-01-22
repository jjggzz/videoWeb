package repository

import (
	"time"
)

type CustomerExample struct {
	criteria []struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}
}

func (ex *CustomerExample) Clear() *CustomerExample {
	ex.criteria = []struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{}
	return ex
}

func (ex *CustomerExample) AndIdIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndIdEqualTo(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdNotEqualTo(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThan(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdGreaterThanOrEqualTo(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndIdLessThan(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdLessThanOrEqualTo(param int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdIn(param []int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdNotIn(param []int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndIdBetween(param1 int64, param2 int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndIdNotBetween(param1 int64, param2 int64) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and id not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndAccessKeyLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndAccessKeyBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndAccessKeyLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndAccessKeyNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and access_key not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndCreateTimeLessThan(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeIn(param []time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotIn(param []time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndCreateTimeBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndCreateTimeNotBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and create_time not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThan(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeGreaterThanOrEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndUpdateTimeLessThan(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeLessThanOrEqualTo(param time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeIn(param []time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotIn(param []time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUpdateTimeNotBetween(param1 time.Time, param2 time.Time) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and update_time not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThan(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndDeleteStatusLessThan(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusIn(param []int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotIn(param []int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusBetween(param1 int32, param2 int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndDeleteStatusNotBetween(param1 int32, param2 int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and delete_status not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndPhoneLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndPhoneBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPhoneLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndPhoneNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and phone not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndUsernameLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndUsernameBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndUsernameLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndUsernameNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and username not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndPasswordLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndPasswordBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndPasswordLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndPasswordNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and password not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndEmailLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndEmailBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndEmailLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndEmailNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and email not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameNotEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameGreaterThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndNicknameLessThan(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameLessThanOrEqualTo(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameNotIn(param []string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndNicknameBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameNotBetween(param1 string, param2 string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndNicknameLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname like ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndNicknameNotLike(param string) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and nickname not like ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusIsNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status is null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusIsNotNull() *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status is not null", noValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status = ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusNotEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status <> ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThan(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status > ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusGreaterThanOrEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status >= ?", param1: param})
	return ex
}
func (ex *CustomerExample) AndStatusLessThan(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status < ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusLessThanOrEqualTo(param int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status <= ?", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusIn(param []int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusNotIn(param []int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status not in (?)", param1: param})
	return ex
}

func (ex *CustomerExample) AndStatusBetween(param1 int32, param2 int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}

func (ex *CustomerExample) AndStatusNotBetween(param1 int32, param2 int32) *CustomerExample {
	ex.criteria = append(ex.criteria, struct {
		fragment     string
		param1       interface{}
		param2       interface{}
		noValue      bool
		betweenValue bool
	}{fragment: "and status not between ? and ?", param1: param1, param2: param2, betweenValue: true})
	return ex
}
