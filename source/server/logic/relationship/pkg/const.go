package pkg

// RegisterType 注册类型
const (
	Phone = iota + 1
	Account
)

// PasswordCost 密码加密等级
const (
	PasswordCost = 12
)

// LoginStatus 登录状态
const (
	Login = iota + 1
	NotLogin
)

// RequestStatus 好友请求状态
const (
	Pending = "0"
	Agreed  = "1"
	Refused = "2"
)

// DefaultFriendGroup 默认好友分组
const (
	DefaultFriendGroup = "我的好友"
)

// 权限
const (
	Normal = 0
	Admin  = 1
	Leader = 2
)

// 逻辑删除
const (
	NotDeleted = 0
	Deleted    = 1
)

// SqlErrorFormat sql错误 log格式
const (
	SqlErrorFormat = "sql 执行错误: %v"
)

// TransactionErrorFormat 事务错误 log格式
const (
	TransactionErrorFormat = "事务执行错误: %v"
)

// GenFuncErrorFormat 生成函数错误 log格式
func GenFuncErrorFormat(funcName string) string {
	return funcName + " 执行错误"
}
