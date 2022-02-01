package e

type Err int

const (
	Success                 Err = 0
	InternalError               = 100
	InvalidParameter            = 101
	Unauthenticated             = 102
	CasbinAddPolicyError        = 201
	CasbinRemovePolicyError     = 202
	InsufficientPermission      = 203
	TokenEmpty                  = 301
	TokenMalformed              = 302
	TokenTimeError              = 303
	TokenSignError              = 304
	TokenRevoked                = 305
	DBCreateError               = 401
	RedisIncrError              = 402
	RedisGetError               = 403
	DbUpdateError               = 404
	UserDuplicated              = 501
	UserNotFound                = 502
	UserPasswordError           = 503
)

var msgText = map[Err]string{
	Success:                 "成功",
	InternalError:           "内部错误",
	InvalidParameter:        "参数错误",
	Unauthenticated:         "未登录",
	CasbinAddPolicyError:    "添加权限失败",
	CasbinRemovePolicyError: "删除权限失败",
	InsufficientPermission:  "权限不足",
	TokenEmpty:              "令牌为空",
	TokenMalformed:          "令牌格式错误",
	TokenTimeError:          "令牌有效期错误",
	TokenSignError:          "令牌签名错误",
	TokenRevoked:            "令牌已失效",
	DBCreateError:           "数据创建失败",
	RedisIncrError:          "Redis自增失败",
	RedisGetError:           "Redis获取失败",
	DbUpdateError:           "数据更新失败",
	UserDuplicated:          "用户已存在",
	UserNotFound:            "用户不存在",
	UserPasswordError:       "用户密码错误",
}

func (code Err) String() string {
	msg, ok := msgText[code]
	if ok {
		return msg
	}
	return msgText[InternalError]
}
