package main

import (
	"fmt"
	"math"
	"sort"
)
func main() {
	var a = [10]int{}
	for i, _ := range a {
		a[i] = 1
	}
	s, _ := sol(a, 9)
	fmt.Println(s)
	for key, val := range s {
		switch val {
		case 1:
			fmt.Printf("A ")
		case 2:
			fmt.Printf("B ")
		case 3:
			fmt.Printf("C ")
		case 4:
			fmt.Printf("D ")
		default:
			fmt.Printf("%d+%d ", key, val)
		}
	}
}
// 递归函数
func sol(a [10]int, i int, ) ([10]int, int) {
	if i < 0 {
		return a, i + 1
	} else {
		for a[i] < 5 {
			a, i = sol(a, i-1)
			if IsRight(a) != 10 {
				a[i]++
			} else {
				return a, i + 1
			}
		}
		if IsRight(a) == 10 {
			return a, i + 1
		} else {
			a[i] = 1
			return a, i + 1
		}
	}
}
//判断答案，返回答对题目总数
func IsRight(a [10]int) int {
	pass := 1
	count := 0
	//question2
	var t int
	switch a[1] {
	case 1:
		t = 3
	case 2:
		t = 4
	case 3:
		t = 1
	case 4:
		t = 2
	}
	if t == a[4] {
		pass++
	} else {
		return 1
	}
	//question 3
	switch a[2] {
	case 1:
		if (a[3] != a[2] && a[2] != a[5] && a[1] != a[2]) {
			pass++
			count++
		} else {
			return 2
		}
	case 2:
		if (a[1] != a[5] && a[2] != a[5] && a[5] != a[3]) {
			pass++
			count++
		} else {
			return 2
		}
	case 3:
		if (a[1] != a[2] && a[1] != a[3] && a[1] != a[5]) {
			pass++
			count++
		} else {
			return 2
		}
	case 4:
		if (a[3] != a[5] && a[1] != a[3] && a[2] != a[3]) {
			pass++
			count++
		} else {
			return 2
		}
	default:
		return 2
	}
	if count != 1 {
		return 2
	}

	//question4
	count = 0
	switch a[3] {
	case 1:
		if a[0] == a[4] {
			pass++
			count++
		} else {
			return pass
		}
	case 2:
		if a[1] == a[6] {
			pass++
			count++
		} else {
			return pass
		}
	case 3:
		if a[0] == a[8] {
			pass++
			count++
		} else {
			return pass
		}
	case 4:
		if a[5] == a[9] {
			pass++
			count++
		} else {
			return pass
		}
	default:
		return pass
	}
	if count != 1 {
		return pass
	}

	//question5
	count = 0
	if a[7] == 1 && a[4] == 1 {
		pass++
		count++
	} else if a[3] == 2 && a[4] == 2 {
		pass++
		count++
	} else if a[8] == 3 && a[4] == 3 {
		pass++
		count++
	} else if a[6] == 4 && a[4] == 4 {
		pass++
		count++
	}
	if count != 1 {
		return pass
	}

	//question 6
	count = 0
	switch a[5] {
	case 1:
		if a[1] == a[3] && a[3] == a[7] {
			pass++
			count++
		}
	case 2:
		if a[0] == a[7] && a[5] == a[7] {
			pass++
			count++
		}
	case 3:
		if a[2] == a[7] && a[7] == a[9] {
			pass++
			count++
		}
	case 4:
		if a[4] == a[7] && a[7] == a[8] {
			pass++
			count++
		}
		if count != 1 {
			return pass
		}

	}
	//question 7
	s := least(a)
	if s == a[6] {
		pass++
	} else {
		return pass
	}
	//question 8
	if notnear(a[0], a[1]) || notnear(a[0], a[4]) || notnear(a[0], a[6]) || notnear(a[0], a[9]) {
		pass++
	} else {
		return pass
	}

	//question 9
	if a[0] == a[5] {
		if a[1] != a[4] || a[5] != a[4] || a[8] != a[4] || a[9] != a[4] {
			pass++
		}
	} else if a[1] == a[4] || a[5] == a[4] || a[8] == a[4] || a[9] == a[4] {
		pass ++
	} else {
		return pass
	}
	//question 10
	if a[9] == countnum(a) {
		pass ++
	}
	return pass

}
//最少的答案数
func least(a [10]int) int {
	var s1, s2, s3, s4 = 0, 0, 0, 0
	for _, v := range a {
		switch v {
		case 1:
			s1++
		case 2:
			s2++
		case 3:
			s3++
		case 4:
			s4++
		}
	}
	arry := []int{s1, s2, s3, s4}
	sort.Ints(arry)
	switch arry[0] {
	case s1:
		return 1
	case s2:
		return 2
	case s3:
		return 3
	case s4:
		return 4
	default:
		return 0
	}
}

func notnear(a int, b int) bool {
	if math.Abs(float64(a)-float64(b)) > 1 {
		return true
	} else {
		return false
	}
}
//统计各个答案总数
func countnum(a [10]int) int {
	var s1, s2, s3, s4 int = 0, 0, 0, 0
	for _, v := range a {
		switch v {
		case 1:
			s1++
		case 2:
			s2++
		case 3:
			s3++
		case 4:
			s4++
		}
	}
	arry := []int{s1, s2, s3, s4}
	sort.Ints(arry)
	switch arry[3] - arry[0] {
	case 1:
		return 4
	case 2:
		return 2
	case 3:
		return 1
	case 4:
		return 3
	default:
		return 0
	}
}
