// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoResp struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	UserId       int64  `json:"userid"`
}

type ChangePwdReq struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

type ChangePwdResp struct {
	Pong int64 `json:"pong"`
}

type ListUserReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListUserItemResp struct {
	Id         int64  `json:"id"`
	CreateTime string `json:"createtime"`
	UpdateTime string `json:"updatetime"`
	DeleteTime string `json:"deletetime"`
	DelState   int64  `json:"delstate"`
	Version    int64  `json:"version"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	NickName   string `json:"nickname"`
	Sex        int64  `json:"sex"`
	Avatar     string `json:"avatar"`
	Info       string `json:"info"`
}

type ListUserResp struct {
	Total int64              `json:"total"`
	List  []ListUserItemResp `json:"list"`
}

type DeleteUserReq struct {
	Id int64 `json:"id"`
}

type DeleteUserResp struct {
	Pong int64 `json:"pong"`
}

type AddHomestayReq struct {
	CreateTime          string `json:"createtime,optional"`
	UpdateTime          string `json:"updatetime,optional"`
	DelTime             string `json:"deltime,optional"`
	Version             int64  `json:"version,optinal"`
	Title               string `json:"title"`
	SubTitle            string `json:"subtitle"`
	Banner              string `json:"banner"`
	Info                string `json:"info"`
	PeopleNum           int64  `json:"peoplenum,optional"`
	HomestayBusinessId  int64  `json:"homestaybusinessid,optional"`
	UserId              int64  `json:"userid,optional"`
	RowState            int64  `json:"rowstate,default=1"`
	FoodInfo            string `json:"foodinfo,optional"`
	FoodPrice           int64  `json:"foodprice,optional"`
	HomestayPrice       int64  `json:"homestayprice,optional"`
	MarketHomestayPrice int64  `json:"markethomestayprice,optional"` //insert use optional
}

type AddHomestayResp struct {
	Pong int64 `json:"pong""`
}

type ListHomestayReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListHomestayItemResp struct {
	Id                  int64  `json:"id"`
	CreateTime          string `json:"createtime"`
	UpdateTime          string `json:"updatetime"`
	DelTime             string `json:"deltime"`
	Version             int64  `json:"version"`
	Title               string `json:"tile"`
	SubTitle            string `json:"subtitle"`
	Banner              string `json:"banner"`
	Info                string `json:"info"`
	PeopleNum           int64  `json:"peoplenum"`
	HomestayBusinessId  int64  `json:"homestaybusinessid"`
	UserId              int64  `json:"userid"`
	RowState            int64  `json:"rowstate"`
	FoodInfo            string `json:"foodinfo"`
	FoodPrice           int64  `json:"foodprice"`
	HomestayPrice       int64  `json:"homestayprice"`
	MarketHomestayPrice int64  `json:"markethomestayprice"`
}

type ListHomestayResp struct {
	Total int64                  `json:"total"`
	List  []ListHomestayItemResp `json:"list"`
}

type DeleteHomestayReq struct {
	Id int64 `json:"id"`
}

type DeleteHomestayResp struct {
	Pong int64 `json:"pong"`
}

type UpdateHomestayReq struct {
	Id                  int64  `json:"id"`
	CreateTime          string `json:"createtime,optional"`
	UpdateTime          string `json:"updatetime,optional"`
	DelTime             string `json:"deltime,optional"`
	Version             int64  `json:"version,optinal"`
	Title               string `json:"tile"`
	SubTitle            string `json:"subtitle"`
	Banner              string `json:"banner"`
	Info                string `json:"info"`
	PeopleNum           int64  `json:"peoplenum,optional"`
	HomestayBusinessId  int64  `json:"homestaybusinessid,optional"`
	UserId              int64  `json:"userid,optional"`
	RowState            int64  `json:"rowstate,optional"`
	FoodInfo            string `json:"foodinfo,optional"`
	FoodPrice           int64  `json:"foodprice,optional"`
	HomestayPrice       int64  `json:"homestayprice,optional"`
	MarketHomestayPrice int64  `json:"markethomestayprice,optional"` //insert use optional
}

type UpdateHomestayResp struct {
	Pong int64 `json:"pong"`
}

type AddHomestayActivityReq struct {
	CreateTime string `json:"createtime,optional"`
	UpdateTime string `json:"updatetime,optional"`
	DeleteTime string `json:"deletetime,optional"`
	Version    int64  `json:"version,optinal"`
	DelState   int64  `json:"delstate,optional"`
	RowStatus  int64  `json:"rowstatus,default=1"`
	RowType    string `json:"rowtype,optional"`
	DataId     int64  `json:"dataid,optional"`
}

