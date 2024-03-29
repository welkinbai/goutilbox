package collections

func FindDupElement[E comparable](slice []E) []E {
	checkMap := make(map[E]bool)
	result := make([]E, 0, len(slice))
	for _, one := range slice {
		if _, ok := checkMap[one]; ok {
			checkMap[one] = true
		} else {
			checkMap[one] = false
		}
	}
	for _, one := range slice {
		if checkMap[one] {
			result = append(result, one)
		}
	}
	return result
}

func FindDupElementInUnComparable[E any, O comparable](slice []E, compareElementFunc func(e E) O) []E {
	checkMap := make(map[O]bool)
	result := make([]E, 0, len(slice))
	for _, one := range slice {
		elementCanCompare := compareElementFunc(one)
		if _, ok := checkMap[elementCanCompare]; ok {
			checkMap[elementCanCompare] = true
		} else {
			checkMap[elementCanCompare] = false
		}
	}
	for _, one := range slice {
		elementCanCompare := compareElementFunc(one)
		if checkMap[elementCanCompare] {
			result = append(result, one)
		}
	}
	return result
}

func GroupListToMap[K comparable, T any](list []T, keyFunc func(t T) K) map[K][]T {
	result := make(map[K][]T, len(list))
	for _, value := range list {
		k := keyFunc(value)
		mapList, ok := result[k]
		if ok {
			mapList = append(mapList, value)
			result[k] = mapList
		} else {
			result[k] = []T{value}
		}
	}
	return result
}
