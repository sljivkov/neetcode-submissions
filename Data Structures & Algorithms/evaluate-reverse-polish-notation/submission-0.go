func evalRPN(tokens []string) int {

	stack := make([]int, 0)

	for _, t := range tokens {
		if num, err:= strconv.Atoi(t); err == nil {
			stack = append(stack, num)

		} else {
			num1 := pop(&stack)
			num2 := pop(&stack)

			switch t {
			case "+":
				stack = append(stack, num1 + num2)
			case "-":
				stack = append(stack, num2 - num1)
			case "*":
				stack = append(stack, num1 * num2)
			default:
				stack = append(stack, num2 / num1)
			}
		}
	}

	return stack[len(stack)-1]
}

func pop(stack *[]int) int {
	s := *stack
	num := s[len(s) - 1]
	*stack = s[:len(s) - 1]
	return num
}