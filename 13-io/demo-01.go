package main

import "io"

type AlphaReader string // of type string

//io.Reader interface implementation
func (alphaReader AlphaReader) Read(buffer []byte) (n int, err error) { //alphaReader is a receiver
	count := 0
	lenBuff :=len(buffer)
	println(lenBuff)
	for idx := 0; idx < len(alphaReader); idx++ {
		if (alphaReader[idx] >= 'a' && alphaReader[idx] <= 'z') || (alphaReader[idx] >= 'A' && alphaReader[idx] <= 'Z') {
			buffer[count] = alphaReader[idx]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	str := "Consequat id eiusmod 23479 laborum non esse consequat cupidatat elit fugiat officia fugiat. Ad culpa aliqua 340 *&^%$&() cupidatat nulla qui minim pariatur cupidatat mollit aliqua velit eu officia. Amet ex voluptate velit dolor nulla nostrud dolor ullamco anim enim. Consequat sit et pariatur esse ipsum amet duis ullamco. Labore ad culpa deserunt est ea. Fugiat duis ad dolore dolor sunt eu qui ipsum aliqua laborum."
	reader := AlphaReader(str)
	buffer := make([]byte, len(str)) //pointer to a buffer
	println(len(buffer))
	noOfBytesRead, err := reader.Read(buffer) //calling reader of AlphaReader
	println(noOfBytesRead)
	if err == io.EOF {
		println(string(buffer[:noOfBytesRead]))
		return
	}
	if err != nil {
		panic(err)
	}
}