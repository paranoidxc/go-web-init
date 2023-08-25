package redis

// redis key
const (
	Prefix       = "hunter"
	KeyPostTime  = "post:time"
	KeyPostScore = "post:score"
	KeyPostVoted = "post:voted"
)

func getRedisKey(key string) string {
	return Prefix + key
}
