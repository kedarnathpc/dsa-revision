func maxFrequency(nums []int, k int) int {
    left, right, total, n, maxi := 0, 0, 0, len(nums), 0

    sort.Ints(nums)

    for right < n {
        total += nums[right]
        for nums[right] * (right - left + 1) > total + k {
            total -= nums[left]
            left++
        }
        maxi = int(math.Max(float64(maxi), float64(right - left + 1)))
        right++
    }

    return maxi
}

