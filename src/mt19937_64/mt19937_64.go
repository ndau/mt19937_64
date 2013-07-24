// Copyright (c) 2012-2013 Bartosz Szczesny
// LICENSE: The BSD 2-Clause License

/*
	note: "unsigned long long" = "uint64" therefore "genrand_int64()" --> "Uint64()"
*/

package mt19937_64

import "sync"

const (
	MT_VEC_LEN  uint64 = 312
	MT_VEC_LEN2 uint64 = 156

	MT_MAG_LEN uint64 = 2
	MT_MAG_0   uint64 = 0
	MT_MAG_1   uint64 = 0xB5026F5AA96619E9

	MT_MOST_33  uint64 = 0xFFFFFFFF80000000
	MT_LEAST_31 uint64 = 0x000000007FFFFFFF

	MT_DEFAULT_SEED       uint64 = 5489
	MT_DEFAULT_SEED_ARRAY uint64 = 19650218

	MT_INIT         uint64 = 6364136223846793005
	MT_INIT_ARRAY_J uint64 = 3935559000370003845
	MT_INIT_ARRAY_I uint64 = 2862933555777941757

	MT_UINT64_SHR_29 uint64 = 0x5555555555555555
	MT_UINT64_SHL_17 uint64 = 0x71D67FFFEDA60000
	MT_UINT64_SHL_37 uint64 = 0xFFF7EEE000000000

	MT_REAL1_DIV float64 = 9007199254740991.0
	MT_REAL2_DIV float64 = 9007199254740992.0
	MT_REAL3_DIV float64 = 4503599627370496.0
)

type MT struct {
	index uint64
	mutex sync.Mutex
	vec   [MT_VEC_LEN]uint64
	mag   [MT_MAG_LEN]uint64
}

func New() *MT {
	return &MT{index: MT_VEC_LEN + 1}
}

func (mt *MT) initLocked(seed uint64) {
	mt.mag[0] = MT_MAG_0
	mt.mag[1] = MT_MAG_1

	mt.vec[0] = seed
	for mt.index = 1; mt.index < MT_VEC_LEN; mt.index++ {
		mt.vec[mt.index] = MT_INIT*(mt.vec[mt.index-1]^(mt.vec[mt.index-1]>>62)) + mt.index
	}
}

func (mt *MT) Init(seed uint64) {
	mt.mutex.Lock()
	mt.initLocked(seed)
	mt.mutex.Unlock()
}

func (mt *MT) InitByArray(initKey []uint64) {
	mt.mutex.Lock()
	mt.initLocked(MT_DEFAULT_SEED_ARRAY)

	var initKeyLen uint64 = uint64(len(initKey))
	var i uint64 = 1
	var j uint64 = 0
	var k uint64 = 0

	if MT_VEC_LEN > initKeyLen {
		k = MT_VEC_LEN
	} else {
		k = initKeyLen
	}

	for ; k != 0; k-- {
		mt.vec[i] = (mt.vec[i] ^ ((mt.vec[i-1] ^ (mt.vec[i-1] >> 62)) * MT_INIT_ARRAY_J)) + j + initKey[j]
		i++
		j++
		if i >= MT_VEC_LEN {
			mt.vec[0] = mt.vec[MT_VEC_LEN-1]
			i = 1
		}
		if j >= initKeyLen {
			j = 0
		}
	}

	for k = MT_VEC_LEN - 1; k != 0; k-- {
		mt.vec[i] = (mt.vec[i] ^ ((mt.vec[i-1] ^ (mt.vec[i-1] >> 62)) * MT_INIT_ARRAY_I)) - i
		i++
		if i >= MT_VEC_LEN {
			mt.vec[0] = mt.vec[MT_VEC_LEN-1]
			i = 1
		}
	}

	mt.vec[0] = 1 << 63
	mt.mutex.Unlock()
}

func (mt *MT) Uint64() uint64 {
	mt.mutex.Lock()
	if MT_VEC_LEN+1 == mt.index {
		mt.initLocked(MT_DEFAULT_SEED)
	}

	var x uint64
	var j uint64
	if MT_VEC_LEN <= mt.index {
		mt.index = 0

		for j = 0; j < MT_VEC_LEN-MT_VEC_LEN2; j++ {
			x = (mt.vec[j] & MT_MOST_33) | (mt.vec[j+1] & MT_LEAST_31)
			mt.vec[j] = mt.vec[j+MT_VEC_LEN2] ^ (x >> 1) ^ mt.mag[x&1]
		}

		for ; j < MT_VEC_LEN-1; j++ {
			x = (mt.vec[j] & MT_MOST_33) | (mt.vec[j+1] & MT_LEAST_31)
			mt.vec[j] = mt.vec[j+MT_VEC_LEN2-MT_VEC_LEN] ^ (x >> 1) ^ mt.mag[x&1]
		}

		x = (mt.vec[MT_VEC_LEN-1] & MT_MOST_33) | (mt.vec[0] & MT_LEAST_31)
		mt.vec[MT_VEC_LEN-1] = mt.vec[MT_VEC_LEN2-1] ^ (x >> 1) ^ mt.mag[x&1]
	}

	x = mt.vec[mt.index]
	x ^= (x >> 29) & MT_UINT64_SHR_29
	x ^= (x << 17) & MT_UINT64_SHL_17
	x ^= (x << 37) & MT_UINT64_SHL_37
	x ^= (x >> 43)

	mt.index++
	mt.mutex.Unlock()
	return x
}

func (mt *MT) Int63() int64 {
	return int64(mt.Uint64() >> 1)
}

func (mt *MT) Real1() float64 {
	return float64(mt.Uint64()>>11) / MT_REAL1_DIV
}

func (mt *MT) Real2() float64 {
	return float64(mt.Uint64()>>11) / MT_REAL2_DIV
}

func (mt *MT) Real3() float64 {
	return (float64(mt.Uint64()>>12) + 0.5) / MT_REAL3_DIV
}
