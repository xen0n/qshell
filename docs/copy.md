# 简介
`copy` 命令用来为存储在七牛空间中的文件创建副本。注意如果目标文件已存在空间中的时候，默认情况下，`copy` 会失败，报错 `614 file exists`，如果一定要强制覆盖目标文件，可以使用选项 `--overwrite` 。

参考文档：[资源复制 (copy)](http://developer.qiniu.com/code/v6/api/kodo-api/rs/copy.html)

# 格式
```
qshell copy [--overwrite] <SrcBucket> <SrcKey> <DestBucket> [-k <DestKey>]
```

# 鉴权
需要在使用了 `account` 设置了 `AccessKey` 和 `SecretKey` 的情况下使用。

# 参数
- SrcBucket: 源空间名称 【必选】
- SrcKey: 源文件名称 【必选】
- DestBucket: 目标空间名称，可以和源空间名称相同【必选】
  
# 选项
- -k: 目标文件名称(DestKey)，如果是 `DestBucket` 和 `SrcBucket` 不同的情况下，这个参数可以不填，默认和 `SrcKey` 相同。【可选】
- -w/--overwrite: 当保存的文件已存在时，强制用新文件覆盖原文件，如果无此选项操作会失败。【可选】

##### 备注：
1 如果复制的副本和原文件在同一个空间，那么必须提供不同于原文件的副本文件名，或者加上覆盖选项 `--overwrite`
2 如果复制的副本和原文件不在同一个空间，那么可以不提供副本文件名，默认和原文件名相同。
3 不支持跨存储区域复制文件, SrcBucket, DestBucket必须在统一存储区域

# 描述
1 复制`if-pbl`空间中的`qiniu.jpg`，并保存在`if-pbl`中，新副本文件名为`2015/01/19/qiniu.jpg`
```
$ qshell copy if-pbl qiniu.jpg if-pbl -k 2015/01/19/qiniu.jpg
```

2 复制`if-pbl`空间中的`qiniu.jpg`，并保存在`if-pri`中，新副本文件名和原文件名相同
```
$ qshell copy if-pbl qiniu.jpg if-pri
```

3 复制`if-pbl`空间中的`qiniu.jpg`，并保存到`if-pri`空间中，保存 Key 为：`qiniu_pri.jpg`，由于`if-pri`已有文件`qiniu_pri.jpg`，所以加上选项`--overwrite`强制覆盖
```
$ qshell copy --overwrite if-pbl qiniu.jpg if-pri -k qiniu_pri.jpg
```
