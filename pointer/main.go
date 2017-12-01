package main

type User struct {
	Name string
}

func main() {
	u := User{Name: "Leto"}
	println(u.Name)
	Modify(u)
	println(u.Name)
	arr := []int{}
	println(arr)
	M(&arr)
	println(arr)
}

func Modify(u User) {
	u.Name = "Duncan"
}

func M(arr *[]int) {
	*arr = append(*arr, 1)
}
