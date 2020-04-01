package test_demo

import "fmt"

func byAdd(s1 string, s2 string) string {
	return s1 + " " + s2
}

func bySprintf(s1 string, s2 string) string{
	return fmt.Sprintf("%s %s", s1, s2 )
}
