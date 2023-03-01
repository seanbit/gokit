package utils

import (
	"math/rand"
	"time"
)

type Rand struct {
	rd *rand.Rand
	offset int64
	seedTime time.Duration
	seedDuration time.Duration
}

func NewRand(offset int64, seedDuration time.Duration) *Rand {
	tm := time.Now()
	return &Rand{
		rd:           rand.New(rand.NewSource(tm.UnixNano() + offset)),
		offset:       offset,
		seedTime:     time.Duration(tm.Unix()),
		seedDuration: seedDuration,
	}
}

func (r *Rand) RandInt(n int) int {
	if time.Duration(time.Now().Unix()) - r.seedTime > r.seedDuration {
		tm := time.Now()
		r.rd.Seed(tm.UnixNano() + r.offset)
		r.seedTime = time.Duration(tm.Unix())
	}
	return r.rd.Intn(n)
}