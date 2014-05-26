GoMq
====

GoMq 是基于Golang 和Redis 开发的消息队列系统，目前仅仅是beta版本。

支持简单的put get操作,采用http协议访问。项目地址：https://github.com/wujingke/GoMq

基于Redis 数据库，可以方便进行性能调优，和支持分布式的消息队列，Golang 自身又可作为webserver，二次开发也非常方便，全部代码才50多行，这仅仅是第一个beta版本，
后面考虑添加监控运行状态的功能和提升性能

==============================================
使用说明
put 操作
curl "http://127.0.0.1:1988/put?data=xxxxxxxxxx"

get 操作
curl "http://127.0.0.1:1988/get"
==============================================
安装＆使用
安装Redis之后，


git clone https://github.com/wujingke/GoMq.git
cd GoMq
go build
./GoMq
即可
