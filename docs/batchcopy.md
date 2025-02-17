# 简介
`batchcopy` 命令用来将一个空间中的文件批量复制到另一个空间，另外你可以在复制的过程中，给文件进行重命名。

当然，如果所指定的源空间和目标空间相同的话，如果这个时候源文件和目标文件名相同，那么复制会失败（这个操作其实没有意义）。
如果复制的目标空间中存在同名的文件，那么默认情况下针对该文件的复制操作也会失败，如果希望强制覆盖，可以指定 `--overwrite` 选项。

# 格式
```
qshell batchcopy [--force] [--success-list <SuccessFileName>] [--failure-list <FailureFileName>] [--sep <Separator>]  [--worker <WorkerCount>] <SrcBucket> <DestBucket> [-i <SrcDestKeyMapFile>]
```

# 帮助
```
qshell batchcopy -h
```

# 鉴权
需要在使用了 `account` 设置了 `AccessKey`, `SecretKey` 和 `Name` 的情况下使用。

# 参数
- SrcBucket：原空间名，可以为公开空间或私有空间。【必选】
- DestBucket：目标空间名，可以为公开空间或私有空间。【必选】

# 选项
- -i/--input-file：该选项接受一个文件参数， 内容每行包含 `原文件名` 和 `目标文件名`，如果你希望 `目标文件名` 和 `原文件名` 相同的话，也可以不指定 `目标文件名`，那么这一行就是只有 `原文件名` 即可。每行多个元素名之间用分割符分隔（默认 tab 制表符）； 如果需要自定义分割符，可以使用 `-F` 或 `--sep` 选项指定自定义的分隔符。如果没有通过该选项指定该文件参数， 从标准输入读取内容。每行具体格式如下：（【可选】）
```
// 不指定目标文件名
<SrcKey>               // SrcKey：原文件名，copy 后目标文件名为 <SrcKey> 

// 指定目标文件名
<SrcKey><Sep><DestKey> // SrcKey：原文件名，<Sep>：分割符，DestKey：目标文件名。
```
- -y/--force：该选项控制工具的默认行为。默认情况下，对于批量操作，工具会要求使用者输入一个验证码，确认下要进行批量文件操作了，避免操作失误的发生。如果不需要这个验证码的提示过程可以使用此选项。【可选】
- -s/--success-list：该选项指定一个文件，qshell 会把操作成功的文件行导入到该文件；默认不导出。【可选】
- -e/--failure-list：该选项指定一个文件， qshell 会把操作失败的文件行加上错误状态码，错误的原因导入该文件；默认不导出。【可选】
- -F/--sep：该选项可以自定义每行输入内容中字段之间的分隔符（文件输入或标准输入，参考 -i 选项说明）；默认为 tab 制表符。【可选】】
- -c/--worker：该选项可以定义 Batch 任务并发数；1 路并发单次操作对象数为 250 ，如果配置为 10 并发，则 10 路并发单次操作对象数为 2500，此值需要和七牛对您的操作上限相吻合，否则会出现非预期错误，正常情况不需要调节此值，如果需要请谨慎调节；默认为 4。【可选】
- --min-worker：最小 Batch 任务并发数；当并发设置过高时，会触发超限错误，为了缓解此问题，qshell 会自动减小并发度，此值为减小的最低值。默认：1【可选】
- --worker-count-increase-period：为了尽可能快的完成操作 qshell 会周期性尝试增加并发度，此值为尝试增加并发数的周期，单位：秒，最小 10，默认 60。【可选】
- --overwrite：默认情况下，如果批量重命名的文件列表中存在目标空间已有同名文件的情况，针对该文件的重命名会失败，如果希望能够强制覆盖目标文件，那么可以使用 `--overwrite` 选项。【可选】
- --enable-record：记录任务执行状态，当下次执行命令时会检测任务执行的状态并跳过已执行的任务。 【可选】
- --record-redo-while-error：依赖于 --enable-record，当检测任务状态时（命令重新执行时，所有任务会从头到尾重新执行；任务执行前会先检测当前任务是否已经执行），如果任务已执行且失败，则再执行一次；默认此选项不生效，当任务执行失败不重新执行。 【可选】

# 示例
1 我们将空间 `if-pbl` 中的一些文件复制到 `if-pri `空间中去。如果是希望原文件名和目标文件名相同的话，可以这样指定 `SrcDestKeyMapFile` 的内容：
```
data/2015/02/01/bg.png
data/2015/02/01/pig.jpg
```

然后使用如下命令就可以把上面的文件就以和原来文件相同的文件名从 `if-pbl` 复制到 `if-pri` 了。
```
$ qshell batchcopy if-pbl if-pri -i tocopy.txt
```

2 如果上面希望在复制的时候，对一些文件进行重命名，那么 `SrcDestKeyMapFile` 可以是这样：
```
data/2015/02/01/bg.png	background.png
data/2015/02/01/pig.jpg
```
从上面我们可以看到，你可以为你希望重命名的文件设置一个新的名字，不希望改变的就不用指定。然后使用命令就可以将文件复制过去了。
```
$ qshell batchcopy if-pbl if-pri -i tocopy.txt
```

3 如果不希望上面的复制过程出现验证码提示，可以使用 `--force` 选项：
```
$ qshell batchcopy --force if-pbl if-pri -i tocopy.txt
```

4 如果目标空间存在同名的文件，可以使用 `--overwrite` 选项来强制覆盖：
```
$ qshell batchcopy --force --overwrite if-pbl if-pri -i tocopy.txt
```

5 如果希望导出复制成功和失败的列表， 可以使用 `--success-list` 和 `--failure-list` 选项
```
$ qshell batchcopy --success-list success.txt --failure-list failure.txt -i tocopy.txt if-pbl if-pri
```

6 如果源文件或者目的文件中包含了空格，那么文件需要使用其他的分隔符， 需要手动使用-F选项指定分隔符, 假如使用 `\t`
```
data/2015/02/01/bg.png\tbackgrou nd.png
data/2015/02/01/p\tig.jpg
```
可以使用如下命令处理:
```
$qshell batchcopy -i tocopy.txt  -F'\t' if-pbl if-pri
```

# 注意
如果没有指定输入文件的话， 会从标准输入读取同样内容格式

