package entity

//go:generate go run go.uber.org/mock/mockgen -destination mock_entity/repository.go . UserRepository
type UserRepository interface {
	Save(user User) error
}
