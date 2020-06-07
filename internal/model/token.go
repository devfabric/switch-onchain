package model

//用户结构体
type UserInfo struct {
	UserName string `json:"username" example:"Tom"`    //机构用户名
	Pwd      string `json:"pwd"  example:"123456"`     //机构用户密码
	OrgCode  string `json:"orgcode" example:"orgcode"` //业务机构id
}

//请求token结构体
type ReqToken struct {
	UserInfo string                 `json:"userinfo" example:"{\"username\":\"client\",\"pwd\":\"12356\",\"orgcode\":\"731221\"}"` //用户请求token数据
	IsEnc    bool                   `json:"isenc" example:"false"`                                                                 //请求数据是否加密 true:加密 false:不加密
	CustomKV map[string]interface{} `json:"customkv,omitempty"`                                                                    //token中可以设置自定义KV
	Audience string                 `json:"audience" example:"switch-directory-chain"`                                             //请求发起者名称
	Expires  int64                  `json:"expires" example:"86400"`                                                               //token过期时间(单位秒) 默认一天
}

//token返回结构
type RespToken struct {
	Token   string `json:"token"`   //token字符串
	Expires int64  `json:"expires"` //token过期时间
}

//验证token
type ReqChkToken struct {
	UserName string `json:"username" example:"client"`                            //机构用户名
	OrgCode  string `json:"orgcode" example:"731221"`                             //业务机构id
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"` //token字符串
}
