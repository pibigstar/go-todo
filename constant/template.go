package constant

import "time"

const (
	// 任务接收消息模板ID
	TemplateReceiveTaskId = "YWAd-Djf_hURHmf1HMobGSrd8AD53j_IoBJ35tffkPQ"
	// redis存储AccessToken前缀
	RedisPrefixAccessToken = "todo:accessToken:%s"
	// redis存储formID前缀
	RedisPrefixFormId = "todo:formID:%s"
	// 收集的用户的formID,OpenID
	CollectionFormIdPrefix = "formIds:%s"
	// 收集的用户的formID的过期时间
	UserFormIdExpire = 7 * 24 * time.Hour
)
