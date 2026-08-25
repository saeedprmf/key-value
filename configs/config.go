package configs



type Config struct {
	Shard_count int  `json:"shard_count"`
}

var config Config

func GetSharCount() int {
	return config.Shard_count
}