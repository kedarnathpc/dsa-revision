func missingNumber(nums []int) int {
    xorr, i := 0,0

    for _, it := range nums {
        xorr ^= (i ^ it)
        i++
    }

    return xorr ^ i
}
