package main

// arry_slice is slice just looks like a arry

func Removeslice(index int, s_slice []int) []int {

	if index < 0 || index < len(s_slice) {
		return append(s_slice[:index], s_slice[index+1:]...)
	}

	panic(" Invalid index !!!")
}
