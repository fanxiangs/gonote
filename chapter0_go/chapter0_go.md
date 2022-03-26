# chapter0_go
## go doc
[go tour](https://golang.google.cn/tour/list)
[Language Specification](https://golang.google.cn/ref/spec)

[go tool](https://golang.google.cn/cmd/go/)

[go modules reference](http://docs.studygolang.com/ref/mod)

## [go modules 的基本使用](https://docs.studygolang.com/blog/using-go-modules)

介绍go modules的常见操作:

- 创建一个module

  ```shell
  go mod init
  ```

- 更新一个dependency或指定版本

  ```shell
  go get golang.org/x/text
  go get rsc.io/sampler@v1.3.1
  ```

- 删除无用的dependency

  ```shell
   go mod tidy
  ```

- 查看当前dependencies版本信息

  ```shell
  go list -m all
  ```

  

## [go modules reference](http://docs.studygolang.com/ref/mod)
