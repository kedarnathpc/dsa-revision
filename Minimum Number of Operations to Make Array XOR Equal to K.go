func minOperations(nums []int, k int) int {
    xorr := 0

    for _, num := range nums {
        xorr ^= num
    }

    if xorr == k {
        return 0
    }

    cnt := 0
    for k != 0 || xorr != 0 {
        if k&1 != xorr&1 {
            cnt++
        }
        k = k >> 1
        xorr = xorr >> 1
    }

    return cnt
}
