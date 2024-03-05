package transaction

type Transaction struct {
	ID        string `bson:"_id"`
	From      string
	To        string
	Amount    int64
	CreatedAt int64
}
