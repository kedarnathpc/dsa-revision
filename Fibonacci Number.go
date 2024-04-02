func fabonacciNum (n int) int {
    if n <= 1 {
        return n
    }
    return fabonacciNum(n - 1) + fabonacciNum(n - 2)
}

func fib(n int) int {
    return fabonacciNum(n)
}
