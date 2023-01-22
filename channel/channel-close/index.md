## channel 操作
1. 退出
2. 阻塞

## channel 用途
1. 发送指令关闭所有协程，先在协程内部用 select 阻塞 <- close，在外部用 close 关闭协程

## channel 退出方式
1. 手动 close
2. 增加关闭条件
   ```
   ta := time.After(5 * time.Second)
   select {
   case <-ta:
   close(closed)
   }
    ```

## 用 channel 配合 close 做协程关闭的局限性
1. 每次要在协程内部增加对channel的判断
2. 也要在外部设置关闭条件
3. 如果程序要限制的是总时长，而不是单个操作的时长，这样每个操作要限制多少时间也是个难题。