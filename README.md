# 1. Todo
[![Build Status](https://travis-ci.org/pibigstar/go-todo.svg?branch=master)](https://travis-ci.org/pibigstar/go-todo)

> 此项目是todo小程序的后台，todo是一个任务发布提醒小程序，你可以加入一个组织，在里面可以给成员发布一些待完成的任务，主要服务于学校和一些公司之间，为了更加方便的管理任务需求而制作的一款小程序应用。

## 2. 使用技术

### 2.1 后端请求接收与处理
使用`goframe`框架

安装：
```go
go get -u github.com/gogf/gf
```

##### 服务绑定:
```go
func init() {
	s := g.Server()
	s.BindHandler("/wxLogin", wxLogin)
}
```
##### 数据校验与获取
```go
type WxLoginRequest struct {
	Code string `json:"code" gvalid:"type@required#code码不能为空"`
}

// 校验
if err := gvalid.CheckStruct(wxLoginRequest, nil); err != nil {
    log.Error("code为空", "err", err.String())
    r.Response.WriteJson(errorResponse(err.String()))
    return
}

// 获取前端请求数据
r.GetToStruct(wxLoginRequest)
```

##### 启动
```go
s := g.Server()
port := config.ServerConfig.Port
s.SetPort(int(port))
host := config.ServerConfig.Host
s.Domain(host)

// 开启日志
s.SetLogPath("log/todo.log")
s.SetAccessLogEnabled(true)
s.SetErrorLogEnabled(true)

s.Run()

```

### 2.2 配置文件读取
使用`viper`框架 

安装:
```go
go get -u github.com/spf13/viper
```
使用：
```go
// 设置配置文件名
configName := fmt.Sprintf("%s-%s", "config", ServerStartupFlags.Environment)
viper.SetConfigName(configName)
// 设置配置文件路径
viper.AddConfigPath("conf")
// 解析配置
viper.ReadInConfig()
// 获取server配置，map类型
viper.GetStringMap("server")
```

### 2.3 日志输出
使用`zap`框架

安装：
```go
go get -u go.uber.org/zap
```
使用：见`utils/log/log.go`


### 2.4 定时任务
使用`cron`框架

安装:
```go
go get -u github.com/robfig/cron
```
使用：
```go
c := cron.New()
	for _, job := range jobs.GetJobs() {
		log.Info("job启动", "job name", job.Name())
		c.AddFunc(job.Cron(), func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("job 运行出错", "job name", job.Name(), "error", err)
				}
			}()
			// 执行任务
			job.Run()
		})
	}
	c.Start()
	defer c.Stop()
```
## 3. 部署

### 3.1 打包成可执行文件
```bash
cd scripts
./build.bat
```

### 3.2 编译成镜像
```bash
docker build -t go-todo .
```
### 3.3 启动容器
```bash
docker run -dit -p 7410:7410 --name todo-container go-todo /bin/bash
```

### 3.4 进入容器
```bash
docker exec -it todo-container /bin/bash
```

### 3.5 删除镜像
```bash
# 停止容器
docker stop todo-container
# 删除容器
docker rm todo-container
# 删除镜像
docker rmi go-todo
```

## 相关项目

- [任务发布系统后端-go语言编写](https://github.com/pibigstar/go-todo)
- [任务发布系统小程序端](https://github.com/pibigstar/wx-todo)
- [任务发布系统后端-react编写](https://github.com/pibigstar/admin-todo)


## 项目结构
<details>
<summary>展开查看</summary>
<pre><code>.
├─conf
├─config
├─constant
├─controller
├─cron
│  └─jobs
├─https
├─log
│  └─todo.log
│      └─access
├─middleware
├─models
│  └─db
├─scritps
├─test
│  ├─config
│  ├─model
│  └─utils
├─utils
│  └─logger
└─vendor
</pre></code>
</details>

