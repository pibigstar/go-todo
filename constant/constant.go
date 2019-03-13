package constant

import "time"

const (
	// 任务接收消息模板ID
	Tmeplate_Receive_Task_ID = "YWAd-Djf_hURHmf1HMobGSrd8AD53j_IoBJ35tffkPQ"
	// redis存储AccessToken前缀
	Redis_Prefix_Access_Token = "todo:accessToken:%s"
	// redis存储formID前缀
	Redis_Prefix_Form_ID = "todo:formID:%s"
	// 收集的用户的formID,OpenID
	Collection_Form_ID_Prefix  = "formIds:%s"
	// 收集的用户的formID的过期时间
	User_Form_ID_Expire = 7 * 24 * time.Hour
)
