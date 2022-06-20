package main

import (
	"L7/utils/reader"
	"L7/utils/writer"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	in, out := "../../dupa", "../../dupa2"
	pD := 0.0005
	if len(os.Args) < 3 {
		fmt.Println(noise(in, out, pD))
	} else {
		p, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(noise(os.Args[1], os.Args[2], p))
	}

}

func noise(inName, outName string, p float64) int {
	r := reader.Reader_createReader(inName)
	w := writer.Writer_createWriter(outName)
	fmt.Println(r.ReadWholeFileGetSizeAndResetReader())
	counter := 0
	rand.Seed(time.Now().UnixNano())
	for r.Reader_PeekBit() {

		b := r.Reader_ReadBit()
		if rand.Float64() < p {
			w.Writer_addBits([]uint8{1 - b})
			counter++
		} else {
			w.Writer_addBits([]uint8{b})
		}

	}
	w.Writer_Flush()
	w.CloseFile()

	return counter
}
