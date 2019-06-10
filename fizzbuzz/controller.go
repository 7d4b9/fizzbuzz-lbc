// Package fizzbuzz provides the fizzbuzz controller
package fizzbuzz

import (
	"strconv"
	"strings"
)

// Controller implements the FizzBuzz controller features
type Controller struct {
	// ressources....., databases, queues,....
}

// FizzBuzz computes a FizzBuzz on requirement
// all multiples of int1 are replaced by str1,
// all multiples of int2 are replaced by str2,
// all multiples of int1 and int2 are replaced by str1str2.
func (app *Controller) FizzBuzz(int1, int2, limit int, str1, str2 string) (string, error) {
	words := []string{}
	for i := 1; i <= limit; i++ {
		word := ""
		if i%int1 == 0 {
			word += str1
		}
		if i%int2 == 0 {
			word += str2
		}
		if word == "" {
			word += strconv.Itoa(i)
		}
		words = append(words, word)
	}
	return strings.Join(words, " "), nil
}
