package mix

import "fmt"

func Count(str string) {
	digitNum, englishNum, spaceNum, otherNum := 0, 0, 0, 0
	for _, elem := range str {
		switch {
		case elem >= 48 && elem <= 57:
			digitNum++
		case elem >= 65 && elem <= 90 || elem >= 97 && elem <= 122:
			englishNum++
		case elem == 32:
			spaceNum++
		default:
			otherNum++
		}

	}
	fmt.Println("数字:%d\t字母:%d\t空格:%d\t其他:%d\t", digitNum, englishNum, spaceNum, otherNum)
}

func Fibonacci() {
	var fibo [50]int
	fibo[0] = 0
	fibo[1] = 1
	fmt.Print(fibo[0])
	fmt.Print(" ")
	fmt.Print(fibo[1])
	for i := 2; i < len(fibo); i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		fmt.Print(" ")
		fmt.Print(fibo[i])

	}
}
