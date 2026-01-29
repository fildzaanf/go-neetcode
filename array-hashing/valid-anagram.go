package arrayhashing

func IsAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
  
    countS, countT := make(map[rune]int), make(map[rune]int)

    for i, v := range s {
        countS[v]++
        countT[rune(t[i])]++
    }

    for key, value := range countS {
        if value != countT[key] {
            return false
        }
    }

    return true 
}