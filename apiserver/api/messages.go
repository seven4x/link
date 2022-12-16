package api

const (
	GlobalActionMustLogin    = "global.action-must-login"
	GlobalErrorAboutDatabase = "global.error-database"
	//参数错误
	GlobalParamWrong = "global.param-wrong"
)

const (
	LinkNotAllowDomain    = "link.not-allow-domain"
	LinkRepeatInSameTopic = "link.repeat-in-same-topic"
)
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
const (
	REGISTER_CODE_ERROR  = "register-code-error "
	REGISTER_NAME_REPEAT = "register-name-repeat "
)
