package main

///this is a pretext to prac13 where we will use maps

import "fmt"

func main() {

	//this is a map without any values. go by default gives maps just as a reference
	//	grades1:= map[string]float32

	//this is a map with values. Use keyword make to initialise maps
	grades := make(map[string]float32)

	grades["babua"] = 100
	grades["kalua"] = 111
	grades["chutwa"] = 122
	fmt.Println(grades)

	//print the map value of kalua
	KaluaGrade := grades["kalua"]
	fmt.Println(KaluaGrade)

	//deleting a key value pair in map
	delete(grades, "kalua")
	fmt.Println(grades)

	//iterate through map, k is the index or the key , whereas people is the value
	for k, people := range grades {
		fmt.Println(k, ":", people)
	}
}
