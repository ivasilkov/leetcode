package _1657__Determine_if_Two_Strings_Are_Close_

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	counter1, counter2 := [26]int{}, [26]int{}
	for i := range word1 {
		counter1[word1[i]-'a']++
		counter2[word2[i]-'a']++
	}

	for i := range counter1 {
		if counter1[i] != 0 && counter2[i] == 0 {
			return false
		}
	}

	m := map[int]int{}
	for i := range counter1 {
		m[counter1[i]]++
		m[counter2[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true

	//loop:
	//for i1, v1 := range counter1 {
	//	for i2, v2 := range counter2 {
	//		if v1 == v2 {
	//			counter2[i2] = 0
	//			counter1[i1] = 0
	//			continue loop
	//		}
	//	}
	//	return false
	//}
	//return true
}
