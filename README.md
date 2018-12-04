# go_study
study go by example

主要是并发编程的实践和做的一些小实验
2018-12-02和2018-12-03两天的业余时间，将interface以及channel的基础做了下编码练习

接下来进行下channel的源代码分析巩固下练习中的心得

2018-12-04
当一个for循坏中有多个select语句来检测channel通道状态时，要避免前面的select阻塞在某个case语句上，从而导致后续的select得不到执行
在N生产者-M消费者 模式下如何优雅的关闭channel通道
select 和 timer结合

