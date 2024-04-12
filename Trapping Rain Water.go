func trap(height []int) int {
    n := len(height)
    left, right := 0, n-1
    maxLeft, maxRight, res := height[left], height[right], 0

    for left < right {
        if height[left] < height[right]{
            left++
            maxLeft = max(maxLeft, height[left])
            res += maxLeft - height[left]
        } else {
            right--
            maxRight = max(maxRight, height[right])
            res += maxRight - height[right]
        }
    }

    return res
}
