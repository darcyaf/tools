# 工具类

## 错误码描述
在API接口中需要特定描述
### 依赖
安装stringer安装包
`go get golang.org/x/tools/cmd/stringer`

### 使用
运行前需要执行`//go:generate stringer -type Code -linecomment -output code_string.go`,生成对应的String()方法
1. 直接使用code
  Code实现了`Error()`方法,所有可以直接在需要返回error中的方法中返回code,可以通过`code.String()`获取描述,`int(code)`获取code的值
2. 使用`NewError()`方法,相对于直接使用, 增加了打印日志到stderr的特性.

## iris的decoder
iris中虽然提供了`formDecoder.RegisterConverter()`来注册自定义的转换器,但是`ctx.ReadForm()`方法中调用的decoder是`schema`包中的一个内部变量,未提供出访问和设置的方法,
故而这里使用
```
//go:linkname formDecoder github.com/iris-contrib/schema.formDecoder
//go:nosplit
```
通过这个方式获取到内部变量从而实现注册,但是`go:linkname`会打破包的依赖结构啥的,建议还是少用.