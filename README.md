# shortlinkd
Yet another URL shortener

# Features

- [x] 302 Redirect
- [x] Config hot reload

# Install

```bash
make
cp shortlinkd /usr/local/bin/
cp shortlinkd.service /etc/systemd/system/
touch /etc/shortlinkd.yaml
systemctl daemon-reload
systemctl start shortlinkd
```

Now access 127.0.0.1:10086 will show index page.

Add your link to config file,and run `systemctl reload shortlinkd` to apply new redirect rules.

# Config

```yaml
list:
  test: https://www.youtube.com/watch?v=dQw4w9WgXcQ
  test2: 
    url: https://google.com
    expired: 2024-06-01 12:22:34
    disabled: true
  test3: http://hao123.com
```
