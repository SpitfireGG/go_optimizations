package main

func main() {

	var acc string
	num := []int{14, 3125432, 5, 2, 354, 5, 455, 454}
	str := []string{"even", "odd"}
	for _, val := range num {

		acc = str[val%2]
		if acc == "even" {
			println(val)
		}
	}
}

