// Insertion Sort
func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        key, j := nums[i], i-1
        
        for j >= 0 && nums[j] > key {
            nums[j+1] = nums[j]
            j--
        }

        nums[j+1] = key
    }

    return nums
}
