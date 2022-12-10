package common

import (
	"fmt"
	"os"
	"strings"
)

func ReadLines(day string) []string {
	data, _ := os.ReadFile(fmt.Sprintf("%s/%s.txt", day, day))

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
