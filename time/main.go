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
	fmt.Println(zodiac[(t.Year()+8)%12])
	fmt.Println(tiangan[(t.Year()+8)%12])
	fmt.Println(dizhi[(t.Year()+6)%10])
}