package user

type User struct {
	ID        uint32
	Email     string
	Canonical string

	Salt     string
	Password string
}
