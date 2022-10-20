package entity

// User is the representation of our table in the database.
// Important note.
// We are NOT going to return the password in the response, so we create a model.
// The model will allow us to send only the necessary fields to the client.
type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
