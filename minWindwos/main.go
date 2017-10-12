package main

type dict struct{
	name byte
	index int
}


func minWindows(s string, t string) string {
	m := make(map[byte]int)
	a := make([]dict, len(t))
	for _, tmp := range a{
		m = 10000
	}
	for idx, x := range s {
		if isIn(t, x) {
			m[x] = idx
			if  
		}
	}
}

func isIn(s string, b byte) bool {
	for _, x := range s {
		if x == b {
			return true
		}
	}
	return false
}

func findMin(mydict []dict, ){
	for _,x := range dict{
		
	}
}
