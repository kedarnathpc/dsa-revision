func makeGood(s string) string {
    var stack []rune
    res := ""

    for _, i := range s {
        if len(stack) > 0 && (stack[len(stack)-1] == i+32 || stack[len(stack)-1]+32 == i) {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, i)
        }
    }

    for _, i := range stack {
        res += string(i)
    }

    return res
}
