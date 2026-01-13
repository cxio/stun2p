// Copyright (c) 2026 @cxio/stun2p
// Released under the MIT license
//////////////////////////////////////////////////////////////////////////////
//
// STUN2 客户端的一个实现。
// 向服务器请求 NAT 类型探测（STUN:Cone）和生存期探测（STUN:Live）。
//
// 用法（应用端）：
// - 创建一个 Client 实例。
// - 向基网查询STUN2服务节点，连接到STUN2服务子网（stun2-service）。
// - 向服务节点请求NAT相关服务：STUN:Addr|STUN:Cone|STUN:Live。
//
// STUN:Addr:
// - 客户端向服务器发送请求（STUN:Addr），获取自身的公网地址。
// - 可以向另一台服务节点发送相同请求以进行对比。
// - 此也为 STUN:Cone 的前置预探测。
// - 此时的 UDP 链路为 QUIC 安全连接。
//
// 注：详细操作请参考 README.md 中的 STUN:Addr 部分。
//
// STUN:Cone:
// - 客户端向服务器请求NAT类型探测（STUN:Cone），发送必要的信息。
// - 客户端断开 QUIC 连接，回退到普通UDP链路。
// - 服务器根据预定的规则构造SN，然后在该普通UDP链路上发送探测包。
// - 服务器用一个新的IP向客户端发送 NewHost 探测包，或者请求另一台服务器协助。
// - 客户端接收 UDP 消息，综合判断自己的 NAT 类型。
//
// 注：详细操作请参考 README.md 中的 STUN:Cone 部分。
//
// STUN:Live:
// - 客户端首先向服务器请求探知自己的公网地址（STUN:Addr）。
// - 客户端断开此 QUIC 连接，回退到普通UDP连接（旧链路）。
// - 客户端向一台新的服务器请求NAT生命期探测（STUN:Live），QUIC 连接。
// - 客户端向新服务器发送必要的信息（起始计数、SN因子、旧链路端口）。
// - 客户端在旧链路上发送 PING 保活消息（刷新NAT映射），计时器重置。
// - 服务器构造SN，在纯UDP的旧链路上发送探测包（3个，冗余）。
// - 客户端控制请求时间间隔，通过多次请求，测试NAT映射的存活期。
//
// 注：详细操作请参考 README.md 中的 STUN:Live 部分。
//
//////////////////////////////////////////////////////////////////////////////
//

package client

import (
	"errors"

	"github.com/cxio/stun2"
)

// NatLevel NAT类型层级引用
type NatLevel = stun2.NatLevel

// STUNTester NAT测试类型
type STUNTester int

const (
	STUN_ADDR STUNTester = 1 + iota // STUN:Addr
	STUN_CONE                       // STUN:Cone
	STUN_LIVE                       // STUN:Live
)

var (
	// UDP拨号错误
	ErrDialUDP = errors.New("dial to server udp failed")

	// 客户端UDP地址未探测
	ErrNotAddr = errors.New("client udp addr is empty")
)

// Client 请求NAT探测的通用客户端。
type Client struct {
	Tester chan STUNTester // STUN 测试类型通知
	// ...
}

// Close 关闭客户端。
func (c *Client) Close() {
	close(c.Tester)
	// ...
}
