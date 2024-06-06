func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 {
        return false
    }

    mp := make(map[int]int)
    for i := 0; i < len(hand); i++ {
        mp[hand[i]]++
    }

    keys := make([]int, 0, len(mp))
    for key := range mp {
        keys = append(keys, key)
    }

    sort.Ints(keys)

    for _, key := range keys {
        count := mp[key]
        if count > 0 {
            for i := 0; i < groupSize; i++ {
                if mp[key + i] > 0 {
                    mp[key]--
                } else {
                    return false
                }
            }
        }
    }

    return true
}
