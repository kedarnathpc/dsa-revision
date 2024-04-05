// Merge Sort

func merge(arr []int, low, mid, high int) {
    var temp[]int
    left, right := low, mid + 1

    for left <= mid && right <= high {
        if arr[left] <= arr[right] {
            temp = append(temp, arr[left])
            left++
        } else {
            temp = append(temp, arr[right])
            right++
        }
    }

    for left <= mid {
        temp = append(temp, arr[left])
        left++
    }

    for right <= high {
        temp = append(temp, arr[right])
        right++
    }

    for i := low; i <= high; i++ {
        arr[i] = temp[i - low]
    }
}

func arraySort(nums []int, low,high int) {
    if low >= high{
        return
    }

    mid := low + (high - low)/2;

    arraySort(nums, low, mid)
    arraySort(nums, mid + 1, high)
    merge(nums, low, mid, high)
}

func sortArray(nums []int) []int {
    arraySort(nums, 0, len(nums)-1)
    return nums
}
