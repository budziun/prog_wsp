package main

import (
	"fmt"
	"math"
)

func policz(dane []float64) float64 {
	var suma float64
	for _, wartosc := range dane {
		suma += wartosc
	}
	return suma
}

func suma_watki(dane []float64, wynik chan float64) {
	var suma float64
	for _, wartosc := range dane {
		suma += wartosc
	}
	wynik <- suma
}

func main() {
	fmt.Println("Lab01")
	fmt.Printf("\nZAD1\n")

	// Proszę stworzyć długą tablicę z wartościami zmiennoprzecinkowymi, n- elementową (na przykład
	// rozmiar = 1024) i wypełnić je wartościami sin(1* 2pi/n), sin(2* 2pi/n)… .
	const tablica_rozmiar = 1024

	tablica := make([]float64, tablica_rozmiar)

	for i := 0; i < tablica_rozmiar; i++ {
		k := float64(i + 1)
		tablica[i] = math.Sin(k * 2.0 * math.Pi / float64(tablica_rozmiar))
	}

	// wyswietlanie 10 elementow zeby sprawdzic
	for i := 0; i < 10; i++ {
		fmt.Println(tablica[i])
	}
	print("\n od tylu teraz\n")
	for i := 1023; i > 1013; i-- {
		fmt.Println(tablica[i])
	}

	fmt.Printf("\nZAD2\n")
	//Następnie proszę policzyć sumę wartości w tej tablicy za pomocą 1 wątku i ją wyświetlić. Powinna być
	//wartość bliska zeru.
	//funkcja policz
	fmt.Println(policz(tablica))

	fmt.Printf("\nZAD3\n")
	//Proszę wyliczyć sumę i pokazać sumę wartości tej samej tablicy przy użyciu metody dziel i zwyciężaj
	//oraz wielowątkowości, bez rekurencji. Proszę dokonać podziału przy pomocy parametru zmiennego w
	//programie oznaczającego liczbę wątków, każdy wątek niech zajmuje się różnym wycinkiem tablicy o
	//podobnej wielkości. Oczywiście, liczba wątków niech będzie znacznie mniejsza niż długość tablicy. Dla
	//uproszczenia można przyjąć, że liczba wątków i długość tablicy to potęgi liczby 2.
	const watki = 128
	dzielenie := tablica_rozmiar / watki
	wynikiChan := make(chan float64, watki)
	for i := 0; i < watki; i++ {
		start := i * dzielenie
		koniec := start + dzielenie
		kawalek := tablica[start:koniec]

		go suma_watki(kawalek, wynikiChan)
	}
	var wynik_ostateczny float64
	for i := 0; i < watki; i++ {
		czastkowa := <-wynikiChan
		fmt.Printf("Wątek %d: %f\n", i+1, czastkowa)
		wynik_ostateczny += czastkowa
	}

	fmt.Println("Suma wielowatkowa:", wynik_ostateczny)
}
