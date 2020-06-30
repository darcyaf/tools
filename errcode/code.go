package errcode

// 生成 code错误描述
//go:generate stringer -type Code -linecomment -output code_string.go

// 所有Code采用追加的方式
// 错误的命名以scope[api]-name-operation-state结构

// 数据库错误 3000+
const (
	DefaultDBCode Code = iota + 3000 // 服务器开小差了，请重试

)

// 输入错误 4000+
const (
	DefaultRequestCode Code = iota + 4000 // 参数错误，请重试
)

// 内部服务错误 5000+
const (
	DefaultServerError Code = iota + 5000 // 服务器开小差了，请重试
)

// 调用其他服务错误 6000+
const (
	DefaultOtherServerCode Code = iota + 6000 // 服务器开小差了，请重试
)

// 后台页面错误 7000+
const (
	_ Code = iota + 7000
)
