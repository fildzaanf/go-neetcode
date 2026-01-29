package arrayhashing

func LengthOfLastWord(s string) int {
    count := 0
    i := len(s) - 1

    for i >= 0 && s[i] == ' '{
        i--
    }

    for i >= 0 && s[i] != ' '{
        count++
        i--
    }

    return count
}