type AddHomestayActivityResp struct {
	Pong int64 `json:"pong"`
}

type DeleteHomestayActivityReq struct {
	Id int64 `json:"id"`
}

type DeleteHomestayActivityResp struct {
	Pong int64 `json:"pong"`
}

type ListHomestayActivityReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListHomestayAcitivityItemResp struct {
	Id         int64  `json:"id"`
	CreateTime string `json:"createtime"`
	UpdateTime string `json:"updatetime"`
	DeleteTime string `json:"deletetime"`
	Version    int64  `json:"version"`
	DelState   int64  `json:"delstate"`
	RowStatus  int64  `json:"rowstatus"`
	RowType    string `json:"rowtype"`
	DataId     int64  `json:"dataid"`
}

type ListHomestayActivityResp struct {
	Total int64                           `json:"total"`
	List  []ListHomestayAcitivityItemResp `json:"list"`
}

type UpdateHomestayActivityReq struct {
	Id         int64  `json:"id"`
	CreateTime string `json:"createtime,optional"`
	UpdateTime string `json:"updatetime,optional"`
	DeleteTime string `json:"deletetime,optional"`
	Version    int64  `json:"version,optinal"`
	DelState   int64  `json:"delstate,optional"`
	RowStatus  int64  `json:"rowstatus,optional"` //not default
	RowType    string `json:"rowtype,optional"`
	DataId     int64  `json:"dataid,optional"`
}

type UpdateHomestayActivityResp struct {
	Pong int64 `json:"pong"`
}

type AddHomestayBusinessReq struct {
	CreateTime  string  `json:"createtime,optional"`
	UpdateTime  string  `json:"updatetime,optional"`
	DeleteTime  string  `json:"deletetime,optional"`
	Version     int64   `json:"version,optinal"`
	DelState    int64   `json:"delstate,optional"`
	Title       string  `json:"title,optional"`
	UserId      int64   `json:"userid,optional"`
	Info        string  `json:"info,optional"`
	BossInfo    string  `json:"bossinfo,optional"`
	LicenseFron string  `json:"licensefron,optional"`
	LicenseBack string  `json:"licenseback,optional"`
	RowState    int64   `json:"rowstate,default=1"`
	Star        float64 `json:"star,optional"`
	Tag         string  `json:"tag,optional"`
	Cover       string  `json:"cover,optional"`
	HeaderImg   string  `json:"headerimg,optional"`
}

type AddHomestayBusinessResp struct {
	Pong int64 `json:"pong"`
}

type DeleteHomestayBusinessReq struct {
	Id int64 `json:"id"`
}

type DeleteHomestayBusinessResp struct {
	Pong int64 `json:"pong"`
}

type UpdateHomestayBusinessReq struct {
	Id          int64   `json:"id"`
	CreateTime  string  `json:"createtime,optional"`
	UpdateTime  string  `json:"updatetime,optional"`
	DeleteTime  string  `json:"deletetime,optional"`
	Version     int64   `json:"version,optinal"`
	DelState    int64   `json:"delstate,optional"`
	Title       string  `json:"title,optional"`
	UserId      int64   `json:"userid,optional"`
	Info        string  `json:"info,optional"`
	BossInfo    string  `json:"bossinfo,optional"`
	LicenseFron string  `json:"licensefron,optional"`
	LicenseBack string  `json:"licenseback,optional"`
	RowState    int64   `json:"rowstate,optional"`
	Star        float64 `json:"star,optional"`
	Tag         string  `json:"tag,optional"`
	Cover       string  `json:"cover,optional"`
	HeaderImg   string  `json:"headerimg,optional"`
}

type UpdateHomestayBusinessResp struct {
	Pong int64 `json:"pong"`
}

type ListHomestayBusinessReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListHomestayBusinessItemResp struct {
	Id          int64   `json:"id"`
	CreateTime  string  `json:"createtime"`
	UpdateTime  string  `json:"updatetime"`
	DeleteTime  string  `json:"deletetime"`
	Version     int64   `json:"version"`
	DelState    int64   `json:"delstate"`
	Title       string  `json:"title"`
	UserId      int64   `json:"userid"`
	Info        string  `json:"info"`
	BossInfo    string  `json:"bossinfo"`
	LicenseFron string  `json:"licensefron"`
	LicenseBack string  `json:"licenseback"`
	RowState    int64   `json:"rowstate"`
	Star        float64 `json:"star"`
	Tag         string  `json:"tag"`
	Cover       string  `json:"cover"`
	HeaderImg   string  `json:"headerimg"`
}

