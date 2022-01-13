/**
 * Created by angelina-zf on 17/2/27.
 */
package leotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Year
// 获取本年 整型
func Year() int {
	return time.Now().Year()
}

// Month
// 获取本月
func Month() time.Month {
	now := time.Now()
	_, month, _ := now.Date()
	return month
}

// Day
// 获取本日
func Day() int {
	return time.Now().Day()
}

// Day
// 获取本日
func Minute() int {
	return time.Now().Minute()
}

// Week
// 获取本日在这周是周几
func Week() time.Weekday {
	now := time.Now()
	return now.Weekday()
}

// TimeToUnix
// 转换为10位的unix时间戳 精确到秒
func TimeToUnixS(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

// UnixSToTime
// 10位的unix时间戳转换为time.Time
func UnixSToTime(st string) time.Time {
	stInt, _ := strconv.Atoi(st)
	return time.Unix(int64(stInt), 0)
}

func CustomLocation(t time.Time, dis time.Duration) time.Time {
	return TimeLocation(t, "Custom", dis)
}

/**
dis 与 UTC的时差 （+-）; 东8区设置为 8
*/
func TimeLocation(t time.Time, name string, dis time.Duration) time.Time {
	return t.In(time.FixedZone(name, int((dis * time.Hour).Seconds())))
}

// 默认的时间格式
const (
	ForMate_yyyymmddhhmmss = "YYYY-MM-DD HH:mm:ss"
	ForMate_yyyymmdd       = "YYYY-MM-DD"
)

// DateFormat 格式化时间
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateFormat(t time.Time, format string) string {
	//t.In(time.FixedZone())
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}
