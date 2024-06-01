package mtime

import "time"

// Today 今天
func Today() (startTime int64, endTime int64) {
	nowTime := time.Now()
	startTime = time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	endTime = nowTime.UnixMilli()

	return
}

// Yestoday 昨天时间
func Yestoday() (startTime int64, endTime int64) {
	start := time.Now().Add(-24 * time.Hour)

	startTime = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	endTime = time.Date(start.Year(), start.Month(), start.Day(), 23, 59, 59, 999, time.Local).UnixMilli()

	return
}

// LastNDays 最近N天
func LastNDays(num time.Duration) (startTime int64, endTime int64) {
	now := time.Now()
	start := now.Add(-24 * time.Hour * (num - 1))

	startTime = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	endTime = now.UnixMilli()

	return
}

// Last3Days 最近三天
func Last3Days() (startTime int64, endTime int64) {
	now := time.Now()
	start := now.Add(-24 * time.Hour * 2)

	startTime = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	endTime = now.UnixMilli()

	return
}

// Last7Days 最近7天
func Last7Days() (startTime int64, endTime int64) {
	now := time.Now()
	start := now.Add(-24 * time.Hour * 6)

	startTime = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	endTime = now.UnixMilli()

	return
}

// LastMonth 最近一个月
func LastMonth() (startTime int64, endTime int64) {
	nowTime := time.Now()
	startTime = time.Date(nowTime.Year(), nowTime.Month()-1, nowTime.Day(), 0, 0, 0, 0, time.Local).UnixMilli()

	endTime = nowTime.UnixMilli()

	return
}

// Last3Months 最近三个月
func Last3Months() (startTime int64, endTime int64) {
	nowTime := time.Now()
	startTime = time.Date(nowTime.Year(), nowTime.Month()-3, nowTime.Day(), 0, 0, 0, 0, time.Local).UnixMilli()

	endTime = nowTime.UnixMilli()

	return
}

// Month 本月
func Month() (startTime int64, endTime int64) {
	nowTime := time.Now()
	startTime = time.Date(nowTime.Year(), nowTime.Month(), 0, 0, 0, 0, 0, time.Local).UnixMilli()

	endTime = nowTime.UnixMilli()

	return
}

// Year 本年
func Year() (startTime int64, endTime int64) {
	nowTime := time.Now()
	startTime = time.Date(nowTime.Year(), 0, 0, 0, 0, 0, 0, time.Local).UnixMilli()

	endTime = nowTime.UnixMilli()

	return
}
