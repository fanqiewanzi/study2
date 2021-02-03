package calculate

import "fmt"

func compare(str1, str2 int32) int {
	//扫描的比栈顶的运算符大就返回1，小就返回-1，同级返回0
	switch str1 {
	case '+', '-':
		if str2 == '(' || str2 == '#' {
			return 1
		} else {
			return -1
		}
	case '*', '/':
		if str2 == '*' || str2 == '/' {
			return -1
		} else {
			return 1
		}
	case '(':
		return 1
	case ')':
		if str2 == '(' {
			return 0
		} else {
			return -1
		}
	case '#':
		if str2 == '#' {
			return 0
		} else {
			return -1
		}
	}
	return 1
}

//
func Calculate(str string) {
	//定义两个栈,初始化为10的长度,一个栈用来存运算对象，一个栈用来存运算符，第一个栈是float64类型，第二栈是int32类型
	numStack := NewSeqStack()
	opStack := NewSeqStack()
	//将第一个运算符设置为最低级的运算符#
	opStack.Push(int32('#'))
	//用来存放每次的运算结果的result
	var result float64 = 0
	//进入循环扫描字符串
	//因为用str[i]的方式读取的值的类型会不同，所以每次读取中我都将str[i]的值进行转换来同一类型方便后续计算和比较
	//此循环最后一个扫描到的就是我们放进去的最低级的运算符#这时候表达式就运算完毕了
	str1 := str + "#"
	for i := 0; i < len(str1); {
		//如果是数字则进数字栈
		if str1[i] >= '0' && str1[i] <= '9' {
			numStack.Push(float64(str1[i]) - '0')
			i++
		} else {
			//进入运算符判断环节，将栈顶运算符取出与当前运算符进行比较
			op := opStack.GetTop()
			k := compare(int32(str1[i]), op.(int32))
			//如果当前运算符比栈顶运算符大则将当前运算符入栈
			if k == 1 {
				opStack.Push(int32(str1[i]))
				i++
			} else if k == -1 {
				//如果当前运算符比栈顶运算符小则进行运算
				first := numStack.Pop().(float64)
				second := numStack.Pop().(float64)
				op := opStack.Pop()
				switch op {
				case '+':
					result = second + first
				case '-':
					result = second - first
				case '*':
					result = second * first
				case '/':
					result = second / first
				}
				//将计算结果放进栈中，这里在下个循环中会将刚才扫描到的运算符再扫描一遍进行判断
				numStack.Push(result)
			} else //同级出栈扫描下一个，因为同级的在规则中只有()和#和#是同级的是不需要运算的，两个运算符比较为同级说明中间的运算符都已运算完毕所以出栈扫描下一个
			{
				opStack.Pop()
				i++
			}
		}
	}
	//输出栈顶的值就是表达式的结果
	fmt.Printf("\n表达式%s结果为%f\n", str, numStack.GetTop())
}
