package tool

import "strconv"

func StrAutoIncrease(s string) (string, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}
	num++
	return strconv.Itoa(num), nil
}
