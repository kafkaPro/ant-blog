package utils

func FuzzyAll(s string) string {
	return "%" + s + "%"
}

func FuzzyLeft(s string) string {
	return "%" + s
}

func FuzzyRight(s string) string {
	return s + "%"
}
