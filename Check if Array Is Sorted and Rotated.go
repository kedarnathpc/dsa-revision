func check(nums []int) bool {
    count := 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            count++
        }
    }

    if nums[len(nums)-1] > nums[0] {
        count++
    }
    
    return count <= 1
}
