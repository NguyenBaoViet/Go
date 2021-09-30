package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_janes`:       {`Shaken, not stirred`, `Martinis`, `Women`},
		`monneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:            {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["viet_nguyen"] = []string{`123`, `sdawd`, `adsad`}

	for k, v := range x {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
