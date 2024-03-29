# 我的电视·服务

启动一个本地的视频源服务。为了更好的稳定性、隐私性、便捷性，可以把一个视频源文件发布到一个本地地址，添加到本地视频播放器。

[my-tv-server](https://github.com/lizongying/my-tv-server)

## 功能

1. 支持原始输出
2. 支持json、txt、m3u之间格式转化
3. 支持发布到本地、局域网、公网
4. 支持linux、mac、windows系统
5. 支持配置到“我的电视·〇”等视频播放器

## 下载

[releases](https://github.com/lizongying/my-tv-0/releases/latest)

## 用法

1. 下载程序
2. 启动程序
3. 进入网址
4. 上传文件
5. 配置地址

```shell
./my-tv-server
```

* 默认端口8000，使用其他端口 `--port 8001`

![image](./screenshots/img.png)

## 测试

```shell
go run ./cmd/server/*
```

## 构建

```shell
make
```

## 镜像

[ghcr.io](https://github.com/lizongying/my-tv-server/pkgs/container/my-tv-server)

[hub.docker.com](https://hub.docker.com/r/lizongying/my-tv-server)

## 赞赏

![image](./screenshots/appreciate.jpeg)