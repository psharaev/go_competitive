package polynomial_hash

type PolynomialHash struct {
	// hashes p[i] = hash(s[0..i-1]) = s0 * x^i-1 + s1 * x^i-2
	hashes []int
	pows   []int
	prime  int
	n      int
}

// New
// base мощность алфавита или любое число больше него
// prime большое случайное простое число
// runeConverter получение порядкового номера символа в алфавите
func New(arr []int, base int, prime int) *PolynomialHash {
	n := len(arr) + 1
	hashes := make([]int, n)
	pows := make([]int, n)
	pows[0] = 1

	for i := 1; i < n; i++ {
		hashes[i] = (hashes[i-1]*base + arr[i-1]) % prime
		pows[i] = (pows[i-1] * base) % prime
	}

	return &PolynomialHash{
		hashes: hashes,
		pows:   pows,
		prime:  prime,
		n:      len(arr),
	}
}

func NewDefault(arr []int) *PolynomialHash {
	return New(arr, 37, 1_000_000_007)
}

// Hash
// l включительно
// r не включительно
func (s *PolynomialHash) Hash(l int, r int) int {
	return (s.hashes[r] - (s.hashes[l]*s.pows[r-l])%s.prime + s.prime) % s.prime
}

func (s *PolynomialHash) HashAll() int {
	return s.Hash(0, s.n)
}
