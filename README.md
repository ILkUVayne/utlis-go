# utils-go

记录在Golang的工作学习中，一些常用的工具函数库，会持续的更新（ps: 更新时间随缘 \owo/ ~）

## Install

~~~bash
go get -u github.com/ILkUVayne/utlis-go/v2
~~~

## Usage

更多的已有的工具函数请查看[工具文档](https://github.com/ILkUVayne/utlis-go/blob/master/utils_doc.md)

~~~go
package main

import "github.com/ILkUVayne/utlis-go/v2/str"

func main() {
    s := "HelloWorld"
    println(str.CamelCaseToUnderscore(s))
}
~~~

