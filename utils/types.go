package utils

type User struct {
	ID       int
	Role     string
	Username string
	Password string
	FName    string
	SName    string
	ClassID  int
	Points   int
}

type Class struct {
	ID        int
	ClassName string
}
