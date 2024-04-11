func removeKdigits(num string, k int) string {
    stack := make([]rune, 0)

    for _, c := range num {
        for len(stack) > 0 && k > 0 && stack[len(stack)-1] > c {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, c)
    }

    for k > 0 && len(stack) > 0 {
        stack = stack[:len(stack)-1]
        k--
    }

    result := string(stack)

    for len(result) > 1 && result[0] == '0' {
        result = result[1:]
    }

    if len(result) > 0 {
        return result
    }

    return "0"
}
