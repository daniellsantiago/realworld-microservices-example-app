package entity

type UserRepository interface {
	Save(user User) error
}
