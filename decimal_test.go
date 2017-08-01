// +build go1.6

package humanize

import (
	"github.com/shopspring/decimal"
	"math"
	"testing"
)

func BenchmarkDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Commaf(1234567890.83584)
	}
}

func TestDecimals(t *testing.T) {
	testList{
		{"0", Decimal(decimal.NewFromFloat(0)), "0"},
		{"10.11", Decimal(decimal.NewFromFloat(10.11)), "10.11"},
		{"100", Decimal(decimal.NewFromFloat(100)), "100"},
		{"1,000", Decimal(decimal.NewFromFloat(1000)), "1,000"},
		{"10,000", Decimal(decimal.NewFromFloat(10000)), "10,000"},
		{"100,000", Decimal(decimal.NewFromFloat(100000)), "100,000"},
		{"834,142.32", Decimal(decimal.NewFromFloat(834142.32)), "834,142.32"},
		{"10,000,000", Decimal(decimal.NewFromFloat(10000000)), "10,000,000"},
		{"10,100,000", Decimal(decimal.NewFromFloat(10100000)), "10,100,000"},
		{"10,010,000", Decimal(decimal.NewFromFloat(10010000)), "10,010,000"},
		{"10,001,000", Decimal(decimal.NewFromFloat(10001000)), "10,001,000"},
		{"123,456,789", Decimal(decimal.NewFromFloat(123456789)), "123,456,789"},
		{"maxf64", Decimal(decimal.NewFromFloat(math.MaxFloat64)), "179,769,313,486,231,570,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000"},
		{"-123,456,789", Decimal(decimal.NewFromFloat(-123456789)), "-123,456,789"},
		{"-10,100,000", Decimal(decimal.NewFromFloat(-10100000)), "-10,100,000"},
		{"-10,010,000", Decimal(decimal.NewFromFloat(-10010000)), "-10,010,000"},
		{"-10,001,000", Decimal(decimal.NewFromFloat(-10001000)), "-10,001,000"},
		{"-10,000,000", Decimal(decimal.NewFromFloat(-10000000)), "-10,000,000"},
		{"-100,000", Decimal(decimal.NewFromFloat(-100000)), "-100,000"},
		{"-10,000", Decimal(decimal.NewFromFloat(-10000)), "-10,000"},
		{"-1,000", Decimal(decimal.NewFromFloat(-1000)), "-1,000"},
		{"-100.11", Decimal(decimal.NewFromFloat(-100.11)), "-100.11"},
		{"-10", Decimal(decimal.NewFromFloat(-10)), "-10"},
	}.validate(t)
}

