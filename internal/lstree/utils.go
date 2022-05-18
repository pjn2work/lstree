package lstree

func fillWith(char string, n int) string {
    res := ""
    for ; n > 0; n-- {
        res += char
    }
    return res
}

func getStringLen(str string) int {
    r := []rune(str)
    return len(r)
}
