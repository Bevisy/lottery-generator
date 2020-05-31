package common

import (
	"fmt"
)

func Help() string {
	return fmt.Sprint("---HELP---\n\n" +
		"Lottery Generator:\n" +
		"    Method: GET\n" +
		"    Path: /lottery\n" +
		"    Return Code: 200\n" +
		"    Example Return: 01 05 12 13 18 23 | 02\n\n" +
		"Blog:\n" +
		"    Method: GET\n" +
		"    Path: /blog\n" +
		"    Return Code: 200\n" +
		"    Example Return: new sites\n")
}
