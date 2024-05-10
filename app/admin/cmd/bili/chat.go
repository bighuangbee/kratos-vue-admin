package main

import (
	"fmt"
	bili "github.com/FishZe/go-bili-chat"
	handle "github.com/FishZe/go-bili-chat/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	bili.ChangeLogLevel(logrus.DebugLevel)
	bili.SetUID(443345583)
	bili.SetHeaderCookie("buvid3=4810AA4A-3821-9976-E15F-70F40DF0714342880infoc; b_nut=1713770742; _uuid=7DC987F9-A6810-6B92-9BB8-25482FF418E143270infoc; buvid4=EC88127B-D94F-9683-7823-CC65CAC1F06B44523-024042207-sXWsjx8Y0%2BfKPF7RnE3LKQ%3D%3D; CURRENT_FNVAL=4048; rpdid=|(u))lkk)lR|0J'u~uRJYJ|Yk; buvid_fp_plain=undefined; DedeUserID=443345583; DedeUserID__ckMd5=a5cbd6d706887823; CURRENT_QUALITY=80; fingerprint=a13a06b4392b162591f9bca7938f1611; enable_web_push=ENABLE; iflogin_when_web_push=1; header_theme_version=CLOSE; LIVE_BUVID=AUTO1017153259269380; SESSDATA=c9b6a27a%2C1730882069%2Cb516a%2A52CjAcCTt0EKSzcAy_j3Ls4df5NZOSuGGTl2JIHBaFexnonoKE2x-m00SGh5VKG1a3wn4SVkQ3d2UzU3Z5NFdCcUp2TWpEMDNzb3AwS0pwc093SXNRNWtBaU0zTmM1WE9NVUhFNEV0OEZibnR1RVc4Q29HNDM3elFhU2dGa1FJYjU0SXdLR3AyZzB3IIEC; bili_jct=869dead6b717926498271281ee75a1ee; b_lsid=102D71D10F_18F61A8F4B4; bp_t_offset_443345583=929827266749268021; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1OTAxNzMsImlhdCI6MTcxNTMzMDkxMywicGx0IjotMX0.KBLg_07TNo0xtGxl59KEmZtUBzetezE7DDVf9gyST8s; bili_ticket_expires=1715590113; bsource=search_bing; buvid_fp=a13a06b4392b162591f9bca7938f1611; PVID=29; sid=ehgx1s5d; home_feed_column=4; browser_resolution=1075-915")

	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	h.AddOption(handle.CmdDanmuMsg, 24577128, func(event handle.MsgEvent) {
		// 打印出弹幕消息
		fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
	})

	h.AddOption(handle.CmdInteractWord, 24577128, func(event handle.MsgEvent) {
		// 打印出弹幕消息
		fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
	})

	// 连接到直播间
	_ = h.AddRoom(24577128)
	//
	// 自定义AuthBody连接到直播间
	//
	//_ = h.AddRawRoom(client.WsAuthBody{
	//	UID:      0,
	//	Roomid:   26097368,
	//	Protover: 3,
	//	Platform: "web",
	//	Type:     2,
	//	Key:      "",
	//})
	//
	// 启动处理器
	h.Run()
}
