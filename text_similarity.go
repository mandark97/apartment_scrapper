package main

// Max .
func Max(more ...int) int {
	manNum := more[0]
	for _, elem := range more {
		if manNum < elem {
			manNum = elem
		}
	}
	return manNum
}

// LCS Longest common subsequence
func LCS(str1, str2 string) float64 {
	len1 := len(str1)
	len2 := len(str2)

	tab := make([][]int, len1+1)
	for i := range tab {
		tab[i] = make([]int, len2+1)
	}

	i, j := 0, 0
	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				tab[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				tab[i][j] = tab[i-1][j-1] + 1
				if i < len1 {
					i++
					j++
				}
			} else {
				tab[i][j] = Max(tab[i-1][j], tab[i][j-1])
			}
		}
	}

	return float64(tab[len1][len2]) / float64((len1+len2)/2)
}
