package main

import "fmt"

func main() {
	r := decodeString("3[a2[c]]")
	fmt.Println(r)
}

type item struct {
	n     int
	bytes []byte
}

func decodeString(str string) string {
	num := 0
	st := []item{{1, []byte{}}}

	for i := range str {
		switch {
		case str[i] == '0':
			num *= 10
		case str[i] > '0' && str[i] <= '9':
			num = num*10 + int(str[i]-'0')
		case str[i] == '[':
			st = append(st, item{num, []byte{}})
			num = 0
		case str[i] == ']':
			tmp := st[len(st)-1]
			st = st[:len(st)-1]
			for j := 0; j < tmp.n; j++ {
				st[len(st)-1].bytes = append(st[len(st)-1].bytes, tmp.bytes...)
			}
		default:
			st[len(st)-1].bytes = append(st[len(st)-1].bytes, str[i])
		}
	}

	return string(st[0].bytes)
}

//
//n := nSt[len(nSt)-1]
//l1 := len(bSt[len(bSt)-2])
//l2 := len(bSt[len(bSt)-1])
//tmp := make([]byte, l1+n*l2)
//copy(tmp, bSt[len(bSt)-2])
//for j := 0; j < n; j++ {
//	copy(tmp[l1+j*l2:], bSt[len(bSt)-1])
//}
//bSt[len(bSt)-2] = tmp
//bSt = bSt[:len(bSt)-1]
//nSt = nSt[:len(nSt)-1]
