package account

const Collection = "account"

type Account struct {
	UUID          string `bson:"_id"`
	WeiXinOpenID  string `bson:"wei_xin_open_id"`
	DeviceID      string `bson:"device_id"`
	Phone         string `bson:"phone"`
	ShorID        int32  `bson:"shor_id"`
	OS            string
	ClientVersion string
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
