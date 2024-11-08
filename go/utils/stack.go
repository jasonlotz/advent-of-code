package utils

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) PopMultiple(count int) ([]string, bool) {
	if len(*s) < count {
		return nil, false
	}
	index := len(*s) - count
	elements := (*s)[index:]
	*s = (*s)[:index]
	return elements, true
}

func (s *Stack) Peek() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	return (*s)[len(*s)-1], true
}
