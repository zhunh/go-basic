# 本项目学习过程中的笔记

## 1. golang init函数
* init函数在包运行时作为初始化包内一些公用变量使用
##2. defer 
* 不宜大量使用，影响性能
* 在可能出错，有 return 语句终止的情况下使用，来保证终止之后仍能执行例如关闭数据库等操作
## 3. import _
*
## 4.sync.map
*
## 5.http.ResponseWriter  ->  WriteHeader

## 6.io.WriteString()

## 7.流控算法
* token bucket

## 8.中间件写法
* 实现handler接口
* 实现该接口的serverHTTP方法