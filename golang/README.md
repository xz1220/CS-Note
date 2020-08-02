# Golang

### RunTime
---
[Go 调度源码解析](https://mp.weixin.qq.com/s/T2nqjNlfGmBTOIeF6jidYA)
- 调度器会在G的时间片有剩余的时候，优先执行G最后一个创建的goroutine，并且将之前创建的goroutine按顺序放到当前G本地队列中
- 不要依赖调度器来保证goroutine的调度顺序

---
[scheduler 源码解析 大佬写的](https://github.com/cch123/golang-notes/blob/master/scheduler.md)

- TODO：还没有细看，需要汇编基础