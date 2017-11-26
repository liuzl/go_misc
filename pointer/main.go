package main

type User struct {
	Name string
}

func main() {
	u := User{Name: "Leto"}
	println(u.Name)
	Modify(u)
	println(u.Name)
}

func Modify(u User) {
	u.Name = "Duncan"
}
