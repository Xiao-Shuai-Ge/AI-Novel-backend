package ai

import (
	"Ai-Novel/application/ai/interfaces/api/internal/logic/ai"
	"Ai-Novel/application/ai/interfaces/api/internal/svc"
	"Ai-Novel/application/ai/interfaces/api/internal/types"
	"Ai-Novel/common/jwtx"
	"Ai-Novel/common/websocketx"
	"Ai-Novel/common/zlog"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 验证token（实际项目中替换为你的JWT实现）
		data, err := svcCtx.JWT.IdentifyToken(req.Token)
		if err != nil {
			logx.Errorf("token验证失败:%v", err)
			go httpx.ErrorCtx(r.Context(), w, errors.New("token验证失败"))
			return
		}

		// 验证token类型
		if data.Class != jwtx.AUTH_ENUMS_ATOKEN {
			logx.Errorf("token类型错误")
			httpx.ErrorCtx(r.Context(), w, errors.New("token类型错误"))
			return
		}

		// 获取用户ID
		userID, err := strconv.ParseInt(data.Userid, 10, 64)
		if err != nil {
			logx.Errorf("%s 转换 int64 错误: %v", data.Userid, err)
			httpx.ErrorCtx(r.Context(), w, errors.New("token转换错误"))
			return
		}

		// 升级连接
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Errorf("WebSocket upgrade error: %v", err)
			return
		}

		zlog.Infof("用户ID %d 连接 websocket", userID)

		// 注册客户端
		websocketx.WebsocketManager.Mutex.Lock()
		websocketx.WebsocketManager.Clients[conn] = userID
		websocketx.WebsocketManager.Users[userID] = conn
		websocketx.WebsocketManager.Mutex.Unlock()

		// 创建Logic实例处理连接
		wsLogic := ai.NewWebsocketLogic(r.Context(), svcCtx, userID, conn)
		go wsLogic.Websocket(&req)
	}
}
