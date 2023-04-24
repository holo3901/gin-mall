package models

type ParamRegister struct {
	NickName   string `form:"nick_name" json:"nick_name" binding:"required"`
	UserName   string `form:"user_name" json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

type ParamProduct struct {
	Info     string `form:"info" json:"info"`
	PageSize int64  `form:"page_size" json:"page_size"`
	PageNum  int64  `form:"page_num" json:"page_num"`
}

type ParamUpdateUser struct {
	NickName string `form:"nick_name" json:"nick_name"`
	Email    string `form:"email" json:"email"`
	PassWord string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ParamUserValid struct {
	Email string `form:"email" json:"email"`
	Token string `form:"token" json:"token" binding:"required"`
}

type ParamProductService struct {
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
}

type Page struct {
	PageSize int64 `form:"pagesize" json:"pagesize"`
	PageNum  int64 `form:"pagenum" json:"pagenum"`
}

type ParamFavorite struct {
	ProductId int64 `form:"product_id" json:"product_id"`
	BossId    int64 `form:"boss_id" json:"boss_id"`
}

type ParamOrder struct {
	ProductID uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
	AddressID uint `form:"address_id" json:"address_id"`
	Money     int  `form:"money" json:"money"`
	BossID    uint `form:"boss_id" json:"boss_id"`
}

type ParamCarts struct {
	ProductID uint `form:"product_id" json:"product_id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	Num       uint `form:"num" json:"num"`
}
