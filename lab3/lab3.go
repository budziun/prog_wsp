package main

import (
	"fmt"
	//"sync"
)

//zad dluga tabela wypelniamy zeby byla suma 0 wart -1/3 ... +1/3
// powinno wyjsc kolo 0 ale nie zero
// dzielimy potem na 4 czesci tablice
// za pomoca kanalu odpalamy 2 kanaly
// w jednym kanale wrzucam info o nr czesci
// a w drugim kanale zbieramy wartosci/ sumy tych czesci -> wyrzucam na ekran sumuje i patrze co sie dzieje
// dziel i zwyciezaj
// stworzyc dwa kanaly na poczatek
// robotnicy zrobic ktory sprawdza czy kanal sie nie zamknal
// jak otrzymal nr czesci to sumuje wartosci z kanalu jak dowidzenia to koniec
// robotnik odczytuje sprawdza czy jest jeszcze robota sumuje i wynik zwraca
// watek glowny -> wrzuca w nr czesci i odczytuje wynik (nie efektywna metoda) ALBO damy wiecej robotnikow np 10 do roznych robot
// i pozniej wszystkie informacje otrzymuje (liczba robot nie moze byc wieksza niz liczba roboli)
// trzeba dodac bufor do zrzucania liczby zadan wtedy 20 zadan = bufor wiekszy, i pozniej czekam na wyniki otrzymamy 10 zadan
// a potem reszte i nic sie nie blokuje
// bufor = kolejka - bedzie se trzymal to i czekal az sie robole zwolnia
// kto ma czekac? lepiej zeby nie czekal ten kotry dluzej robi

const rozmiar = 10000

var tablica []float64

func wykonawca(wyk_id int, kanal_we chan int, kanal_wy chan float64) {
	dlugoscCzesci := rozmiar / 4
	for {
		nrCzesci, czy_otwarty := <-kanal_we
		if !czy_otwarty {
			break
		}
		poczatek := nrCzesci * dlugoscCzesci
		koniec := poczatek + dlugoscCzesci

		var suma float64 = 0.0
		for i := poczatek; i < koniec; i++ {
			suma += tablica[i]
		}
		fmt.Printf("[Watek %d] Część %d: suma = %f\n", wyk_id, nrCzesci, suma)
		kanal_wy <- suma
	}
}

func main() {
	fmt.Println("LAB3")
	//zad1
	tablica = make([]float64, rozmiar)

	for i := 0; i < rozmiar; i++ {
		if i < rozmiar/2 {
			tablica[i] = 1.0 / 3.0
		} else {
			tablica[i] = -1.0 / 3.0
		}
	}
	var suma float64
	for _, wartosc := range tablica {
		suma += wartosc
	}
	fmt.Printf("Suma tablicy: %.15f\n", suma)

	// zad2
	/*
		kanal_we := make(chan int, 4)
		go wykonawca(1, kanal_we)
		go wykonawca(2, kanal_we)

		for i := 0; i < 4; i++ {
			kanal_we <- i
		}
		close(kanal_we)
		time.Sleep(1 * time.Second)
		fmt.Println("Koniec") */
	//zad 3
	kanal_we := make(chan int, 4)
	kanal_wy := make(chan float64, 4)

	go wykonawca(1, kanal_we, kanal_wy)
	go wykonawca(2, kanal_we, kanal_wy)

	for i := 0; i < 4; i++ {
		kanal_we <- i
	}
	close(kanal_we)
	/* zad3
	for i := 0; i < 4; i++ {
		suma := <-kanal_wy
		fmt.Printf("suma otrzymana: %.6f\n", suma)
	}
	*/
	//zad 4
	var final_suma float64 = 0.0
	for i := 0; i < 4; i++ {
		suma := <-kanal_wy
		fmt.Printf("suma otrzymana: %.6f\n", suma)

		final_suma += suma
	}

	fmt.Printf("Koncowa suma od roboli: %.15f\n", final_suma)
	fmt.Println("KONIEC")

}
