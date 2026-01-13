# 内网节点NAT探测和打洞服务（子网）

## 功能

- 提供节点公网地址查询（STUN:Addr）。
- 支持节点NAT类型探测（STUN:Cone）。
- 支持节点NAT存活期探测（STUN:Live）。
- 协助节点UDP打洞（STUN:Punch）。


## 使用方式

```go
// 应用节点搜寻STUN服务
stunPeers := baseNet.Search("stun2-service")

// 连接到STUN服务子网
stunNet := p2p.New("stun2-service", stunPeers)

// 使用服务
natType := stunNet.Request("detect-nat-type")
liveTime := stunNet.Request("detect-living")
helpers := stunNet.Request("punch-help", targetAddr)
```
