package main

import (
	"L7/utils/reader"
	"L7/utils/utilities"
	"L7/utils/writer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	code(os.Args[1], os.Args[2])
}

func code(inName, outName string) {
	r := reader.Reader_createReader(inName)
	w := writer.Writer_createWriter(outName)

	hammingCodes := utilities.CycleHamming()
	fmt.Println(r.ReadWholeFileGetSizeAndResetReader())
	for r.Reader_PeekBit() {
		four_bits := 0
		for i := 0; i < 4; i++ {
			four_bits = (four_bits << 1) + int(r.Reader_ReadBit())

		}
		w.Writer_addBytes([]byte{hammingCodes[four_bits]})
	}
	w.Writer_Flush()
	w.CloseFile()
}
