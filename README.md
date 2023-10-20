<div align="center">

# <img src="https://github.com/whitexiong/Gatex-ui/raw/master/src/assets/logo.png" width="150" alt="GateX Logo"> GateX

**GateX** 是一个为现代网络应用程序设计的健壮网关解决方案。结合了 Go 的高效性与 Vue3 的灵活性，项目集成了聊天室功能模仿微信界面可添加AI用户进行问答，
同时集成了 openVpn 方便企业进行部署开发。

[![Go Version](https://img.shields.io/badge/Go-1.20-blue?style=flat&logo=go)](https://golang.org/)
[![Vue3](https://img.shields.io/badge/Vue-3.x-green.svg?style=flat&logo=vue.js)](https://vuejs.org/)
[![Element-UI](https://img.shields.io/badge/Element--UI-3.x-9cf.svg?style=flat)](https://element.eleme.io/)
[![License](https://img.shields.io/badge/license-proprietary-red.svg)](path_to_license.md)

</div>

## 🌟 特性

- 🔒 **身份验证**：使用 JWT 和 Casbin 进行身份认证和权限控制。
- 🌐 **路由管理**：轻松定义与管理 API 端点，RBAC 的权限模型。
- 🖥️ **现代用户界面**：采用 Vue3 与 Element-UI 构建。
- 🗣️ **聊天室**：对接了科大讯飞的星火模型、自己可扩展其他AI，使用了 websocket 进行在线聊天并可发送图片、视频。
- 🛡️ **OpenVPN**：一键生成证书、实时查看在线用户。
- 🛡️ **Gogs**：集成了 gogs 自建git仓库， 方便进行项目管理。

## 🔧 技术栈

- 🖥️ **后端**：[Go 1.20](https://golang.org/)
- 🎨 **前端**：[Vue3](https://vuejs.org/), [Element-UI](https://element.eleme.io/)，组件完全自定义。
- 📦 **包管理**：[npm 6.14.11](https://www.npmjs.com/), [node v14.16.0](https://nodejs.org/)
- 🔑 **访问控制**：[Casbin](https://casbin.org/)

## 🚀 如何启动项目

### 后端 (Go)

```bash
# 进入后端目录
cd backend

# 安装依赖
go mod tidy

# 启动后端服务
go run main.go


# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动前端开发服务器
npm run serve

访问前端应用：http://localhost:8080

```

## 🚀 未来计划

- 🛍️ 项目管理。
- 🔑 本地化CI/CD。
- 📊 集成监控与日志功能。
- 🚦 引入网关功能，如限流、断路器和负载均衡。
- 🌌 逐步转向微服务架构，确保更高的可扩展性和可维护性。

## 📸 项目预览

以下是项目的部分页面预览：

<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174129.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174143.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174147.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174150.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174154.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174158.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174202.png" width="400" height="auto">
<img src="https://github.com/whitexiong/Gatex-ui/raw/master/docs/images/20231019174205.png" width="400" height="auto">


## 🤝 贡献

🙌 欢迎感兴趣的开发者参与贡献。请参阅 [贡献指南](path_to_contributing_guide.md) 以获取详细信息。

## 📄 许可

🔒 该项目采用专有许可证。使用或实施该项目的任何部分均需要付费。详情请参阅 [LICENSE](path_to_license.md) 文件。

