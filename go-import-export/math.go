package main

func Add(a int, b int) int { // Exported because Add starts with capital A
	return a + b
}

func subtract(a int, b int) int { // Not exported
	return a - b
}
