package mix

import "fmt"

func Count(str string) {
	//定义四个变量分别存储数字个数，英文个数，空格个数和其他字符个数初始化为0
	digitNum, englishNum, spaceNum, otherNum := 0, 0, 0, 0
	//对字符串中的字符进行判断
	for _, elem := range str {
		switch {
		case elem >= '0' && elem <= '9':
			digitNum++
		case elem >= 'a' && elem <= 'z' || elem >= 'A' && elem <= 'Z':
			englishNum++
		case elem == ' ':
			spaceNum++
		default:
			otherNum++
		}
	}
	fmt.Printf("字符串中%s包含：\n数字:%d\t字母:%d\t空格:%d\t其他:%d\n", str, digitNum, englishNum, spaceNum, otherNum)
}

func Fibonacci() {
	//定义一个长50的数组并初始化前两个变量为0和1
	fibo := [50]int{0, 1}
	//打印输出斐波拉契数组
	fmt.Printf("斐波拉契数组为:%d %d", fibo[0], fibo[1])
	for i := 2; i < len(fibo); i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		fmt.Printf(" %d", fibo[i])
	}
}
