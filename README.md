## LanBlog 一站式个人博客解决方案

[![Go Report Card](https://goreportcard.com/badge/github.com/sinksmell/LanBlog)](https://goreportcard.com/report/github.com/sinksmell/LanBlog)
[![GoDoc](https://godoc.org/github.com/sinksmell/LanBlog?status.svg)](https://godoc.org/github.com/sinksmell/LanBlog)
[![Build Status](https://travis-ci.com/sinksmell/LanBlog.svg?branch=master)](https://travis-ci.com/sinksmell/LanBlog)


**感谢以下开源项目作者及参与者的无私奉献**
> * [Beego](https://github.com/astaxie/beego/)
> * [Vue](https://github.com/vuejs/vue)
> * [D2Admin](https://github.com/d2-projects/d2-admin)
> * 其他相关开源项目

**技术栈**
> Vue.js + axios(ajax) +Beego Restful api + Mysql + Nginx

### **项目介绍**

### 效果图
> * 暂时只迁移一篇文章
> * [演示地址](http://47.101.222.133)

* 首页

![](https://i.loli.net/2019/02/28/5c77f2090412a.png)
* 侧边栏
![](https://i.loli.net/2019/02/28/5c77f2401504d.png)

* 阅读界面
![](https://i.loli.net/2019/02/28/5c77f2752fd85.png)

* 后台登录界面

![](https://i.loli.net/2019/02/28/5c77f4a302bd3.png)

* 后台管理界面
![](https://i.loli.net/2019/02/28/5c7773fb63702.png)


### 安装&使用
> 以Ubuntu为例

### 简单部署
> 下载对应的 压缩包 解压运行 具体步骤待补充...

### 手动编译安装
**Step1 安装mysql**

```shell
sudo apt update
sudo apt install mysql-server mysql-common mysql-client
```
安装完成后创建 myblog数据库或者其他名称,与项目目录conf下app.conf中mysqldb保持一致即可

``` shell
	mysql -u root -p
	// 进入mysql后创建数据库
	
	create database myblog;
	
	//创建完成后退出
	
	exit;
	
``` 

**Step2 安装Nginx**

``` shell
 sudo apt install nginx
```

**Step 3 安装编译环境**
> 若提示没有权限,请以root身份运行

* 下载并安装go语言,配置环境变量


``` shell
cd /usr/local
wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz

tar zxvf  go1.12.linux-amd64.tar.gz

echo 'export GOROOT=/usr/local/go' >> ~/.bashrc 

echo 'export GOPATH=/var/www' >> ~/.bashrc 
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc 

source ~/.bashrc
	
``` 

* 查看是否安装成功
> 输入go version查看go版本 输入go 查看命令提示


	go version
	go 

如果出现以下提示,则安装成功

![](https://i.loli.net/2019/03/03/5c7b8034bbdc4.png)

* 克隆项目到本地 

``` shell

cd /var/www

mkdir src

cd src

git clone https://github.com/sinksmell/LanBlog.git

```

* 安装依赖

``` 

go get github.com/astaxie/beego

go get  github.com/beego/bee

go get github.com/dgrijalva/jwt-go

go get github.com/go-sql-driver/mysql


```

**Step 3 安装编译环境**

* 修改Nginx配置文件

> 后台管理 blogCMS.conf

``` conf
server {
listen 8088;
server_name localhost;
charset utf-8;
access_log /var/www/blogCMS.log  main;

location / {
  root /var/www/src/LanBlog/views;
  index index.html;
}

location ~ /(script|image|img|js|fonts|css)/ {
  expires 1d;
  root /var/www/src/LanBlog/static/blogCMS ;
}
    location /api {
            proxy_pass   http://localhost:8088/v1;
            add_header Access-Control-Allow-Methods *;
            add_header Access-Control-Max-Age 3600;
            add_header Access-Control-Allow-Credentials true;
            add_header Access-Control-Allow-Origin $http_origin;
            add_header Access-Control-Allow-Headers $http_access_control_request_headers;

            if ($request_method = OPTIONS ) {
                return 200;
            }
        }

}
```

> blog.conf

``` conf
server {
    listen 80;
    server_name localhost;
    access_log /var/www/blog.log  main;

    location / {
        root /var/www/src/LanBlog/views ;
        index home.html ;
    }

    location  ~ /(css|js|fonts|img|node_modules)/ {
       # access_log off;
        expires 1d;
        root /var/www/src/LanBlog/static/blog ;
    }

    location /api {
            proxy_pass   http://localhost:8088/v1;
            add_header Access-Control-Allow-Methods *;
            add_header Access-Control-Max-Age 3600;
            add_header Access-Control-Allow-Credentials true;
            add_header Access-Control-Allow-Origin $http_origin;
            add_header Access-Control-Allow-Headers $http_access_control_request_headers;

            if ($request_method = OPTIONS ) {
                return 200;
            }
        }
}

```

* 运行项目 

``` 
sudo  service nginx start

cd /var/www/src/LanBlog

bee run 

```

* 运行效果图
![](https://i.loli.net/2019/03/03/5c7b81c7e4722.png)

* 访问80端口