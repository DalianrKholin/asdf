package dataBaseModel

type Creds struct {
	Password string `bson:"password" json:"password"`
	Mail     string `bson:"mail" json:"email"`
}
