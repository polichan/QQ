# WeChat SDK for Go
![Go](https://github.com/polichan/qq/workflows/Go/badge.svg?branch=release-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/wechat)](https://goreportcard.com/report/github.com/silenceper/wechat)
[![pkg](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/polichan/qq?tab=doc)
![version](https://img.shields.io/badge/version-v2-green)

使用Golang开发的微信SDK，简单、易用。
>注意：当前版本为v2版本，v1版本已废弃


## 文档 && 例子
[API列表](https://github.com/polichan/qq/tree/v2/doc/api)

[Wechat SDK 2.0 文档](https://silenceper.com/wechat)

[Wechat SDK 2.0 例子](https://github.com/gowechat/example)


## 快速开始
```
import "github.com/polichan/qq"
```

以下是一个微信公众号处理消息接收以及回复的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
q := qq.NewQQ()
memory := cache.NewMemory()
cfg := &offConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Token:     "xxx",
    //EncodingAESKey: "xxxx",
    Cache: memory,
}
```

## 目录说明
- miniprogram: 小程序API
- doc: api文档

## 贡献
- 在[API列表](https://github.com/polichan/qq/tree/v2/doc/api)中查看哪些API未实现
- 提交issue，描述需要贡献的内容
- 完成更改后，提交PR

## License

Apache License, Version 2.0
