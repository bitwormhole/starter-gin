# Starter-Gin 项目

Starter-Gin 在 Starter 的基础上，实现了以依赖注入的方式，来配置并启动 Gin 服务。



## 安装

在安装 Starter-Gin 包之前, 需要先 [安装 Starter 框架](https://github.com/bitwormhole/starter) 。

1. 首先安装 Starter 框架, 然后你就能通过下列命令来安装 Starter-Gin 了。
 
        $ go get -u github.com/bitwormhole/starter-gin

2. 把它导入到你的代码:
 
        import ginstarter "github.com/bitwormhole/starter-gin"


## 快速开始
在工程文件夹下新建一个源文件，例如：example.go, 然后输入下列代码。

    package main

    import ginstarter "github.com/bitwormhole/starter-gin"

    func main() {
        i := ginstarter.InitGin()
        i.Use(ginstarter.Module())
        i.Run()
    }

运行这个例子：

    $ go run example.go

## 更多

要了解更多内容，请访问：
https://bitwormhole.com/starter-gin/
