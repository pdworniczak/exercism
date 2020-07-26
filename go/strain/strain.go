package strain

// Ints enchanted collection of int values
type Ints []int

// Lists enchanted collection of int collection values
type Lists [][]int

// Strings enchanted collection of string values
type Strings []string

// Keep funciton return collection of values that fulfill predicate
func (col Ints) Keep(pred func(int) bool) Ints {
	var result Ints

	for _, val := range col {
		if pred(val) {
			result = append(result, val)
		}
	}

	return result
}

// Discard function return collection filtered from values that fillfill predicate
func (col Ints) Discard(pred func(int) bool) Ints {
	var result Ints

	for _, val := range col {
		if !pred(val) {
			result = append(result, val)
		}
	}
	return result
}

// Keep funciton return collection of values that fulfill predicate
func (list Lists) Keep(pred func([]int) bool) Lists {
	var result Lists

	for _, val := range list {
		if pred(val) {
			result = append(result, val)
		}
	}

	return result
}

// Keep funciton return collection of values that fulfill predicate
func (str Strings) Keep(pred func(string) bool) Strings {
	var result Strings

	for _, val := range str {
		if pred(val) {
			result = append(result, val)
		}
	}

	return result
}
