package main

import "fmt"

func fibonacci() func() int {
	fiblower := 0
	fibhigher := 1

	return func() int {
		tmp := fiblower
		fiblower = fibhigher
		fibhigher += tmp

		return tmp
	}
}

func fibdynamic() func(n int) int {
	fiblist := []int{0, 1}

	var fibfoo func(n int) int

	fibfoo = func(n int) int {
		if len(fiblist) <= n {
			fiblist = append(fiblist, fibfoo(n - 1) + fibfoo(n - 2))
		}
		return fiblist[n]
	}

	return fibfoo
}

func fibnaive() func(n int) int {
	var fibfoo func(n int) int

	fibfoo = func(n int) int {
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		} else {
			return fibfoo(n - 1) + fibfoo(n - 2)
		}
	}

	return fibfoo
}

func fib_run() {
	g := fibnaive()
	for i := 0; i < 50; i++ {
		fmt.Println(i, ": ", g(i))
	}

}