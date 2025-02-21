Interfaces.go
This Go program demonstrates the use of interfaces by defining an animal interface with a speak() method. It then implements this interface in two struct types: cat and dog.
Animal Sounds Interface
This Go program demonstrates the use of interfaces by defining an Animal interface with a Sound() method. It implements this interface in three struct types: Dog, Cat, and Cow. Each struct has a Sound method that returns a string representing the sound the animal makes. The main function creates instances of Dog, Cat, and Cow, stores them in a slice of type Animal, and iterates over the slice to print the sound each animal makes.
Arrays

Declared with a fixed size and elements of the same type.
Example: var arr [5]int
Initialized: arr = [5]int{1, 2, 3, 4, 5}
Modified: arr[2] = 30
Length: len(arr)
Slices

Dynamic, flexible view into an array's elements.
Example: a := []int{}
Initialized: a = []int{1, 2, 3, 4, 5}
Appended: a = append(a, 6, 7, 8)
Modified: a[3] = 50
Created with make: a = make([]int, 5)
Length and Capacity: len(a), cap(a)
Maps

Unordered collection of key-value pairs.
Example: b := map[string]string{}
Initialized: b = make(map[string]string)
Adding elements: b["car"] = "bmw"
Modified: b["car"] = "THAR"
Deleted: delete(b, "train")
