package main

import "fmt"

func main() {
	// * map literal
	m := map[string]int{
		"Amir":  27,
		"Faeze": 24,
	}
	fmt.Println(m)
	fmt.Printf("%#v\n", m) // * using pound
	fmt.Println("The age of Amir was", m["Amir"])

	// * create map using make
	an := make(map[string]int)
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	an["Lucas"] = 28
	an["Steph"] = 37
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	for k := range m {
		fmt.Println(k)
	}

	// * delete a key
	delete(an, "Steph")
	delete(an, "invalid") // * won't panic
	fmt.Println(an)
	fmt.Println(len(an))
	fmt.Println(an["invalid"]) // * won't panic -> return the zero-value

	if v, ok := an["invalid"]; ok {
		fmt.Println("Found the key", v)
	} else {
		fmt.Println("The key doesn't exists")
	}

	fmt.Println("Lucas is", an["Lucas"])
	an["Lucas"]++
	fmt.Println("Lucas now is", an["Lucas"])

	ax := make(map[string][]string)
	ax["Amir"] = []string{"Hossein", "Zare"}
	fmt.Println("ax", ax)
}
