package dto

type AppConfig struct {
	DatabaseDSN string
	ServerPort  string
	//JWTConfig   JWTConfig
}

//type JWTConfig struct {
//	SecretKey        string
//	AccessExpiration time.Duration
//}
