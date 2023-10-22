# prometheus-notifier

<div align=center>
<img src="https://img.shields.io/badge/golang-^1.21-red"/>
<img src="https://img.shields.io/badge/grafana-^10-green"/>
<img src="https://img.shields.io/badge/alertmanager-^0.26-blue"/>
</div>

### 已支持的通知方式

- [x] feishu 飞书群机器人
- [ ] feishu_app 飞书应用机器人
- [ ] dingtalk 钉钉
- [ ] wechat 企业微信
- [ ] email 邮箱
- [ ] telegram 电报


### 快速启动

```bash
docker run -d -p 8080:8080 lingcoder/prometheus-notifier:latest
```
