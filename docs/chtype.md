# 简介
`chtype` 指令用来为空间中的一个文件修改 **存储类型**。

# 格式
```
qshell chtype <Bucket> <Key> <FileType>
```

# 鉴权
需要在使用了 `account` 设置了 `AccessKey` 和 `SecretKey` 的情况下使用。

# 参数
- Bucket：空间名，可以为公开空间或私有空间。【必选】
- Key：空间中的文件名。【必选】
- FileType：给文件指定的新的存储类型，其中可选值为 `0` 代表 `普通存储`，`1` 代表 `低频存储`，`2` 代表 `归档存储`，`3` 代表 `深度归档存储`。【必选】

注：
`归档存储` 直接转 `普通存储` 或 `低频存储` 会失败，需要通过再存储的方式。

# 示例
修改 `if-pbl` 空间中`qiniu.png` 图片的存储类型为 `低频存储`
```
$ qshell chtype if-pbl qiniu.png 1
```

修改完成，我们检查一下文件的存储类型：
```
$ qshell stat if-pbl qiniu.png
```

输出
```
Bucket:             if-pbl
Key:                qiniu.png
Hash:               FrUHIqhkDDd77-AtiDcOwi94YIeM
Fsize:              5331
PutTime:            14285516077733591
MimeType:           image/jpeg
FileType:           1 -> 低频存储
```
我们发现，文件的存储类型已经被修改为 `低频存储` 了。
