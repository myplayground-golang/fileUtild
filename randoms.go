package fileUtild

import (
	"math/rand"
	"time"
)

// RandomString generate a random string by required length
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed( time.Now().UTC().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// RandomInt generate a random int which is in range 0 ~ max
func RandomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// RandomIntNoRange get a random int with no range scale
func RandomIntNoRange() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

// RandomIntWithRange get a random int with range scale by min ~ max
func RandomIntWithRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// RandomFloatNoRange get a random float with no range scale
func RandomFloatNoRange() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

// RandomIntList get a list of random int with no range scale
func RandomIntList() []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(8)
}