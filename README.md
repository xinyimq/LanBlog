## LanBlog 一站式个人博客解决方案
**感谢以下开源项目作者及参与者的无私奉献**
> * [Beego](https://github.com/astaxie/beego/)
> * [Vue](https://github.com/vuejs/vue)
> * [D2Admin](https://github.com/d2-projects/d2-admin)
> * 其他相关开源项目

**技术栈**
> Vue.js + axios(ajax) +Beego Restful api + Mysql + Nginx

### **项目介绍**

### 效果图
> 暂时只迁移一篇文章

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

#### 手动编译安装
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

**Step3 克隆项目**
> 拷贝项目到你指定的目录下,这里选择 /var/www/
> 没有安装git 需要先安装git  sudo apt install git

``` 
	cd /var/www/
	git clone https://github.com/sinksmell/LanBlog.git
```
**Step 4 安装编译环境**

```
	cd /usr/local
	wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz
	
``` 
