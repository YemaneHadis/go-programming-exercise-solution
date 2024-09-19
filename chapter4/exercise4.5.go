// Write an inplace function to eliminate adjacent duplicates in a[]string slice

package main

func eliminateDuplicate(strings []string) []string {
	j := 0
	for i := 0; i < len(strings)-1; i++ {
		if i == 0 {
			j++
			continue
		}
		if strings[i] != strings[i-1] {
			strings[j] = strings[i]
			j++
		}

	}
	return strings[:j+1]
}
