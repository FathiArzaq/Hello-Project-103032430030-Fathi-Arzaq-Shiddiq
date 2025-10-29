// TESSS

package main

import (
	"fmt"
)

func hitungLuasKelilingLingkaran(r float64) (float64, float64) {
	const pi = 3.14
	l := pi * r * r
	k := 2 * pi * r
	return l, k
}

func hitungLuasKelilingPersegi(s float64) (float64, float64) {
	l := s * s
	k := 4 * s
	return l, k
}

func hitungTotal(lL, lP, kL, kP float64) (float64, float64) {
	totLuas := lL + lP
	totKel := kL + kP
	return totLuas, totKel
}

func main() {
	var r, s float64
	fmt.Println("radius lingkaran, sisi persegi:")

	for {
		fmt.Scan(&r, &s)
		if r == 0 && s == 0 {
			break
		}

		lLingkaran, kLingkaran := hitungLuasKelilingLingkaran(r)
		lPersegi, kPersegi := hitungLuasKelilingPersegi(s)
		totalLuas, totalKel := hitungTotal(lLingkaran, lPersegi, kLingkaran, kPersegi)

		fmt.Printf("R: %.2f, S: %.2f, LL: %.2f, LP: %.2f, KL: %.2f, KP: %.2f, TL: %.2f, TP: %.2f\n",
			r, s, lLingkaran, lPersegi, kLingkaran, kPersegi, totalLuas, totalKel)
	}
}
