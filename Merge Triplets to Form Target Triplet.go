func mergeTriplets(triplets [][]int, target []int) bool {
    a, b, c := false, false, false

    for i := range triplets {
        if triplets[i][0] <= target[0] && triplets[i][1] <= target[1] && triplets[i][2] <= target[2] {
            if target[0] == triplets[i][0] {
                a = true
            }
            if target[1] == triplets[i][1] {
                b = true
            }
            if target[2] == triplets[i][2] {
                c = true
            }
        }
    } 

    return a && b && c
}
