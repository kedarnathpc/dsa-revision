func findMaxConsecutiveOnes(nums []int) int {
    count, maxi := 0, 0

    for i := 0; i < len(nums); i++ {
        count += nums[i]
        if nums[i] == 0 {
            count = 0
        }
        maxi = int(math.Max(float64(maxi), float64(count)))
    }

    return maxi
}
