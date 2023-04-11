# coupon-issue-project
회사에서 진행했던 쿠폰 발급 프로젝트 복습

redis 실행

```bash
docker-compose up
```

초당 100개 pop

```bash
k6 run --vus 100 --duration 30s script.js
```
