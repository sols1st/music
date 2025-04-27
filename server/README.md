# server for music platform

## 配置
### 在项目根目录新建 config.yaml 格式如下
```yaml
system:
  port: "3456"

mysql:
  host: "127.0.0.1" 
  port: "3306"
  db-name: "music-go"
  username: "root"
  password: ""
  config: "charset=utf8mb4&parseTime=True&loc=Local"

upload:
  upload_dir: "./upload"
  avatar_dir: "./upload/avatar"
  song_dir: "./upload/song"
  song_pic_dir: "./upload/songPic"
  song_lrc_dir: "./upload/songLrc"
  singer_pic_dir: "./upload/singer"
  song_list_pic_dir: "./upload/songListPic"
  banner_pic_dir: "./upload/banner"
  max-size: 10 # MB
  allowed-types:
    - ".mp3"
    - ".wav"
    - ".flac"
    - ".jpg"
    - ".png"
    - ".jpeg"

```
