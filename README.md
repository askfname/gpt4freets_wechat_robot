## 🚨 由于 GPT4Free-TS 不再直接公开 forefront 的代码，所以基于 forefront 的 GPT4 无法使用，请在配置文件中设置其他站点，例如：chatdemo、you、phind、vita。

# gpt4freets_wechat_robot
调用 [GPT4Free-TS](https://github.com/xiangsx/gpt4free-ts) 的接口，实现个人微信接入 ChatGPT，和 GPT 机器人免费聊天。支持私聊回复和群聊艾特回复。

### 实现功能

* 机器人私聊回复
* 机器人群聊@回复
* 私聊回复前缀设置
* 好友添加自动通过可配置

### 常见问题
> 如无法登录：`login error: write storage.json: bad file descriptor`
删除掉 storage.json 文件重新登录。

> 如无法登录：`login error: wechat network error: Get "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage": 301 response missing Location header`
一般是微信登录权限问题，先确保 PC 端能否正常登录。

> 其他无法登录问题：
尝试删除掉 storage.json 文件，结束进程（linux 一般是 kill -9 进程 id）之后重启程序，重新扫码登录。
如果为 docket 部署，Supervisord 进程管理工具会自动重启程序。

> 机器人无法关联上下文：
上下文最大只支持 4 条，超过之后会自动切换账号。

### 使用前提
* 已经部署 [GPT4Free-TS](https://github.com/xiangsx/gpt4free-ts)
* 微信必须实名认证
* 微信必须绑定 +86 手机号

### 注意事项
* 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途
* 请注意收发敏感信息，本项目不做信息过滤

### Docker 运行
你可以使用 Docker 快速运行本项目。

#### 1. 基于环境变量运行

```sh
# 运行项目，环境变量参考下方配置说明
docker run -itd --name wechatbot --add-host=host.docker.internal:host-gateway --restart=always \
 -e AUTO_PASS=false \
 -e SESSION_TIMEOUT=60s \
 -e MODEL=gpt3.5-turbo \
 -e SITE=chatdemo \
 -e REPLY_PREFIX=来自GPT的回复: \
 -e TIMEOUT=150 \
 -e URL=http://host.docker.internal:3000 \
 nasheep/wechatbot:latest

# 查看二维码
docker exec -it wechatbot bash 
tail -f -n 50 /app/run.log 
```

运行命令中映射的配置文件参考下边的配置文件说明。

#### 2. 基于配置文件挂载运行

```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
cp config.dev.json config.json  # 其中 config.dev.json 从项目的根目录获取

# 运行项目
docker run -itd --name wechatbot --add-host=host.docker.internal:host-gateway -v `pwd`/config.json:/app/config.json nasheep/wechatbot:latest

# 查看二维码
docker exec -it wechatbot bash 
tail -f -n 50 /app/run.log 
```

其中配置文件参考下边的配置文件说明。


### 源码运行
适合了解 Go 语言编程的同学

````
# 获取项目
git clone https://github.com/nasheep/gpt4freets_wechat_robot.git

# 进入项目目录
cd gpt4freets_wechat_robot

# 复制配置文件
cp config.dev.json config.json

# 启动项目
go run main.go
````

### 配置说明

```json
{
  "auto_pass": true,                          // 是否自动通过好友添加
  "session_timeout": 60,                      // 会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文
  "model": "gpt3.5-turbo",                    // 模型类型，可选参数：gpt3.5-turbo、net-gpt3.5-turbo、gpt4
  "site": "chatdemo",                         // API 来源，可选参数：chatdemo、you、phind、vita、forefront
  "reply_prefix": "来自GPT的回复：",           // 私聊回复前缀
  "timeout": 150,                             // 请求 API 接口的超时时间（秒）
  "url": "http://host.docker.internal:3000"   // gpt4free-ts 的部署地址
}
```

### 友情提示
本项目是 fork 他人的项目来进行学习和使用，请勿商用，可以下载下来做自定义的功能。
项目基于 [openwechat](https://github.com/eatmoreapple/openwechat) 开发。
