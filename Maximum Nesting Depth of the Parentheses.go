func maxDepth(s string) int {
    count, maxi := 0,0

    for _, i := range s {
        if i == '(' {
            count++
        } else if i == ')' {
            count--
        }
        maxi = max(maxi, count)
    }

    return maxi
}