type ListHomestayBusinessResp struct {
	Total int64                          `json:"total"`
	LIst  []ListHomestayBusinessItemResp `json:"list"`
}

type AddHomestaycommentReq struct {
	CreateTime string  `json:"createtime,optional"`
	UpdateTime string  `json:"updatetime,optional"`
	DeleteTime string  `json:"deletetime,optional"`
	Version    int64   `json:"version,optinal"`
	DelState   int64   `json:"delstate,optional"`
	HomestayId int64   `json:"homwstayid,optional"`
	UserId     int64   `json:"userid,optional"`
	Star       float64 `json:"star,optional"`
	Content    string  `json:"content,optioanal"`
}

type AddHomestaycommentResp struct {
	Pong int64 `json:"pong"`
}

type DeleteHomestaycommentReq struct {
	Id int64 `json:"id"`
}

type DeleteHomestaycommentResp struct {
	Pong int64 `json:"pong"`
}

type ListHomestaycommentReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListHomestaycommentItemResp struct {
	Id         int64   `json:"id"`
	CreateTime string  `json:"createtime"`
	UpdateTime string  `json:"updatetime"`
	DeleteTime string  `json:"deletetime"`
	Version    int64   `json:"version"`
	DelState   int64   `json:"delstate"`
	HomestayId int64   `json:"homwstayid"`
	UserId     int64   `json:"userid"`
	Start      float64 `json:"start"`
	Content    string  `json:"content"`
}

type ListHomestaycommentResp struct {
	Total int64                         `json:"total"`
	List  []ListHomestaycommentItemResp `jsoan:"list"`
}

type UpdateHomestaycommentReq struct {
	Id         int64   `json:"id"`
	CreateTime string  `json:"createtime,optional"`
	UpdateTime string  `json:"updatetime,optional"`
	DeleteTime string  `json:"deletetime,optional"`
	Version    int64   `json:"version,optinal"`
	DelState   int64   `json:"delstate,optional"`
	HomestayId int64   `json:"homwstayid,optional"`
	UserId     int64   `json:"userid,optional"`
	Star       float64 `json:"star,optional"`
	Content    string  `json:"content,optioanal"`
}

type UpdateHomestaycommentResp struct {
	Pong int64 `json:"pong"`
}

type ListPaymentReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListPaymentItemResp struct {
	Id          int64  `json:"id"`
	Sn          string `json:"sn"`
	CreateTime  string `json:"createtime"`
	UpdateTime  string `json:"updatetime"`
	DeleteTime  string `json:"deletetime"`
	DelState    int64  `json:"delstate"`
	Version     int64  `json:"version"`
	UserId      int64  `json:"userid"`
	PayMode     string `json:"paymode"`
	TradeType   string `json:"tradetype"`
	TradeState  string `json:"tradete"`
	OrderSn     string `json:"ordersn"`
	ServiceType string `json:"servicetype"`
	PayStatus   string `json:"paystatus"`
	Paytime     string `json:"paytime"`
}

type ListPaymentResp struct {
	Total int64                 `json:"total"`
	List  []ListPaymentItemResp `json:"list"`
}

type AddOrderReq struct {
	CreateTime          string `json:"createtime,optional"`
	UpdateTime          string `json:"updatetime,optional"`
	DeleteTime          string `json:"deletetime,optional"`
	Version             int64  `json:"version,optional"`
	DelState            int64  `json:"delstate,optional"`
	Sn                  string `json:"sn,optional"`
	HomestayId          int64  `json:"homwstayid,optional"`
	UserId              int64  `json:"userid,optional"`
	Title               string `json:"tile,optional"`
	SubTitle            string `json:"subtitle,optional"`
	Cover               string `json:"cover,optional"`
	Info                string `json:"info,optional"`
	PeopleNum           int64  `json:"peoplenum,optional"`
	RowType             int64  `json:"rowtype,optional"`
	NeedFood            int64  `json:"needfood,optional"`
	FoodInfo            string `json:"foodinfo,optional"`
	FoodPrice           int64  `json:"foodprice,optional"`
	HomestayPrice       int64  `json:"homestayprice,optional"`
	MarketHomestayPrice int64  `json:"markethomestayprice,optional"`
	HomwstayBusinessId  int64  `json:"homestaybusinessid,optional"`
	HomestayUserId      int64  `json:"homestayuserid,optional"`
	LiveStartDate       int64  `json:"livestartdate,optional"`
	LiveEndDate         int64  `json:"liveenddate,optional"`
	LivePeopleNum       int64  `json:"livepeoplenum,optional"`
	TradeState          int64  `json:"tradestate,optional"`
	TradeCode           string `json:"tradecode,optional"`
	Remark              string `json:"remark,optional"`
	OrderTotalPrice     int64  `json:"ordertotalprice,optional"`
	FoodTotalPrice      int64  `json:"foodtotalprice,optional"`
	HomestayTotalPrice  int64  `json:"homestaytotalprice,optional"`
}

