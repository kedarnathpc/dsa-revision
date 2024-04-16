func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func asteroidCollision(nums []int) []int {
    stack := make([]int, 0)

    for _, a := range nums {
        for len(stack) > 0 && a < 0 && stack[len(stack)-1] > 0 {
            if stack[len(stack)-1] > -a {
                a = 0
                break
            } else if stack[len(stack)-1] < -a {
                stack = stack[:len(stack)-1]
            } else {
                stack = stack[:len(stack)-1]
                a = 0
                break
            }
        }

        if a != 0 {
            stack = append(stack, a)
        }
    }

    return stack
}
