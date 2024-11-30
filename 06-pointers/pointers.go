package main

func birthday(age *int) {
	// *age = *age + 1
	// *age++
	(*age)++
}

func main() {
	num := 27
	pointer := &num
	println("num:", num)
	println("pinter:", pointer)
	println("value at pointer:", *pointer)

	var age int = 26
	println("age:", age)

	birthday(&age)
	println("new age:", age)
}
