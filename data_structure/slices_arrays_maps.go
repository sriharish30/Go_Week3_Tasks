package main

import "fmt"

func main() {
	//arrays
	fmt.Println("======ARRAYS======")
	var arr [5]int
	fmt.Println("empty array:", arr)
	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("initialised array:", arr)
	arr[2] = 30
	fmt.Println("modified array:", arr)
	fmt.Println("length of an array:", len(arr))
	//slices
	fmt.Println("======SLICES======")
	a := []int{}
	fmt.Println("empty slices:", a)
	a = []int{1, 2, 3, 4, 5}
	fmt.Println("initialised slices:", a)
	a = append(a, 6, 7, 8)
	fmt.Println("appended slice:", a)
	a[3] = 50
	fmt.Println("modified slices:", a)
	fmt.Println("element at index 5: ", a[5])
	a = make([]int, 5)
	fmt.Println("my slice:", a)
	fmt.Println("length of the slice:", len(a))
	fmt.Println("capacity of the slice:", cap(a))
	//maps
	fmt.Println("======MAPS======")
	b := map[string]string{}
	fmt.Println("empty map:", b)
	b = make(map[string]string)
	b["car"] = "bmw"
	b["bike"] = "bullet"
	fmt.Println("initialised map:", b)
	b["train"] = "metro"
	b["bus "] = "volvo"
	fmt.Println("maps after adding elements:", b)
	b["bike"] = "ROYAL ENFEILD "
	b["car"] = "THAR"
	fmt.Println("modified the maps:", b)
	delete(b, "train")
	fmt.Println("map after deleting train:", b)
	for k, y := range b {
		fmt.Printf("%v:%v", k, y)
	}
}
