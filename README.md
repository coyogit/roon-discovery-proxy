# Roon Discovery Proxy

將 Roon 的 Multicast 訊息 (239.255.90.90:9003) 轉發為 Broadcast，
解決手機 VPN 回家時無法搜尋 Roon Core 的問題。

## 使用方式

```bash
docker run --rm --network=host roon-discovery-proxy:local
