package main

import (
    "fmt"
)

func hitungMenang(g, k int, jm *int) {
    if g > k {
        *jm++
    }
}

func hitungDraw(g, k int, jd *int) {
    if g == k {
        *jd++
    }
}

func hitungKalah(g, k int, jk *int) {
    if g < k {
        *jk++
    }
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
    *jg += g
    *jk += k
    *jsg = *jg - *jk
}

func hitungJumPoint(jm, jd int, jp *int) {
    *jp = (jm * 3) + jd
}

func main() {
    var N int
    fmt.Print("jlh pertandingan: ")
    fmt.Scan(&N)

    jm, jd, jk, jg, kgg, jsg, jp := 0, 0, 0, 0, 0, 0, 0

    for i := 0; i < N; i++ {
        var g, k int
        fmt.Print("skor:")
        fmt.Scan(&g, &k)

        hitungMenang(g, k, &jm)
        hitungDraw(g, k, &jd)
        hitungKalah(g, k, &jk)
        hitungJumGolKegolanSelisih(g, k, &jg, &kgg, &jsg)
    }

    hitungJumPoint(jm, jd, &jp)

    fmt.Printf("%d\n", N,)
    fmt.Printf("%d\n", jm)
    fmt.Printf("%d\n", jd)
    fmt.Printf("%d\n", jk)
    fmt.Printf("%d\n", jg)
    fmt.Printf("%d\n", kgg)
    fmt.Printf("%d\n", jsg)
    fmt.Printf("%d\n", jp)
}
