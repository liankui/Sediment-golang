package main

import (
	"fmt"
	"strings"
	"time"
)

/*
time包有几种格式：
1.Time
2.string
3.Duration 时间间隔 nanosecond 纳秒，格式为24h0m0s
*/

// 使用一个结构封装time.Time可以转化带时区的时间为不带时区的时间
// 例如：2020-09-30T16:47:06+08:00 -> 2020-09-30 16:47:06
type JsonTime struct {
	time.Time
}
// 使用JsonTime时，如果值为null，json对应的值为0

func main() {
	// time.Parse 将string->Time，并返回Time
	startTimeString := "1970-01-01 00:00:00"
	startTime, _ := time.Parse("2006-01-02 15:04:05", startTimeString)
	fmt.Println("startTime:", startTime)	// startTime: 1970-01-01 00:00:00 +0000 UTC

	endTimeString := "2100-06-05 00:00:00"
	endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeString)
	fmt.Println("endTime:", endTime)	// endTime: 2100-06-05 00:00:00 +0000 UTC

	// After用于比较2个Time格式的时间点
	if endTime.After(startTime) {
		fmt.Println("endTime is after startTime")
	}

	// time.ParseInLocation 将string转化为当地的Time,并返回Time
	time1String := "2020-06-09 10:51:06"
	Time1, _ := time.ParseInLocation("2006-01-02 15:04:05", time1String, time.Local)
	fmt.Println("Time1:", Time1)	// Time1: 2020-06-09 10:51:06 +0800 CST

	// 将 2020年06月09日 转化为正常时间格式
	s1 := "2020年06月09日"
	str1 := s1[:4]+"-"+s1[7:9]+"-"+s1[12:14] // 一个中文字符对应3个字节
	Time12, _ := time.ParseInLocation("2006-01-02", str1, time.Local) // Time12: 2020-06-09 00:00:00 +0800 CST
	format := Time12.Format("2006-01-02")
	fmt.Println("Time12:", format)	// Time12: 2020-06-09
	// 方法二：使用replacer
	replacer := strings.NewReplacer("年", "-", "月", "-", "日", "")
	fixed := replacer.Replace(s1)
	fmt.Println("fixed:", fixed) // fixed: 2020-06-09

	// Time1.AddDate 增加天数
	Time2 := Time1.AddDate(0, 0, 1)
	fmt.Println("Time2:", Time2)	// Time2: 2020-06-10 10:51:06 +0800 CST

	// Sub 前面的时间 减 后面的时间
	DurationSub := Time2.Sub(Time1)
	fmt.Println("DurationSub:", DurationSub) // 24h0m0s
	fmt.Println("DurationSub.String:", DurationSub.String()) // 24h0m0s

	// Add, Time 加 时间间隔，返回一个Time
	Time3 := Time1.Add(DurationSub)
	fmt.Println("Time3:", Time3)	// Time3: 2020-06-10 10:51:06 +0800 CST 和 Time2时间相同

	// 将时间调至3day前。 Time最大时间尺度是Hour
	TimeSub3Day := Time1.Add(-time.Hour * 24 * 3)
	fmt.Println("TimeSub3Day:", TimeSub3Day)	// TimeSub3Day: 2020-06-06 10:51:06 +0800 CST

	// 将时间调至1年前
	TimeSub1Year := Time1.AddDate(-1, 0, 0)
	fmt.Println("TimeSub1Year:", TimeSub1Year)	// TimeSub1Year: 2019-06-09 10:51:06 +0800 CST

}
