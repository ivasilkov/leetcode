package _0599__Minimum_Index_Sum_of_Two_Lists

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i, v := range list1 {
		m[v] = i
	}

	min := len(list1) + len(list2)
	var out []string
	for i := range list2 {
		v, ok := m[list2[i]]
		if !ok {
			continue
		}
		if min > v+i {
			min = v + i
			out = []string{list2[i]}
		} else if min == v+i {
			out = append(out, list2[i])
		}
	}

	return out
}
