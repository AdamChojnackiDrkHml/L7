package main

import (
	"L7/utils/reader"
	"L7/utils/utilities"
	"L7/utils/writer"
	"fmt"
	"os"
)

var checkToShift = [8]uint8{0, 7, 6, 4, 1, 5, 2, 3}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	a, b := dekoder(os.Args[1], os.Args[2])
	fmt.Println(a, b)
}

func reverseCycleHamming() [256]uint8 {
	h := utilities.CycleHamming()
	var r [256]uint8
	for i := 0; i < 256; i++ {
		r[i] = 16
	}
	for number := uint8(0); number < 16; number++ {
		r[h[number]] = number
	}
	return r
}

func checkMatrix() [256]uint8 {
	var m [256]uint8
	i := uint8(0)
	for j := 0; j < 256; j++ {
		a := ((i >> 7) & 1)
		b := ((i >> 6) & 1)
		c := ((i >> 5) & 1)
		d := ((i >> 4) & 1)
		e := ((i >> 3) & 1)
		f := ((i >> 2) & 1)
		g := ((i >> 1) & 1)
		check := (((c + e + f + g) % 2) << 2) +
			(((b + d + e + f) % 2) << 1) +
			((a + c + d + e) % 2)

		m[i] = check
		i++
	}
	return m
}

func dekoder(inName, outName string) (int, int) {
	r := reader.Reader_createReader(inName)
	w := writer.Writer_createWriter(outName)
	reverseHamming := reverseCycleHamming()
	checkCodes := checkMatrix()
	counter := 0
	fix_counter := 0

	for {
		nextByte, err := r.Reader_readByte()
		if err != nil {
			break
		}
		parity := (nextByte & 1)
		ones_counter := uint8(0)
		for i := 7; i > 0; i-- {
			ones_counter += ((nextByte >> i) & 1)
		}
		ones_counter %= 2
		check := checkCodes[nextByte]
		to_write := uint8(0)
		if parity == ones_counter { // parzysta liczba błędów
			if check != 0 {
				counter++
			}
			to_write = reverseHamming[nextByte]
		} else { // nieparzysta liczba błędów
			shift := checkToShift[check]
			nextByte = nextByte - (((nextByte >> shift) & 1) << shift) + ((1 - ((nextByte >> shift) & 1)) << shift)
			to_write = reverseHamming[nextByte]
			fix_counter++
		}
		for i := 3; i >= 0; i-- {
			w.Writer_addBits([]uint8{(to_write >> i) & 1})
		}
	}
	w.Writer_Flush()
	w.CloseFile()
	return counter, fix_counter
}
