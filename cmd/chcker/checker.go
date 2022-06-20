package main

import (
	"L7/utils/reader"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	fmt.Println(check(os.Args[1], os.Args[2]))
}

func check(inName1, inName2 string) int {
	r1 := reader.Reader_createReader(inName1)
	r2 := reader.Reader_createReader(inName2)

	size1, size2 := r1.ReadWholeFileGetSizeAndResetReader(), r2.ReadWholeFileGetSizeAndResetReader()
	var bits1 [4]uint8
	var bits2 [4]uint8
	counter := 0
	i := int64(0)
	for i < size1 && i < size2 {
		for j := 0; j < 4; j++ {
			bits1[j] = r1.Reader_ReadBit()
			bits2[j] = r2.Reader_ReadBit()
		}
		for j := 0; j < 4; j++ {
			if bits1[j] != bits2[j] {
				counter++
				break
			}
		}
		i += 4
	}
	if i < size1 || i < size2 {
		var howManyLeft int64
		if size1 < size2 {
			howManyLeft = size2 - size1
		} else {
			howManyLeft = size1 - size2
		}
		howManyLeft /= 4
		fmt.Println("MISMATCH", "Pozostało ", howManyLeft, " niesprawdzonych bloków (różna długość plików).")
	}
	return counter
}
