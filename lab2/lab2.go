package main

import (
	"fmt"
	"sync"
	//"time"
)

var zm_glob int
var mut sync.Mutex
var wg sync.WaitGroup
var mutr sync.RWMutex

const rozmiar = 10000

var tablica []float64

func statystyka() {
	if len(tablica) == 0 {
		fmt.Println("PUSTA")
		return
	}
	var suma float64
	maks := tablica[len(tablica)-1]

	for i := len(tablica) - 1; i >= 0; i-- {
		suma += tablica[i]
		if tablica[i] > maks {
			maks = tablica[i]
		}
	}
	fmt.Printf("[Statystyka] Suma: %.2f, Maksymalna wartość: %.2f\n", suma, maks)
}

func czytacz(id int) {
	defer wg.Done()

	fmt.Printf("Czytacz %d: czekam na dostęp do odczytu...\n", id)
	mutr.RLock()
	fmt.Printf("Czytacz %d: uzyskałem RLock, uruchamiam statystyka()\n", id)

	statystyka()

	mutr.RUnlock()
	fmt.Printf("Czytacz %d: zakończyłem pracę (RUnlock)\n", id)
}

func pisarz_mnoznik(id int) {
	defer wg.Done()
	fmt.Printf("Pisarz mznoznik %d czeka na dostep\n", id)
	mutr.Lock()
	fmt.Printf("Pisarz mnzonik %d dostalem mnzoe razy dwwwwa\n", id)

	for i := 0; i < len(tablica); i++ {
		tablica[i] *= 2.0
	}
	fmt.Printf("Koniec dla %d\n", id)
	mutr.Unlock()
}

func main() {
	fmt.Println("LAB2")
	tablica = make([]float64, rozmiar)

	for i := 0; i < rozmiar; i++ {
		if i < rozmiar/2 {
			tablica[i] = 1.0
		} else {
			tablica[i] = -1.0
		}
	}
	fmt.Printf("Rozmiar: %d\n", len(tablica))
	fmt.Printf("Pierwszy element: %.1f, środkowy: %.1f, ostatni: %.1f\n",
		tablica[0], tablica[rozmiar/2], tablica[rozmiar-1])

	//do pkt 2
	//statystyka()

	//zad 3
	/*
		wg.Add(3)
		go czytacz(1)
		go czytacz(2)
		go czytacz(3)
		wg.Wait()
		fmt.Println("Koniec pracy watki zakonczyly robote")
	*/
	//zad 4
	/*
		fmt.Println("Start")
		statystyka()
		wg.Add(3)
		go pisarz_mnoznik(1)
		go pisarz_mnoznik(2)
		go pisarz_mnoznik(3)

		wg.Wait()
		fmt.Println("KONIEC MNOZENIA")
		statystyka()
	*/
	//zad5
	wg.Add(5)

	go czytacz(1)
	go pisarz_mnoznik(101)
	go czytacz(2)
	go pisarz_mnoznik(102)
	go czytacz(3)

	wg.Wait()
	fmt.Println("\nWszystkie watki zakonczyly dzialanie. Stan koncowy:")
	statystyka()
}
