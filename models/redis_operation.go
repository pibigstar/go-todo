package models

import (
	"encoding/json"
	"fmt"
	"github.com/pibigstar/go-todo/constant"
	"github.com/pibigstar/go-todo/models/db"
	"github.com/pkg/errors"
	"sort"
	"time"
)

// 跟redis相关的操作

type FormIDStruct struct {
	FormID string `json:"formId"`
	Expire time.Time `json:"expire"`
}

type formIDs []FormIDStruct

// 将收集的formID放入redis中
func CollectFormID(openID string, formID string) error{
	formIDStr, err := db.Redis.Get(fmt.Sprintf(constant.Collection_Form_ID_Prefix, openID)).Result()
	var formIds []FormIDStruct
	var data = []byte(formIDStr)
	json.Unmarshal(data, &formIds)

	newFormID := FormIDStruct{
		FormID:formID,
		Expire: time.Now().Add(constant.User_Form_ID_Expire),
	}

	formIds = append(formIds, newFormID)
	bytes, err := json.Marshal(formIds)
	if err!=nil {
		log.Error("解析formIds失败","err",err.Error())
		return err
	}
	_, err = db.Redis.Set(fmt.Sprintf(constant.Collection_Form_ID_Prefix, openID), string(bytes), constant.User_Form_ID_Expire).Result()
	if err!=nil{
		log.Error("存储formIds失败","err",err.Error())
		return err
	}
	return nil
}

// 获取收集的formID
func GetCollectionFormID(openID string)(string,error){
	formIDStr, _ := db.Redis.Get(fmt.Sprintf(constant.Collection_Form_ID_Prefix, openID)).Result()
	var formIds formIDs
	var data = []byte(formIDStr)
	json.Unmarshal(data, &formIds)

	// 排除过期的
	var newFormIds formIDs
	for _, id := range formIds {
		if id.Expire.After(time.Now()) {
			newFormIds = append(newFormIds, id)
		}
	}
	// 按过期时间排序
	sort.Sort(newFormIds)
	if newFormIds.Len() > 0{
		ids, formID := newFormIds.Remove()
		bytes, _ := json.Marshal(ids)
		db.Redis.Set(fmt.Sprintf(constant.Collection_Form_ID_Prefix, openID), string(bytes), constant.User_Form_ID_Expire).Result()
		return formID,nil
	}
	return "",errors.New("无可用的formID")
}


func (ids formIDs)Len()int {
	return ids.Len()
}

func (ids formIDs)Less(i, j int) bool {
	if ids[i].Expire.Before(ids[j].Expire){
		return true
	}
	return false
}

func (ids formIDs)Swap(i, j int) {
	temp := ids[j]
	ids[j] = ids[i]
	ids[i] = temp
}
// 移除第一个元素，并返回最新集合和移除元素
func (ids formIDs) Remove() ([]FormIDStruct,string)  {
	var newFormID []FormIDStruct
	var formID string
	if ids.Len() > 0 {
		for index,id := range ids {
			if index == 0 {
				formID = id.FormID
			} else {
				newFormID = append(newFormID, id)
			}
		}
	}
	return newFormID,formID
}