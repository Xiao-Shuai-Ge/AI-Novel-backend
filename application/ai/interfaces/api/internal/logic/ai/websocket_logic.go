package ai

import (
	"Ai-Novel/common/websocketx"
	"Ai-Novel/common/zlog"
	"context"
	"github.com/gorilla/websocket"

	"Ai-Novel/application/ai/interfaces/api/internal/svc"
	"Ai-Novel/application/ai/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	userID int64
	conn   *websocket.Conn
}

// websocket接口
func NewWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext, userID int64, conn *websocket.Conn) *WebsocketLogic {
	return &WebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		userID: userID,
		conn:   conn,
	}
}

func (l *WebsocketLogic) Websocket(req *types.WebsocketReq) (resp *types.WebsocketResp, err error) {
	// 处理消息
	go func() {
		defer func() {
			websocketx.WebsocketManager.Mutex.Lock()
			delete(websocketx.WebsocketManager.Clients, l.conn)
			delete(websocketx.WebsocketManager.Users, l.userID)
			websocketx.WebsocketManager.Mutex.Unlock()

			zlog.Infof("用户ID %d 连接断开", l.userID)
			l.conn.Close()
		}()

		for {
			_, message, err := l.conn.ReadMessage()
			if err != nil {
				break
			}

			// 处理消息
			zlog.Infof("用户ID %d 发送消息 %s", l.userID, string(message))
			websocketx.WebsocketManager.Broadcast <- websocketx.Message{
				Content: "接受到消息: " + string(message),
				ToType:  "user",
				To:      l.userID,
			}
		}
	}()
	return &types.WebsocketResp{}, nil
}
