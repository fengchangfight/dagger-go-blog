# dagger-go-blog
一个用golang和vuejs(nuxt.js)实现的博客系统，轻快，简洁，对会编程的人来说，高度可定制，实例：www.fengchang.cc

## 说明
### 我嫌wordpress太慢故而手写的这个，所以速度应该还是不错的。
### 架构
前后端分离，所有组件：mysql, redis, nodejs环境（运行前端），golang

### 一些修改
本来是我个人的博客，但是为了共享代码，我删除了一些跟个人相关的信息，例如个人图片，config设置（包含mysql,redis等数据库密码和地址,端口等）。
如想直接代码拉下来就跑恐怕不行，不过，照着出错提示一步步调试（少了什么文件就加，哪个配置不对就改），应该很快能跑起来。毕竟这个代码我从头到尾写也只花了不到一周的业余时间。

### 数据库表
如果你安装了liquibase，直接运行daggerblogmisc中的liquibase-changelog.xml文件，不懂Liquibase的可以学一下，用这个比直接运行sql文件建表容易。
