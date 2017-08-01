// +build go1.6

package humanize

import (
	"math"
	"math/big"
	"testing"
)

func BenchmarkBigCommaf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Commaf(1234567890.83584)
	}
}

func TestBigCommafs(t *testing.T) {
	testList{
		{"0", BigCommaf(big.NewFloat(0)), "0"},
		{"10.11", BigCommaf(big.NewFloat(10.11)), "10.11"},
		{"100", BigCommaf(big.NewFloat(100)), "100"},
		{"1,000", BigCommaf(big.NewFloat(1000)), "1,000"},
		{"10,000", BigCommaf(big.NewFloat(10000)), "10,000"},
		{"100,000", BigCommaf(big.NewFloat(100000)), "100,000"},
		{"834,142.32", BigCommaf(big.NewFloat(834142.32)), "834,142.32"},
		{"10,000,000", BigCommaf(big.NewFloat(10000000)), "10,000,000"},
		{"10,100,000", BigCommaf(big.NewFloat(10100000)), "10,100,000"},
		{"10,010,000", BigCommaf(big.NewFloat(10010000)), "10,010,000"},
		{"10,001,000", BigCommaf(big.NewFloat(10001000)), "10,001,000"},
		{"123,456,789", BigCommaf(big.NewFloat(123456789)), "123,456,789"},
		{"maxf64", BigCommaf(big.NewFloat(math.MaxFloat64)), "179,769,313,486,231,570,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000"},
		{"minf64", BigCommaf(big.NewFloat(math.SmallestNonzeroFloat64)), "0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004940656458412465"},
		{"-123,456,789", BigCommaf(big.NewFloat(-123456789)), "-123,456,789"},
		{"-10,100,000", BigCommaf(big.NewFloat(-10100000)), "-10,100,000"},
		{"-10,010,000", BigCommaf(big.NewFloat(-10010000)), "-10,010,000"},
		{"-10,001,000", BigCommaf(big.NewFloat(-10001000)), "-10,001,000"},
		{"-10,000,000", BigCommaf(big.NewFloat(-10000000)), "-10,000,000"},
		{"-100,000", BigCommaf(big.NewFloat(-100000)), "-100,000"},
		{"-10,000", BigCommaf(big.NewFloat(-10000)), "-10,000"},
		{"-1,000", BigCommaf(big.NewFloat(-1000)), "-1,000"},
		{"-100.11", BigCommaf(big.NewFloat(-100.11)), "-100.11"},
		{"-10", BigCommaf(big.NewFloat(-10)), "-10"},
	}.validate(t)
}

func TestCustomBigCommafs(t *testing.T) {
	testList{
		{"0", CustomBigCommaf(big.NewFloat(0), -1, '.', ','), "0"},
		{"10.11", CustomBigCommaf(big.NewFloat(10.11), -1, '.', ','), "10,11"},
		{"100", CustomBigCommaf(big.NewFloat(100), -1, '.', ','), "100"},
		{"1,000", CustomBigCommaf(big.NewFloat(1000), -1, '.', ','), "1.000"},
		{"10,000", CustomBigCommaf(big.NewFloat(10000), -1, '.', ','), "10.000"},
		{"100,000", CustomBigCommaf(big.NewFloat(100000), -1, '.', ','), "100.000"},
		{"834,142.32", CustomBigCommaf(big.NewFloat(834142.32), -1, '.', ' '), "834.142 32"},
		{"10,000,000", CustomBigCommaf(big.NewFloat(10000000), -1, '^', ','), "10^000^000"},
		{"10,100,000", CustomBigCommaf(big.NewFloat(10100000), -1, '%', ','), "10%100%000"},
		{"10,010,000", CustomBigCommaf(big.NewFloat(10010000), -1, '+', ','), "10+010+000"},
		{"10,001,000", CustomBigCommaf(big.NewFloat(10001000), -1, '/', ','), "10/001/000"},
		{"123,456,789", CustomBigCommaf(big.NewFloat(123456789), -1, '.', ','), "123.456.789"},
		{"maxf64", CustomBigCommaf(big.NewFloat(math.MaxFloat64), -1, ' ', ','), "179 769 313 486 231 570 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000 000"},
		{"minf64", CustomBigCommaf(big.NewFloat(math.SmallestNonzeroFloat64), -1, '.', '|'), "0|000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004940656458412465"},
		{"minf64 fixed", CustomBigCommaf(big.NewFloat(math.SmallestNonzeroFloat64), 4, '.', '|'), "0|0000"},
		{"-123,456,789", CustomBigCommaf(big.NewFloat(-123456789), -1, ':', ','), "-123:456:789"},
		{"-10,100,000", CustomBigCommaf(big.NewFloat(-10100000), -1, '|', ','), "-10|100|000"},
		{"-10,010,000", CustomBigCommaf(big.NewFloat(-10010000), -1, '*', ','), "-10*010*000"},
		{"-10,001,000", CustomBigCommaf(big.NewFloat(-10001000), -1, '?', ','), "-10?001?000"},
		{"-10,000,000", CustomBigCommaf(big.NewFloat(-10000000), -1, '!', ','), "-10!000!000"},
		{"-100,000", CustomBigCommaf(big.NewFloat(-100000), -1, '.', ','), "-100.000"},
		{"-10,000", CustomBigCommaf(big.NewFloat(-10000), -1, '.', ','), "-10.000"},
		{"-1,000", CustomBigCommaf(big.NewFloat(-1000), -1, '.', ','), "-1.000"},
		{"-100.11", CustomBigCommaf(big.NewFloat(-100.11), -1, '.', ','), "-100,11"},
		{"-10", CustomBigCommaf(big.NewFloat(-10), -1, '.', ','), "-10"},
		{"834,142.3211 fixed to 2", CustomBigCommaf(big.NewFloat(834142.3211), 2, '.', ' '), "834.142 32"},
		{"834,142.3211 fixed to 0", CustomBigCommaf(big.NewFloat(834142.3211), 0, '.', ' '), "834.142"},
		{"834,142.3250", CustomBigCommaf(big.NewFloat(834142.3250), -1, '.', ' '), "834.142 325"},
		{"834,142.3250 fixed to 2 rounds down", CustomBigCommaf(big.NewFloat(834142.3250), 2, '.', ' '), "834.142 32"},
		{"834,142.3251", CustomBigCommaf(big.NewFloat(834142.3251), -1, '.', ' '), "834.142 3251"},
		{"834,142.3251 fixed to 2 rounds up", CustomBigCommaf(big.NewFloat(834142.3251), 2, '.', ' '), "834.142 33"},
	}.validate(t)
}
