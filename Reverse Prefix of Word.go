func reverse(word []byte, i, j int) {
    for i <= j {
        word[i], word[j] = word[j], word[i]
        i++
        j--
    }
}

func reversePrefix(word string, ch byte) string {
    wordByte := []byte(word)
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            reverse(wordByte, 0, i)
            break
        }
    }

    return string(wordByte)
}
