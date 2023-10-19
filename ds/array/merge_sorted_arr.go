package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 5, 8, 10, 12}

	rs := mergeArr(arr1, arr2)
	fmt.Println(rs)
}

// complexity O(a + b)
func mergeArr(a, b []int) (r []int) {
	if len(a) == 0 {
		r = b
		return
	}

	if len(b) == 0 {
		r = a
		return
	}

	ai := 0
	bi := 0
	maxA := len(a) - 1
	maxB := len(b) - 1
	aX := a[ai]
	bX := b[bi]

	for ai <= maxA || bi <= maxB {
		if ai > maxA {
			r = append(r, bX)
			bi++
			if bi <= maxB {
				bX = b[bi]
			}
			continue
		}
		if bi > maxB {
			r = append(r, aX)
			ai++
			if ai <= maxA {
				aX = a[ai]
			}
			continue
		}
		if aX < bX {
			r = append(r, aX)
			ai++
			if ai <= maxA {
				aX = a[ai]
			}
		} else {
			r = append(r, bX)
			bi++
			if bi <= maxB {
				bX = b[bi]
			}
		}
	}

	return
}
