---
title: 开发环境
icon: gears
order: 5
category:
  - Guide
tag:
  - disable

navbar: false

breadcrumb: false
pageInfo: false
contributors: false
editLink: false
lastUpdated: false
prev: false
next: false
comment: false
footer: false

backtotop: false
---

## 依赖注入（wire）
~~~ shell
  # 安装wire
  go install github.com/google/wire/cmd/wire@latest
  # 注意将GOPATH/bin添加到path环境变量中
  
  # 测试输出
  wire help
  
  # 修改privader
  # 生成代码
  make wire
~~~