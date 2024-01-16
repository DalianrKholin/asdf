package model


type Creds struct{
    Password string `bson:"password"`
    Email string    `bson:"mail"`
}