package user

type User struct{
	Name string
	Age int
}

func CreateUser(name string, age int) User{
    return User{
		Name: name,
		Age: age,
	}
}

func GetName(u User)string{
	return u.Name
}