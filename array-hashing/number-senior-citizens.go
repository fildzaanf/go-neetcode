package arrayhashing

import "strconv"

func CountSeniors(details []string) int {
	count := 0
	for word := 0; word < len(details); word++ {
		age, _ := strconv.Atoi(string(details[word][11]))
		if age > 6 {
			count++
		}
	}
	return count
}

