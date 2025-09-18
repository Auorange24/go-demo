package main

type fd = int

func isReady(x int) bool {
	if x % 2 == 0 {
		return true
	} else {
		return false
	}
}

func handle(x int) {
	// handle id
}

func main() {
	fds := []fd{1, 2, 3, 4}
	for {
		for i := 0 ; i < len(fds) ; i ++ {
			if isReady(fds[i]) {
				handle(fds[i])
			}
		}
		
	}
}