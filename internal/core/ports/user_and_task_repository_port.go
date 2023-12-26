package ports

type UserAndTaskRepository interface {
	UserRepository
	TaskRepository
}
