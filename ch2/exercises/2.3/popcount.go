package popcount

func PopCount(x uint64) int {
	cnt := 0
	for i := uint64(0); i < 64; i++ {
		if (x & (1 << i)) > 0 {
			cnt++
		}
	}
	return cnt
}
