func findComplement(num int) int {
    bitLength, n := 0, num

    for num > 0 {
        bitLength++
        num >>= 1
    }

    mask := (1 << bitLength) - 1

    return n ^ mask
}
