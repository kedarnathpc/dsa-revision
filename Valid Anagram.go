func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    slice1 := make([]int, 26)
    slice2 := make([]int, 26)

    for i := 0; i < len(s); i++ {
        slice1[s[i]-'a'] = slice1[s[i]-'a'] + 1
        slice2[t[i]-'a'] = slice2[t[i]-'a'] + 1
    }

    for i := 0; i < 26; i++ {
        if slice1[i] != slice2[i] {
            return false
        }
    }

    return true
}
