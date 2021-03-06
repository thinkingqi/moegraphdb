package graphdb

import (
	//"sync"
	"fmt"
	"encoding/json"
)

type (
	User struct{
		Info map[string]interface{}
		//Lock_fans  sync.RWMutex
		//Lock_likes sync.RWMutex
		Uid        uint
		Fans       map[uint]*User
		Likes      map[uint]*User
	}
)

func (this *User)String()string{
	info,_:=json.Marshal(this.Info)
	return "{ Uid:"+fmt.Sprint(this.Uid)+ ",Info:"+string(info)+",FansCount:"+fmt.Sprint(this.FansCount())+",LikesCount:"+fmt.Sprint(this.LikesCount()) +")"
}

// 粉丝数
func (this *User)FansCount()(int){
	if(this!=nil&&this.Fans!=nil){
		return len(this.Fans)
	}
	return 0
}
// 关注数
func (this *User)LikesCount()(int){
	if(this!=nil&&this.Likes!=nil){
		return len(this.Likes)
	}
	return 0
}

// 用户更多信息
func (this *User)SetInfo(info map[string]interface{}){
	this.Info=info
}

func (this *User)GetInfo()(info map[string]interface{}){
	info=this.Info
	return
}
