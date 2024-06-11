package database

type DbClient struct {
	Client string
}

func NewDb() DbClient {
	return DbClient{
		Client: "",
	}
}
