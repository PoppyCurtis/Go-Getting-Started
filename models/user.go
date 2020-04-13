package models

//Supporting package. For sub packages just need directory name

type User struct {
	ID        int
	FirstName string
	LastName  string
}

//variable block
var (
	users  []*User
	nextID = 1
)

//value of 1 implies that data type is going to be an int -> don't need to declare 'int'

func getUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
