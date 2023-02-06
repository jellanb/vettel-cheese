package collections

type Users struct {
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Lastname string `bson:"lastname"`
	Password string `bson:"password"`
	Rol      int    `bson:"rol"`
}
