package helper

const (
	EnvFile                   = "env"
	MinLegalByteAmount        = 50
	MaxLegalByteAmount        = 8000
	TimestampLayout    string = "2006-01-02 03:04:05"

	//AggregatorSocketUrl = "ws://localhost:9990/ws/aggregator"
	AggregatorSocketUrl = "ws://aggregator:9990/ws/aggregator" // docker network
)
