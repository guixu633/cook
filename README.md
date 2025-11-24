# Cook Project

这是一个前后端分离的项目，用于管理家庭菜单。

## 目录结构

- `frontend/`: React (Vite + TypeScript) 前端项目
- `backend/`: Golang 后端项目

## 本地开发 (Development)

建议使用项目根目录的 `Makefile` 进行管理。

1. **安装依赖**:
   ```bash
   make install
   ```

2. **启动数据库** (需要 Docker):
   ```bash
   make db-up
   ```

3. **启动服务**:
   打开两个终端分别运行：
   ```bash
   make run-backend  # 启动后端 (http://localhost:8080)
   make run-frontend # 启动前端 (http://localhost:5173)
   ```

## 部署 (Deployment)

目标服务器: `43.135.156.166`

### 1. 后端部署

构建 Linux 可执行文件:
```bash
make build-linux
```
这将生成 `backend/server-linux` 文件。

**部署步骤**:
1. 将 `backend/server-linux` 上传到服务器。
2. 将 `backend/config.yaml` 上传到服务器（与可执行文件同级）。
3. 确保服务器上安装了 PostgreSQL 或有可用的数据库连接。
4. 运行: `./server-linux`

### 2. 前端部署

构建静态资源:
```bash
make build-frontend
```
这将生成 `frontend/dist/` 目录。

**部署步骤**:
1. 将 `frontend/dist/` 目录上传到服务器（例如 `/var/www/cook`）。
2. 配置 Nginx 指向该目录。

### 前端连接远程后端

如果您在本地开发，但想连接远程的测试/生产后端，`vite.config.ts` 已配置代理指向 `43.135.156.166`。

如果您构建前端 (`npm run build`) 并部署，代码中 (`src/api/meal.ts`) 已配置在生产模式下直接请求 `http://43.135.156.166:8080/api`。
**注意**: 这需要后端开启 CORS 允许前端域名访问，或者使用 Nginx 反向代理将 `/api` 转发到后端端口。
