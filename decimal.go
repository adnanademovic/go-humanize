// +build go1.6

package humanize

import (
	"bytes"
	"github.com/shopspring/decimal"
	"strings"
)

// Decimal produces a string form of the given decimal.Decimal in base 10
// with commas after every three orders of magnitude.
func Decimal(v decimal.Decimal) string {
	return CustomDecimal(v, -1, ',', '.')
}

// CustomDecimal produces a string form of the given decimal.Decimal in base 10
// with given rune after every three orders of magnitude,
// as well as the provided rune for the decimal point.
// Provide non-negative number of decimals for a fixed format, or -1 otherwise.
func CustomDecimal(v decimal.Decimal, decimals int32, kilo, point rune) string {
	buf := &bytes.Buffer{}
	if v.Sign() < 0 {
		buf.Write([]byte{'-'})
		v = v.Abs()
	}
	text := ""
	if decimals < 0 {
		text = v.String()
	} else {
		text = v.StringFixed(decimals)
	}
	parts := strings.Split(text, ".")
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
