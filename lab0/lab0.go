package main

import "fmt" // biblioteka przetwarzania tekstu i konsola to
// import "time" potrzebne do time.Sleep

func silnia(n int) int {
	if n <= 1 {
		return 1
	}
	return n * silnia(n-1)
}

func main() {
	var x1 int32
	x1 = 123
	fmt.Printf("%d", x1)
	// wszystkie zmienne musza byc uzyte !!!
	x2 := x1 // przypisuejmy wartosc i tworzymy nowa zmienna
	// typ wynika z tego co przekazalismy
	fmt.Printf("%d", x2)
	// implicidy - konwersja jawna
	// emplicidy - niejawna konwersja
	var x int32
	x -= 3              // wszystkie zmienne maja wartosc zadeklarowana np. int32 ma 0 wartosc
	fmt.Printf("%d", x) // output: -3
	var a int32
	a = 2
	for a < 10 {
		a += 5
	}
	fmt.Println(a)
	fmt.Println("Hello world!") //wyswietl na konsoli
	fmt.Println(silnia(10))
	// w go nie ma refernecji
	// albo zmienna albo wskaznik
	// wskaznik * a tym & odwolanie
	// z tego bedziemy czesto korzytac
	// dziala to jak w c++
	// go polecneie uruchamia kod w nowym watku mega wazne

}