func TestCustomDecimals(t *testing.T) {
	testList{
		{"0", CustomDecimal(decimal.NewFromFloat(0), -1, '.', ','), "0"},
		{"10.11", CustomDecimal(decimal.NewFromFloat(10.11), -1, '.', ','), "10,11"},
		{"100", CustomDecimal(decimal.NewFromFloat(100), -1, '.', ','), "100"},
		{"1,000", CustomDecimal(decimal.NewFromFloat(1000), -1, '.', ','), "1.000"},
		{"10,000", CustomDecimal(decimal.NewFromFloat(10000), -1, '.', ','), "10.000"},
		{"100,000", CustomDecimal(decimal.NewFromFloat(100000), -1, '.', ','), "100.000"},
		{"834,142.32", CustomDecimal(decimal.NewFromFloat(834142.32), -1, '.', ' '), "834.142 32"},
		{"10,000,000", CustomDecimal(decimal.NewFromFloat(10000000), -1, '^', ','), "10^000^000"},
		{"10,100,000", CustomDecimal(decimal.NewFromFloat(10100000), -1, '%', ','), "10%100%000"},
		{"10,010,000", CustomDecimal(decimal.NewFromFloat(10010000), -1, '+', ','), "10+010+000"},
		{"10,001,000", CustomDecimal(decimal.NewFromFloat(10001000), -1, '/', ','), "10/001/000"},
		{"123,456,789", CustomDecimal(decimal.NewFromFloat(123456789), -1, '.', ','), "123.456.789"},
		{"maxf64", CustomDecimal(decimal.NewFromFloat(math.MaxFloat64), -1, ' ', ','), "179 769 313 486 231 570 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000"},
		{"minf64 fixed", CustomDecimal(decimal.NewFromFloat(math.SmallestNonzeroFloat64), 4, '.', '|'), "0|0000"},
		{"-123,456,789", CustomDecimal(decimal.NewFromFloat(-123456789), -1, ':', ','), "-123:456:789"},
		{"-10,100,000", CustomDecimal(decimal.NewFromFloat(-10100000), -1, '|', ','), "-10|100|000"},
		{"-10,010,000", CustomDecimal(decimal.NewFromFloat(-10010000), -1, '*', ','), "-10*010*000"},
		{"-10,001,000", CustomDecimal(decimal.NewFromFloat(-10001000), -1, '?', ','), "-10?001?000"},
		{"-10,000,000", CustomDecimal(decimal.NewFromFloat(-10000000), -1, '!', ','), "-10!000!000"},
		{"-100,000", CustomDecimal(decimal.NewFromFloat(-100000), -1, '.', ','), "-100.000"},
		{"-10,000", CustomDecimal(decimal.NewFromFloat(-10000), -1, '.', ','), "-10.000"},
		{"-1,000", CustomDecimal(decimal.NewFromFloat(-1000), -1, '.', ','), "-1.000"},
		{"-100.11", CustomDecimal(decimal.NewFromFloat(-100.11), -1, '.', ','), "-100,11"},
		{"-10", CustomDecimal(decimal.NewFromFloat(-10), -1, '.', ','), "-10"},
		{"834,142.3211 fixed to 2", CustomDecimal(decimal.NewFromFloat(834142.3211), 2, '.', ' '), "834.142 32"},
		{"834,142.3211 fixed to 0", CustomDecimal(decimal.NewFromFloat(834142.3211), 0, '.', ' '), "834.142"},
		{"834,142.3249", CustomDecimal(decimal.NewFromFloat(834142.3249), -1, '.', ' '), "834.142 3249"},
		{"834,142.3249 fixed to 2 rounds down", CustomDecimal(decimal.NewFromFloat(834142.3249), 2, '.', ' '), "834.142 32"},
		{"834,142.3250", CustomDecimal(decimal.NewFromFloat(834142.3250), -1, '.', ' '), "834.142 325"},
		{"834,142.3250 fixed to 2 rounds up", CustomDecimal(decimal.NewFromFloat(834142.3250), 2, '.', ' '), "834.142 33"},
		{"834,142.3251", CustomDecimal(decimal.NewFromFloat(834142.3251), -1, '.', ' '), "834.142 3251"},
		{"834,142.3251 fixed to 2 rounds up", CustomDecimal(decimal.NewFromFloat(834142.3251), 2, '.', ' '), "834.142 33"},
		{"834,143.3249", CustomDecimal(decimal.NewFromFloat(834143.3249), -1, '.', ' '), "834.143 3249"},
		{"834,143.3249 fixed to 2 rounds down", CustomDecimal(decimal.NewFromFloat(834143.3249), 2, '.', ' '), "834.143 32"},
		{"834,143.3250", CustomDecimal(decimal.NewFromFloat(834143.3250), -1, '.', ' '), "834.143 325"},
		{"834,143.3250 fixed to 2 rounds up", CustomDecimal(decimal.NewFromFloat(834143.3250), 2, '.', ' '), "834.143 33"},
		{"834,143.3251", CustomDecimal(decimal.NewFromFloat(834143.3251), -1, '.', ' '), "834.143 3251"},
		{"834,143.3251 fixed to 2 rounds up", CustomDecimal(decimal.NewFromFloat(834143.3251), 2, '.', ' '), "834.143 33"},
	}.validate(t)
}
