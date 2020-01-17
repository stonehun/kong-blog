package models

type User struct {
	Base
	Name     string `gorm:"unique" json:"name"`
	Password string `json:"password"` //密码
	NickName string `json:"nick_name"`
	Email    string `json:"email"`  //邮箱
	Mobile   string `json:"mobile"` //手机
	Age      int    `json:"age"`    //年龄
	Remark   string `json:"remark"` //备注
}
