// implementing a swap fucn using sawp

package main

func swap(x, y *int) {

	temp := *x
	*x = *y
	*y = temp

	// // same as with no pointer ????????
	// return *x, *y
	// dont need return cause i already sawped the poinnter
	// so x pointer = y and visa vrca
	// why use pointer when i can just use the var to do it
}
