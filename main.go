package main

import (
	"fmt"
	"strconv"
)

var ans []string

func calcAllStrings(s string, nums int) {
	if nums == -1 {
		ans = append(ans, s)
		return
	}
	calcAllStrings(s+strconv.Itoa(nums), nums-1)

	if nums != 9 {
		calcAllStrings(s+"+"+strconv.Itoa(nums), nums-1)
	}
	calcAllStrings(s+"-"+strconv.Itoa(nums), nums-1)
}

func calcResult(s string) bool {
	if len(s) == 10 {
		return false
	}
	sum := 0
	tmp := 0
	var flag byte = '+'

	// у нас нет особых символов UTF8, поэтому можно итерировать по
	for i := 0; i < len(s); i++ {
		if s[i] != 43 && s[i] != 45 {
			tmp = tmp*10 + int(s[i]-48)
		} else {
			if flag == 43 {
				flag = s[i]
				sum += tmp
				tmp = 0
			} else {
				flag = s[i]
				sum -= tmp
				tmp = 0
			}
		}
	}
	if flag == 43 {
		sum += tmp
	} else {
		sum -= tmp
	}
	if sum == 200 {
		return true
	}
	return false

}

func main() {
	ans = make([]string, 0)
	calcAllStrings("", 9)
	for _, v := range ans {
		if calcResult(v) {
			fmt.Println(v)
		}
	}

}
