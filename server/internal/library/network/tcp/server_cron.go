// Package tcp
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcp

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
)

// getCronKey 生成服务端定时任务名称
func (server *Server) getCronKey(s string) string {
	return fmt.Sprintf("tcp.server_%s_%s", s, server.name)
}

// stopCron 停止定时任务
func (server *Server) stopCron() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

// startCron 启动定时任务
func (server *Server) startCron() {
	// 心跳超时检查
	if gcron.Search(server.getCronKey(consts.TCPCronHeartbeatVerify)) == nil {
		_, _ = gcron.AddSingleton(server.Ctx, "@every 300s", func(ctx context.Context) {
			if server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.heartbeat < gtime.Timestamp()-consts.TCPHeartbeatTimeout {
					_ = client.Conn.Close()
					server.Logger.Debugf(server.Ctx, "client heartbeat timeout, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(consts.TCPCronHeartbeatVerify))
	}

	// 认证检查
	if gcron.Search(server.getCronKey(consts.TCPCronAuthVerify)) == nil {
		_, _ = gcron.AddSingleton(server.Ctx, "@every 300s", func(ctx context.Context) {
			if server.clients == nil {
				return
			}
			for _, client := range server.clients {
				if client.Auth.EndAt.Before(gtime.Now()) {
					_ = client.Conn.Close()
					server.Logger.Debugf(server.Ctx, "client auth expired, close conn. auth:%+v", client.Auth)
				}
			}
		}, server.getCronKey(consts.TCPCronAuthVerify))
	}
}
