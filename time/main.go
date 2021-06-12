package main

import (
	"fmt"
	"time"
)

func main() {
	zodiac := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	tiangan := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	dizhi := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	t := time.Now()
	fmt.Printf("%+v\n", t)
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	fmt.Println(zodiac[(t.Year()+8)%12])
	fmt.Println(tiangan[(t.Year()+8)%12])
	fmt.Println(dizhi[(t.Year()+6)%10])

	fmt.Println(time.Unix(-100000, 0).Format("20060102030405"))

	ts, err := time.Parse(time.RFC3339, "")
	fmt.Println(ts, err)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	ts, err = time.ParseInLocation("1/02/2006, 3:04:05 PM", "6/10/2021, 8:26:03 PM", loc)
	fmt.Println(ts, err)
}
