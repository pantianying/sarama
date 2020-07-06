package sarama

type TopicRule interface {
	GetClientTopicRule(clientTopic string) ClientTopicRule
	// 判断该topic是否有测试配置
	CheckClientTopicIsRule(clientTopic string) bool
	// 通过测试topic找到真实的clientTopic
	// 如果不是测试的则返回 ""
	GetClientTopicByResponseTopic(brokerTopic string) string
}

type ClientTopicRule interface {
	// 确认是不是需要替换topic
	// 传入msg的key和value，返回是否匹配以及根据哪个配置ConfigKey匹配到
	CheckIsReplaceTopic(key, value []byte) (bool, string)
	// 通过配置值返回brokerTopic
	BrokerTopic(ConfigKey string) string
}
