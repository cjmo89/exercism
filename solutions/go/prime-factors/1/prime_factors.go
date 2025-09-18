package prime

func Factors(n int64) []int64 {
	results := []int64{}
	var i int64
	for i = 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			results = append(results, i)
		}
	}
	return results
}
