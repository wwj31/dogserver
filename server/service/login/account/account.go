package account

const Collection = "account"

type Account struct {
	UUID          string `bson:"_id"`
	PlatformID    string
	SID           string
	OS            string
	ClientVersion string
	Language      string
	Country       string
	LastLoginRID  string
	Roles         map[string]Role
}

type Role struct {
	RID      string
	CreateAt string
}

func New() *Account {
	return &Account{}
}
