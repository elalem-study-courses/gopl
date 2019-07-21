package popcount

import "sync"

var once sync.Once
var pc [256]byte

// func init() {

// }

func initializePC() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	once.Do(initializePC)
	return int(pc[byte(x)] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
