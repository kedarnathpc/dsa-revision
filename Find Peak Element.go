func findPeakElement(nums []int) int {
    left, right, n := 1, len(nums)-2, len(nums)

    if n == 1 {
        return 0
    }
    if nums[0] > nums[1] {
        return 0
    } 
    if nums[n-1] > nums[n-2] {
        return n-1
    }

    for left <= right {
        mid := left + (right - left)/2
        if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1]{
            return mid
        } else if nums[mid] > nums[mid + 1] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}
