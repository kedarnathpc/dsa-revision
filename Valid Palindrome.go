func checkPalindrome(s string, left, right int) bool {
    if left >= right {
        return true
    }
    if s[left] != s[right] { 
        return false
    }
    return checkPalindrome(s, left + 1, right - 1)
}

func isPalindrome(s string) bool {
    str := ""
    for _, c := range s {
        if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
            str += strings.ToLower(string(c))
        }
    }
    
    return checkPalindrome(str, 0, len(str) - 1)
}
