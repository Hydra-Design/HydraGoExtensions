package hge

type NoElementMatchingPredicate struct{}

func (e *NoElementMatchingPredicate) Error() string {
	return "no element matching the given predicate"
}

func ForEach[Element any](slice []Element, run func(Element)) {
	for _, curElement := range slice {
		run(curElement)
	}
}

func Delete[Element comparable](array []Element, item Element) []Element {
	for index, curItem := range array {
		if curItem == item {
			return append(array[:index], array[index+1:]...)
		}
	}
	return array
}

func DeleteMatching[Element any](array []Element, predicate func(Element) bool) []Element {
	for index, curItem := range array {
		if predicate(curItem) {
			return append(array[:index], array[index+1:]...)
		}
	}
	return array
}

func FirstWhere[Element any](slice []Element, predicate func(Element) bool) (e Element, err error) {
	for _, curElement := range slice {
		if predicate(curElement) {
			e = curElement
			return
		}
	}
	err = &NoElementMatchingPredicate{}
	return
}

func LastWhere[Element any](slice []Element, predicate func(Element) bool) (e Element, err error) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			e = slice[i]
			return
		}
	}
	err = &NoElementMatchingPredicate{}
	return
}

func Map[Element any, Output any](slice []Element, run func(Element) Output) []Output {
	var out []Output
	for _, curElement := range slice {
		out = append(out, run(curElement))
	}
	return out
}

func Contains[Element comparable](slice []Element, item Element) bool {
	for _, curElement := range slice {
		if curElement == item { return true }
	}
	return false
}

func ContainsWhere[Element any](slice []Element, predicate func(Element) bool) bool {
	for _, curElement := range slice {
		if predicate(curElement) { return true }
	}
	return false
}