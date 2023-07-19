package binary

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
	"errors"
)

// TODO: add solution stub
func ParseBinary(binary string) (expected int, err error) {
	if strings.IndexFunc(binary, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsSpace(r) || r > '1'
	}) != -1 {
		err = errors.New(`Invalid input, contains space characters, 
		letters or invalid numbers`)
		return
	}

	i := len(binary) - 1
	for _, ch := range binary {
		n, err := strconv.Atoi(string(ch))
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		expected += n * int(math.Pow(float64(2), float64(i)))
		i--
	}
	
	return
}
