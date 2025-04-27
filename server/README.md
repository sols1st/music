# Practice for mysql in golang

## 配置
### 在项目根目录新建 config.yaml 格式如下
```yaml
system:
  port: "8000"

mysql:
  host: "127.0.0.1"
  port: "3306"
  db-name: "learndb"
  username: "root"
  password: "xxxxxx"
  config: "charset=utf8mb4&parseTime=True&loc=Local"
```

## 课程选题
### 在线音乐分享系统
- 用户注册与登录
- 专辑与音乐编辑与发布
- 音乐点唱与评论
- 注册用户管理
- 专辑、音乐管理
