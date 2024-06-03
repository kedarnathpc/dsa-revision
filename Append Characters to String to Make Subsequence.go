func appendCharacters(s string, t string) int {
    i, j, m, n := 0, 0, len(s), len(t);

    for i < m && j < n {
        if (s[i] == t[j]) {
            j++
        }
        i++
    }

    return n - j
}
