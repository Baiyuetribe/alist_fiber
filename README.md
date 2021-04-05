# ALIST
 一款基于go+vue的阿里云盘,支持在线播放视频音乐，在线浏览文档图片。直连地址为阿里云oss，当前不限速，体验友好。

## 项目说明
最近学习go语言的过程中，最大的感触是语法虽少但是上手比较困难，正好看到这个项目，采用前后端分离开发，前端vue+后端go语言用的gin框架。所以，为了让自己掌握go语言，我使用go语言最新的fiber框架重构了后端，前后端都移除了部分功能。

## 在线演示地址：
- https://pan.baiyue.one
- 国内DEMO演示2： http://103.40.247.22:3000    [【该机器由茶猫云赞助,2天无理由退款+新购9折优惠】](https://www.chamaoyun.com/?u=D50390)     

![](img/demo.png)

## 部署方法

**请参考博客文章 https://baiyue.one/archives/1726.html**

## 命令说明：
二进制文件，可直接执行，相当于linux上的普通程序。
```bash
# 启动
./alist
# 为保持后台持续运行，使用nohup
nohup ./alist > log.log 2>&1 &

# 停止服务=终止进程
ps -ef|grep alist | grep -v grep | awk '{print $2}' | xargs kill -9
```

## 参考项目
- https://github.com/Xhofe/alist

## 其他自研开源程序
- 佰阅发卡 https://github.com/Baiyuetribe/kamiFaka
- 短视频去水印解析glink https://github.com/Baiyuetribe/glink