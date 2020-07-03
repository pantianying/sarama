package sarama

type TopicRule interface {
	CheckIsReplaceTopic(key, value []byte) bool
}
