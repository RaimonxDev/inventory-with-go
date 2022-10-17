package repository

type Repository interface {
	//SaveUser(user model.User)
	GetUserByEmail(id int)
}
