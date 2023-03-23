package entities

type AWSCredentials struct {
	ServerPort      string `json:"AUTH_SERVER_PORT"`
	RabbitMQConn    string `json:"RABBITMQ_CONN"`
	DBConnString    string `json:"DB_CONN_STRING"`
	TokenExpiryTime string `json:"TOKEN_EXPIRY_TIME"`
}
