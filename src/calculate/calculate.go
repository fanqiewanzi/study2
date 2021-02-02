package calculate

import "fmt"

func compare(str1, str2 uint8) int {
	//第一个比第二个大就返回1，小就返回-1，同级返回0
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
	default:
	}
	return 1
}

func Calculate(str string) {
	//定义两个栈,初始化为10的长度,一个栈用来存运算对象，一个栈用来存运算符
	//numStack是uint8类型,opStack是int32
	stack1 := make([]interface{}, 10)
	stack2 := make([]interface{}, 10)
	numStack := NewSeqStack(-1, stack1)
	opStack := NewSeqStack(-1, stack2)
	//将第一个运算符设置为最低级的运算符#
	opStack.Push('#')
	//用来存放每次的运算结果的result
	result := 0.0
	//进入循环扫描字符串
	for i := 0; i < len(str); {
		//如果是数字则进数字栈
		if str[i] >= '0' && str[i] <= '9' {
			numStack.Push(str[i])
			i++
		} else {
			//进入运算符判断环节，将栈顶运算符取出与当前运算符进行比较
			op := opStack.GetTop()
			k := compare(str[i], op.(uint8))
			//如果当前运算符比栈顶运算符大则将当前运算符入栈
			if k == 1 {
				opStack.Push(str[i])
			} else if k == -1 {
				//比当前运算符小则进行运算
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
				numStack.Push(result)
			} else //同级出栈扫描下一个
			{
				opStack.Pop()
				i++
			}
		}
	}
	fmt.Println(numStack.GetTop())
}
