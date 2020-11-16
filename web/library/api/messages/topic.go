package messages

const (
	//不允许创建根级节点
	TopicRootNotAllowed = "topic.msg.root-not-allowed"
	//文本内容审核未通过
	TopicContentNotAllowed    = "topic.msg.content-not-allowed"
	TopicBackendDatabaseError = "topic.msg.database-error"
	TopicRefTopicNoExist      = "topic.msg.ref-topic-not-exist"
	TopicRepeatInSamePosition = "topic.msg.topic-repeat-in-same-position"
	TopicNotFound             = "topic.msg.topic-not-found"
)
