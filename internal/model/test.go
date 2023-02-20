package model

type SignUpParam struct {
	Age        uint8  `json:"age" form:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" form:"name" binding:"contains=tom,checkName"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password" form:"password" binding:"required"`
	RePassword string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
	Date       string `json:"date" form:"date" binding:"required,datetime=2006-01-02,checkDate"`
}
