# gin-vue
```
Gin-vue是一个基于 vue 和 gin 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，
动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，支持运行于kubernetes。
目前处于开发阶段
```
# 贡献指南
###  Issue 规范
* issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。
* 在提交 issue 之前，请搜索相关内容是否已被提出。
### Pull Request 规范
* 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。
* commit 信息要以[文件名]: 描述信息 的形式填写，例如 README.md: fix xxx bug。
* 如果是修复 bug，请在 PR 中给出描述信息。
# 使用说明
```
- golang版本 >= v1.19
```
# server项目
### 本地运行
```
# 克隆项目
git clone https://github.com/zhime/gin-vue.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go build

# 运行
./server
```
### kubernetes运行
```
待完善。。。
```
