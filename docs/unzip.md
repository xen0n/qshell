# 简介
`unzip` 命令用来解压 `zip` 文件。因为七牛支持的是UTF8编码的文件名，Windows自带的zip工具使用的是GBK编码的文件名。为了兼容这两种编码，所以有了 `unzip` 命令。

# 格式
```
qshell unzip <QiniuZipFilePath> [<UnzipToDir>]
```
 
# 参数
- QiniuZipFilePath：zip 文件路径。 【必须】
  
# 可选
- --dir：解压到指定目录，默认为命令运行的当前目录。 【可选】

# 示例
```
$ qshell unzip hellp.zip /home/Temp
```