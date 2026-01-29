package arrayhashing

func ScoreOfString(s string) int {
    total := 0

    for i := 0; i < len(s)-1; i++ {
        diff := int(s[i+1]) - int(s[i])
        if diff < 0 {
            diff = -diff
        }
        total += diff
    }

    return total
}
