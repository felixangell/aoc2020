package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func bounds(array []uint64) (uint64, uint64) {
	min, max := array[0], array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func sumExistsIn(verify []uint64, val uint64) bool {
	for i := 0; i < len(verify); i++ {
		for j := 0; j < len(verify); j++ {
			s := verify[i] + verify[j]
			if s == val {
				return true
			}
		}
	}
	return false
}

func validate(p int, data []uint64) (uint64, error) {
	nop := data[p:]
	pre := data[:p]
	for idx, v := range nop {
		if !sumExistsIn(pre, v) {
			return v, errors.New("no sum found")
		}

		ai := idx + p
		pre = data[ai-p : ai]
	}
	return 0, nil
}

func main() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var data []uint64
	for scanner.Scan() {
		val, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		data = append(data, val)
	}

	bval, err := validate(5, data)
	if err == nil {
		panic("uh oh")
	}

	fmt.Println("looking for ", bval)

	// find subarray producing sum of bval
	c, s := data[0], uint64(0)

	for i := uint64(1); i <= uint64(len(data)); i++ {
		for c > bval && s < i-1 {
			c = c - data[s]
			s++
		}
		fmt.Println(c)
		if c == bval {
			sa := data[s : i-1]
			fmt.Println("range of subarray found", sa, " = ", bval)

			min, max := bounds(sa)
			fmt.Println("super secret value", min+max)

			break
		}
		if i < uint64(len(data)) {
			c = c + data[i]
		}
	}
}
