# 简介
`batchchgm` 命令用来批量修改七牛空间中文件的 MimeType。

# 格式
```
qshell batchchgm [--force] [--success-list <SuccessFileName>] [--failure-list <FailureFileName>] [--sep <Separator>]  [--worker <WorkerCount>] <Bucket> [-i <KeyMimeMapFile>] 
```

# 帮助
qshell batchchgm -h

# 鉴权
需要在使用了 `account` 设置了 `AccessKey`, `SecretKey` 和 `Name` 的情况下使用。

# 参数
- Bucket：空间名，可以为公开空间或私有空间。【必选】

# 选项
- -i/--input-file：该选项指定输入文件, 文件内容每行包含 `文件名称` 和 `新的 MimeType`；每行多个元素名之间用分割符分隔（默认 tab 制表符）； 如果需要自定义分割符，可以使用 `-F` 或 `--sep` 选项指定自定义的分隔符。 如果没有通过该选项指定该文件参数， 从标准输入读取内容；文件每行具体格式如下：（【可选】）
```
<Key><Sep><MimeType> // <Key>：文件名，<Sep>：分割符，<MimeType>：文件新的 MimeType。
```
- -y/--force：该选项控制工具的默认行为。默认情况下，对于批量操作，工具会要求使用者输入一个验证码，确认下要进行批量文件操作了，避免操作失误的发生。如果不需要这个验证码的提示过程可以使用此选项。【可选】
- -s/--success-list：该选项指定一个文件，qshell 会把操作成功的文件行导入到该文件；默认不导出。【可选】
- -e/--failure-list：该选项指定一个文件，qshell 会把操作失败的文件行加上错误状态码，错误的原因导入该文件；默认不导出。【可选】
- -F/--sep：该选项可以自定义每行输入内容中字段之间的分隔符（文件输入或标准输入，参考 -i 选项说明）；默认为 tab 制表符。【可选】
- -c/--worker：该选项可以定义 Batch 任务并发数；1 路并发单次操作对象数为 250 ，如果配置为 10 并发，则 10 路并发单次操作对象数为 2500，此值需要和七牛对您的操作上限相吻合，否则会出现非预期错误，正常情况不需要调节此值，如果需要请谨慎调节；默认为 4。【可选】
- --min-worker：最小 Batch 任务并发数；当并发设置过高时，会触发超限错误，为了缓解此问题，qshell 会自动减小并发度，此值为减小的最低值。默认：1【可选】
- --worker-count-increase-period：为了尽可能快的完成操作 qshell 会周期性尝试增加并发度，此值为尝试增加并发数的周期，单位：秒，最小 10，默认 60。【可选】
- --enable-record：记录任务执行状态，当下次执行命令时会检测任务执行的状态并跳过已执行的任务。 【可选】
- --record-redo-while-error：依赖于 --enable-record，当检测任务状态时（命令重新执行时，所有任务会从头到尾重新执行；任务执行前会先检测当前任务是否已经执行），如果任务已执行且失败，则再执行一次；默认此选项不生效，当任务执行失败不重新执行。 【可选】

# 示例
比如我们要将空间 `if-pbl` 中的一些文件的 MimeType 修改为新的值。
那么提供的 `KeyMimeMapFile` 的内容有如下格式：
```
data/2015/02/01/bg.png	image/png
data/2015/02/01/pig.jpg	image/jpeg
```

注意：上面文件名和MimeType中间的书写方式不是空格，而是制表符 `tab` 键，否则执行的时候不会报错，但也不会把MimeType(文件类型)批量修改成功。在上面的列表中， `data/2015/02/01/bg.png` 的新MimeType就是 `image/png`，诸如此类。

把上面的内容保存在文件 `tochange.txt` 中，然后使用如下的命令：
```
$ qshell batchchgm if-pbl -i tochange.txt
```

# 注意
如果没有指定输入文件的话, 默认会从标准输入读取同样格式的内容
