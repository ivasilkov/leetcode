package main

import "fmt"

func main() {
	//fmt.Println('a' - 'A')
	in1 := "leEeetcode"
	fmt.Println(makeGood(in1))
	in2 := "abBAcC"
	fmt.Println(makeGood(in2))
	in3 := "s"
	fmt.Println(makeGood(in3))
}

func makeGood(s string) string {
	out := []byte(s)
	writePtr := 1
	for readPtr := 1; readPtr < len(out); readPtr++ {
		if writePtr > 0 {
			diff := out[readPtr] - out[writePtr-1]
			if diff == 224 || diff == 32 {
				writePtr--
				continue
			}
		}
		out[writePtr] = out[readPtr]
		writePtr++
	}
	return string(out[:writePtr])
}
