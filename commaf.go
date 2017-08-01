// +build go1.6

package humanize

import (
	"bytes"
	"math/big"
	"strings"
)

// BigCommaf produces a string form of the given big.Float in base 10
// with commas after every three orders of magnitude.
func BigCommaf(v *big.Float) string {
	return CustomBigCommaf(v, -1, ',', '.')
}

// CustomBigCommaf produces a string form of the given big.Float in base 10
// with given rune after every three orders of magnitude,
// as well as the provided rune for the decimal point.
// Provide non-negative number of decimals for a fixed format, or -1 otherwise.
func CustomBigCommaf(v *big.Float, decimals int, kilo, point rune) string {
	buf := &bytes.Buffer{}
	if v.Sign() < 0 {
		buf.Write([]byte{'-'})
		v.Abs(v)
	}

	parts := strings.Split(v.Text('f', decimals), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.WriteRune(kilo)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.WriteRune(kilo)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.WriteRune(point)
		buf.WriteString(parts[1])
	}
	return buf.String()
}
