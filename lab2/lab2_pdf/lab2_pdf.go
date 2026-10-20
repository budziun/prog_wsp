package main

import (
	"fmt"
	"sync"
	"time"
)

var zm_glob int
var mut sync.Mutex
var wg sync.WaitGroup
var mutr sync.RWMutex

// nie 0 moze nam wyjsc w trakcie po skonczeniu transakcji ma wyjsc 0
// opcja statystyka - jak na samym koncu odpalimy to wychodzi ladne 0

// tutaj dodalismy mutex czyli semator binarny - dzieki niemu wiecej watkow moze pracowac
// Lock() - metoda ta sprawdza, czy ktoś inny już zablokował ten obiekt. Jeśli obiekt nie jest zablokowany, wtedy go blokuje. Jeśli obiekt jest zablokowany, wtedy wątek usypia się czekając dopóki, obiekt ten zostanie odblokowany, następnie go sam blokuje.
// Unlock() - metoda odblokowuję ten obiekt, jeśli jakieś inne wątki zostały uśpione przez blokadę na tym obiekcie, wtedy wybudza jeden z nich.
/* func dodawacz() {
	for i := 0; i < 100; i++ {
		mut.Lock()
		tmp := zm_glob
		tmp++
		time.Sleep(1 * time.Millisecond)
		zm_glob = tmp
		mut.Unlock()
	}
	fmt.Println(zm_glob)
} */
//dodawacz z waitgroup
func dodawacz() {
	fmt.Println("start watku")
	for i := 0; i < 100; i++ {
		mut.Lock()
		tmp := zm_glob + 1
		time.Sleep(1 * time.Millisecond)
		zm_glob = tmp
		mut.Unlock()
	}
	fmt.Println("koniec watku", zm_glob)
	wg.Done()
}

var flaga_stopu bool

func spamer() {
	for !flaga_stopu {
		fmt.Println(("plum-plum, przerwa 0.4s start"))
		time.Sleep(400 * time.Millisecond)
		fmt.Println("plum-plum, koniec przerwy")
	}
}

func czytacz(id int) {
	fmt.Println("id= ", id, ", start watku czytacz()")
	defer wg.Done()
	mutr.RLock() // poczekaj, aż pisarz się skończy, ale nie blokuj innych czytaczy()
	fmt.Println("id= ", id, ", tuz po RLock(), zmienna= ", zm_glob)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("id= ", id, ", tuz przed RUnlock(), zmienna= ", zm_glob)
	mutr.RUnlock() // zwolnienie, zrzeczenie sie praw do czystego odczytu
	fmt.Println("id= ", id, ", koniec watku czytacz()")

}

func pisarz_zwiekszacz(id int) {
	fmt.Println("id= ", id, ", start watku pisarz_zwiekszacz()")
	defer wg.Done()
	mutr.Lock() // poczekaj, aż nikt czyta i pisze
	fmt.Println("id= ", id, ", tuz po Lock(), zmienna= ", zm_glob)
	tmp := zm_glob
	time.Sleep(10 * time.Millisecond)
	zm_glob = tmp + 1
	fmt.Println("id= ", id, ", tuz przed Unlock(), zmienna= ", zm_glob)
	mutr.Unlock() // zwolnienie, zrzeczenie sie praw do czystego odczytu
	fmt.Println("id= ", id, ", koniec watku pisarz_zwiekszacz()")
}

func main() {
	fmt.Println("LAB2")
	//zm_glob = 0
	//go dodawacz()
	//go dodawacz()
	//go dodawacz()
	//time.Sleep(3 * time.Second)
	//fmt.Println("koniec, ", zm_glob)

	//flaga_stopu = false
	//go spamer()
	//time.Sleep(2 * time.Second)
	//flaga_stopu = true
	//fmt.Println("flaga stopu zostala ustawiona, czekamy jeszcze 1s")
	//time.Sleep(1 * time.Second)

	/*
		wg.Add(1)
		go dodawacz()
		wg.Add(1)
		go dodawacz()
		wg.Add(1)
		go dodawacz()
		wg.Wait()
		fmt.Println("koniec main, ", zm_glob)
	*/
	zm_glob = 11
	wg.Add(5) // zamiast dodawać pojedyczo, niezbyt to eleganckie, ale bezpieczne
	go czytacz(101)
	time.Sleep(1 * time.Millisecond)
	go czytacz(102)
	go czytacz(103)
	go pisarz_zwiekszacz(801)
	time.Sleep(1 * time.Millisecond)
	go pisarz_zwiekszacz(802)
	wg.Wait()
	fmt.Println("koniec main, ", zm_glob)

}
