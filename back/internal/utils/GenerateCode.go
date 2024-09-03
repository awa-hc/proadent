package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateCode(code string) (string, error) {
	parts := strings.Split(code, "-")
	fmt.Println(code, "testing...")

	if len(parts) != 2 {
		return "", &ValidationError{Field: "code", Message: "invalid code format."}
	}

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", &ValidationError{Field: "code", Message: "invalid code format."}
	}
	NewCode := parts[0] + "-" + strconv.Itoa(num+1)

	return NewCode, nil
}
