package utils

import (
	"fmt"
	"github.com/beewit/beekit/utils/convert"
	"github.com/beewit/beekit/utils/ulid"
	"math"
	"math/rand"
	"time"
)

// Random random
type Random struct{}

// NewRandom new random
func NewRandom() *Random {
	return &Random{}
}

// Number random number
func (rd Random) Number(length float64) int {
	i := int(math.Pow(10, length-1))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9*i) + i
}

// String random string
func (rd Random) String(length int) string {
	b := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	d := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		d = append(d, b[r.Intn(62)])
	}
	return string(d)
}

// ULID random ulid
func (rd Random) ULID() string {
	t := time.Now()
	entrop := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entrop).String()
}

func (rd Random) NumberByFloat(start, end float64) float64 {
	r := rand.New(rand.NewSource(ID()))
	min := convert.MustInt(start * 100)
	max := convert.MustInt(end * 100)
	n := min + r.Intn(max-min+1)
	f := float64(n) / float64(100)
	return convert.MustFloat64(fmt.Sprintf("%.2f", f))
}

func (rd Random) NumberByInt(start, end int) int {
	r := rand.New(rand.NewSource(ID()))
	return start + r.Intn(end-start+1)
}
