package main

import ( "fmt"
"bytes")

func main() {
	//basicExamples()
	//stringExamples()
	capacityExamples()
}

func capacityExamples() {
	extendExamples()
}

func extendExamples() {
	var iBuffer [10]int
    slice := iBuffer[0:0]
    for i := 0; i < 20; i++ {
    	if cap(slice) == len(slice) {
    		fmt.Println("slice is full!")
    		return
		}
        slice = extendSlice(slice, i)
        fmt.Println(slice)
    }
}

func extendSlice(slice []int, element int) []int {
    n := len(slice)
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	index := bytes.LastIndex(*p, []byte("/"))
	if index >= 0 {
		*p = (*p)[0:index]
	}
}

func (p path) ToUpper() {
    for i, b := range p {
        if 'a' <= b && b <= 'z' {
            p[i] = b + 'A' - 'a'
        }
    }
}

func stringExamples() {
	//truncatePathExample()
	toUpperPathExample()
}

func toUpperPathExample() {
	pathName := path("/usr/bin/tso")
    pathName.ToUpper()
    fmt.Printf("%s\n", pathName)
}

func truncatePathExample() {
	pathName := path("/usr/bin/tso")
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
}

func basicExamples() {
	var buffer [256]byte
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }

	passPointerToSliceExample(slice)
	//changeSliceIndexExample(slice)
	//changeSliceDataExample(buffer, slice)

	fmt.Println("slice", slice)
}

func passPointerToSliceExample(slice []byte) {
	fmt.Println("Before: len(slice) =", len(slice))
    ptrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}

func changeSliceIndexExample(slice []byte) {
	fmt.Println("Before: len(slice) =", len(slice))
    newSlice := subtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("slice", slice)
    fmt.Println("After:  len(newSlice) =", len(newSlice))
    fmt.Println("newSlice", newSlice)
}

func changeSliceDataExample(buffer [256]byte, slice []byte) {
	fmt.Println("before", slice)
    addOneToEachElement(slice)
    fmt.Println("after", slice)
    fmt.Println("buffer", buffer)
}

func ptrSubtractOneFromLength(slicePtr *[]byte) {
    //slice := *slicePtr
    //*slicePtr = slice[0 : len(slice)-1]
    *slicePtr = (*slicePtr)[0:len(*slicePtr)-1]
}

func subtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func addOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}
