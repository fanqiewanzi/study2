package calculate

import "fmt"

func compare(str1, str2 string) int {
	//第一个比第二个大就返回1，小就返回-1，同级返回0
	switch str1 {
	case "+", "-":
		if str2 == "(" || str2 == "#" {
			return 1
		} else {
			return -1
		}
	case "*", "/":
		if str2 == "*" || str2 == "/" {
			return -1
		} else {
			return 1
		}
	case "(":
		return 1
	case ")":
		if str2 == "(" {
			return 0
		} else {
			return -1
		}
	case "#":
		if str2 == "#" {
			return 0
		} else {
			return -1
		}
	default:
	}
	return 1
}

func Calculate(str string) {
	//定义两个栈,初始化为10的长度
	dNum := 0
	oNum := 1
	digitalStack := make([]float64, 10)
	oprateStack := make([]string, 10)
	//将第一个运算符设置为最低级的运算符#
	oprateStack[0] = "#"
	result := 0.0

	for i := 0; i < len(str); {
		if str[i] >= 48 && str[i] <= 57 {
			fmt.Println(str[i])
			digitalStack[dNum] = float64(str[i] - 48.0)
			dNum++
		} else {
			op := oprateStack[oNum]
			k := compare(string(str[i]), op)
			if k == 1 {
				oNum++
				oprateStack[oNum] = string(str[i])
			} else if k == -1 {
				dNum--
				first := digitalStack[dNum]
				dNum--
				second := digitalStack[dNum]
				op := oprateStack[oNum]
				oNum--
				switch op {
				case "+":
					result = second + first
				case "-":
					result = second - first
				case "*":
					result = second * first
				case "/":
					result = second / first
				}
				dNum++
				digitalStack[dNum] = result
			} else //同级出栈扫描下一个
			{
				oNum--
				i++
			}
		}
	}
	dNum--
	fmt.Println(digitalStack[dNum])
}
