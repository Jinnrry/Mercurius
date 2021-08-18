# Mercurius

> 墨丘利（拉丁語：Mercurius）是罗马神话中为众神传递信息的使者

Mercurius是一个使用Go语言编写的内网穿透工具。 相较于[frp](https://github.com/fatedier/frp) [ngrok](https://github.com/inconshreveable/ngrok) 等工具，Mercurius主要是为了解决传输过程中干扰问题。 frp ngrok的流量跨国时都会受到极大的干扰，导致运行相当不稳定。

Mercurius基于TCP实现了私有协议，使用AES-CBC-128对数据包加密，一定程度上保证了数据的安全性。

# Tips

目前项目处于beta阶段，并不能保证程序稳定。传输协议目前也缺少实践检验。后续可能考虑把server和client之间的通讯改成websocket，使用ssl加密，增加伪装性。


## 如何使用

[下载](https://github.com/Jinnrry/Mercurius/releases) 你对应的平台和架构，下载配置文件模板，修改配置文件。 

在内网机器上运行client程序：
```shell script
client -c ./config.json
```
在公网机器运行server程序:
```shell script
server -c ./config.json
```


## 配置文件说明

```json
{
  "common": {
    "token": "xbbbdasdf",  //传输秘钥
    "protocol": "tcp" //传输协议，目前只支持tcp,websocket协议正在开发中
  },
  "server": {
    "port": 11011,     //客户端和服务端通讯端口
    "ip": "127.0.0.1"  // 服务端ip
  },
  "client": {
    "services": [
      {
        "local_ip": "127.0.0.1",  // 需要代理的本地ip
        "local_port": 80,          //client端端口
        "remot_port": 8881,       // 外网端口
        "type": "tcp"     // 目前只支持tcp
      },
      {
        "local_ip": "127.0.0.1",
        "local_port": 3306,
        "remot_port": 3309,
        "type": "tcp"
      }
    ]
  }
}
```