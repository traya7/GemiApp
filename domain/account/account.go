package account

type Account struct {
	ID       string `bson:"_id"`
	Username string
	Password string
	Balance  float64
	Role     string
}
