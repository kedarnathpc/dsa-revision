func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    
    first, second, third, i := 0, 1, 1, 3

    for i <= n {
        sum := first + second + third
        first = second
        second = third
        third = sum
        i++
    }

    return third
}
