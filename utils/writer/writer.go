package writer

import (
	"os"
)

type Writer struct {
	path       string
	file       *os.File
	byteBuffer []byte
	bitBuffer  []byte
	written    int
}

func (writer *Writer) openOrCreateFile() {
	file, err := os.OpenFile(writer.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	writer.written = 0
	writer.file = file

}

func Writer_createWriter(path string) *Writer {
	writer := &Writer{path: path}

	writer.openOrCreateFile()

	return writer
}

func (w *Writer) write() {
	_, err := w.file.Write(w.byteBuffer)
	// f1, _ := w.file.Stat()
	// fmt.Println(f1.Size(), n, len(w.byteBuffer))
	if err != nil {
		panic(err)
	}

	w.byteBuffer = w.byteBuffer[len(w.byteBuffer):]
	w.written += len(w.byteBuffer)

}

func (writer *Writer) CloseFile() {
	writer.file.Close()
}

func (w *Writer) addByte(byteBits []byte) {

	newByte := getByteFromBits(byteBits)
	w.byteBuffer = append(w.byteBuffer, newByte)

	if len(w.byteBuffer) == 256 {
		w.write()
	}
}

func (w *Writer) Writer_addBits(bits []byte) {

	w.bitBuffer = append(w.bitBuffer, bits...)
	for len(w.bitBuffer) >= 8 {
		w.addByte(w.bitBuffer[:8])
		w.bitBuffer = w.bitBuffer[8:]
	}

}

func (w *Writer) Writer_addBytes(bytes []byte) {

	w.byteBuffer = append(w.byteBuffer, bytes...)

	for len(w.byteBuffer) >= 256 {
		w.write()
	}
}

func getByteFromBits(bits []byte) byte {
	acc := byte(0)

	for _, n := range bits {
		acc *= 2
		acc += n
	}

	return acc
}

func (w *Writer) Writer_Flush() {
	if len(w.bitBuffer) != 0 {
		fill := make([]byte, 8-len(w.bitBuffer))
		w.Writer_addBits(fill)
	}

	if len(w.byteBuffer) != 0 {
		w.write()
	}

	// fmt.Println(w.written)
	// f1, _ := w.file.Stat()
	// fmt.Println(f1.Size())
}
