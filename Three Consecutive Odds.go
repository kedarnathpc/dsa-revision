func threeConsecutiveOdds(arr []int) bool {
    count := 0

    for _, n := range arr {
        if n & 1 == 1 {
            count++
            if count == 3 {
                return true
            }
        } else {
            count = 0
        }
    }

    return false
}