type AddOrderResp struct {
	Pong int64 `json:"pong"`
}

type UpdateOrderReq struct {
	Id                  int64  `json:"id"`
	CreateTime          string `json:"createtime,optional"`
	UpdateTime          string `json:"updatetime,optional"`
	DeleteTime          string `json:"deletetime,optional"`
	Version             int64  `json:"version,optional"`
	DelState            int64  `json:"delstate,optional"`
	Sn                  string `json:"sn,optional"`
	HomestayId          int64  `json:"homwstayid,optional"`
	UserId              int64  `json:"userid,optional"`
	Title               string `json:"tile,optional"`
	SubTitle            string `json:"subtitle,optional"`
	Cover               string `json:"cover,optional"`
	Info                string `json:"info,optional"`
	PeopleNum           int64  `json:"peoplenum,optional"`
	RowType             int64  `json:"rowtype,optional"`
	NeedFood            int64  `json:"needfood,optional"`
	FoodInfo            string `json:"foodinfo,optional"`
	FoodPrice           int64  `json:"foodprice,optional"`
	HomestayPrice       int64  `json:"homestayprice,optional"`
	MarketHomestayPrice int64  `json:"markethomestayprice,optional"`
	HomwstayBusinessId  int64  `json:"homestaybusinessid,optional"`
	HomestayUserId      int64  `json:"homestayuserid,optional"`
	LiveStartDate       int64  `json:"livestartdate,optional"`
	LiveEndDate         int64  `json:"liveenddate,optional"`
	LivePeopleNum       int64  `json:"livepeoplenum,optional"`
	TradeState          int64  `json:"tradestate,optional"`
	TradeCode           string `json:"tradecode,optional"`
	Remark              string `json:"remark,optional"`
	OrderTotalPrice     int64  `json:"ordertotalprice,optional"`
	FoodTotalPrice      int64  `json:"foodtotalprice,optional"`
	HomestayTotalPrice  int64  `json:"homestaytotalprice,optional"`
}

type UpdateOrderResp struct {
	Pong int64 `json:"pong"`
}

type ListOrderReq struct {
	Page     int64 `json:"page,default=1"`
	PageSize int64 `json:"pagesize,default=20"`
}

type ListOrderItemResp struct {
	Id                  int64  `json:"id"`
	CreateTime          string `json:"createtime"`
	UpdateTime          string `json:"updatetime"`
	DeleteTime          string `json:"deletetime"`
	Version             int64  `json:"version"`
	DelState            int64  `json:"delstate"`
	Sn                  string `json:"sn"`
	HomestayId          int64  `json:"homwstayid"`
	UserId              int64  `json:"userid"`
	Title               string `json:"tile"`
	SubTitle            string `json:"subtitle"`
	Cover               string `json:"cover"`
	Info                string `json:"info"`
	PeopleNum           int64  `json:"peoplenum"`
	RowType             int64  `json:"rowtype"`
	NeedFood            int64  `json:"needfood"`
	FoodInfo            string `json:"foodinfo"`
	FoodPrice           int64  `json:"foodprice"`
	HomestayPrice       int64  `json:"homestayprice"`
	MarketHomestayPrice int64  `json:"markethomestayprice"`
	HomwstayBusinessId  int64  `json:"homestaybusinessid"`
	HomestayUserId      int64  `json:"homestayuserid"`
	LiveStartDate       int64  `json:"livestartdate"`
	LiveEndDate         int64  `json:"liveenddate"`
	LivePeopleNum       int64  `json:"livepeoplenum"`
	TradeState          int64  `json:"tradestate"`
	TradeCode           string `json:"tradecode"`
	Remark              string `json:"remark"`
	OrderTotalPrice     int64  `json:"ordertotalprice"`
	FoodTotalPrice      int64  `json:"foodtotalprice"`
	HomestayTotalPrice  int64  `json:"homestaytotalprice"`
}

type ListOrderResp struct {
	Total int64               `json:"total"`
	Lisrt []ListOrderItemResp `json:"list"`
}

type DeleteOrderReq struct {
	Id int64 `json:"id"`
}

type DeleteOrderResp struct {
	Pong int64 `json:"pong"`
}
