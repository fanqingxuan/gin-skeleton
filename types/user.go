package types

type UserIndexReq struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Message string `json:"message" form:"message"`
}

type UserIndexReply struct {
	Message string `json:"message"`
}

type UserInfoReq struct {
	UserId uint `json:"id" form:"id" binding:"required"`
}

type UserInfoReply struct {
	Message string `json:"message"`
}

type UserAddReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Age      int64  `json:"age" form:"age"`
}

type UserAddReply struct {
	UserId int64 `json:"userid"`
}
