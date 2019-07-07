# go_study
目录
练习分为三部分：
1、基础语法
2、第三方库的使用
3、golang编码常见错误

一、基础语法
bufio_exercise
bufio包的使用

channel_behavior

channel_exercise

close_channel_graceful

defer_exercise

goroutine_exercise

https_request

interface_exercise

map_exercise

rateLimit_exercise

rpc_exercise

select_exercise

slice_exercise

二、第三方库使用
delve_exercise
演示golang调试工具delve的使用

elastic_exercise
golang调用elastic search的接口

goquery_selector_exercise
爬虫开发时会使用的选择器

iris_exercise
web开发

三、golang编程常见错误
7月7号开始 搜集golang中容易出错的用法或流程然后整理到github上



2018-12-04
当一个for循坏中有多个select语句来检测channel通道状态时，要避免前面的select阻塞在某个case语句上，从而导致后续的select得不到执行
在N生产者-M消费者 模式下如何优雅的关闭channel通道

select 和 timer结合（未完成）

2019-04-10（待完成）
golang 单元测试

2019-04-11（待完成）
gperf分析性能

