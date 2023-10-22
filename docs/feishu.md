```yml
route:
group_by: [ ... ]
group_wait: 60s
group_interval: 5m
repeat_interval: 1h
receiver: 'webhook'

receivers:
  - name: 'webhook'
    webhook_configs:
      - url: 'http://IP:8080/notifier/feishu?lang=zh-CN&url=https://open.feishu.cn/open-apis/bot/v2/hook/{your-token}'
        send_resolved: true

```
