package main

import (
	"fmt"
)

func deferSample1() {
	fmt.Println("POINT:1")
	defer fmt.Println("DEFER:1")
	fmt.Println("POINT:2")
	defer fmt.Println("DEFER:2")
	fmt.Println("POINT:3")
	defer fmt.Println("DEFER:3")
	fmt.Println("END")
}

// 途中でreturnされた場合は処理がそこで終了
func deferSample2(i int) {
	fmt.Println("POINT:1")
	defer fmt.Println("DEFER:1")
	if i == 1 {
		return
	}
	fmt.Println("POINT:2")
	defer fmt.Println("DEFER:2")
	if i == 2 {
		return
	}
	fmt.Println("POINT:3")
	defer fmt.Println("DEFER:3")
	if i == 3 {
		return
	}
	fmt.Println("END")
}

// defer は実行処理を通らない限り実行されない
func deferSample3(i int) {
	fmt.Println("POINT:1")
	if i > 1 {
		defer fmt.Println("DEFER:1")
	}
	fmt.Println("POINT:2")
	if i > 2 {
		defer fmt.Println("DEFER:2")
	}
	fmt.Println("POINT:3")
	if i > 3 {
		defer fmt.Println("DEFER:3")
	}
	fmt.Println("END")
}

// return値の名前をあらかじめ設定しておく
func sampleReturn() (err error) {
	defer func() {
		fmt.Println(err)
	}()
	err = fmt.Errorf("ERROR")
	return err
}

func main() {
	fmt.Println("-0--------------------------")
	deferSample1()
	fmt.Println("-1--------------------------")
	deferSample2(1)
	fmt.Println("-2--------------------------")
	deferSample2(2)
	fmt.Println("-3--------------------------")
	deferSample2(3)
	fmt.Println("-4--------------------------")
	deferSample2(4)
	fmt.Println("-5--------------------------")
	deferSample3(3)
	fmt.Println("-6--------------------------")
	sampleReturn()
}
