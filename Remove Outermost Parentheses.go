func removeOuterParentheses(s string) string {
    stack := make([]rune, 0)
    res := ""

    for _, c := range s {
        if c == '(' {
            if len(stack) != 0 {
                res += string(c)
            }
            stack = append(stack, c)
        } else {
            if len(stack) == 1 {
                stack = stack[:len(stack)-1]
            } else {
                res += string(c)
                stack = stack[:len(stack)-1]
            }
        }
    }

    return res
}
