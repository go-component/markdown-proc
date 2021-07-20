# markdown-proc

用于 markdown 图片、文字等等处理，目前只实现图片

# 限制
```shell
 Golang 版本 >= 1.11
```

# 安装
```html
go get -u github.com/go-component/markdown-proc
```

## 命令行参数解释

```shell
  -f string
        filepath of markdown
  -m int
        processing mode 
        0: image 
        1: word
        
        default 0
        
  -o string
        output path

```

- -f：本地 markdown 文件路径
- -m：处理模式，0：图片，1：文字，默认图片
- -o：加工后输出的路径

## 功能

### 图片加工

扫描远程图片链接，下载到本地并替换链接

```shell
markdown-proc -m 0 -f 1.md -o /output

或

markdown-proc -f 1.md -o /output
```

### 文字加工

TODO
