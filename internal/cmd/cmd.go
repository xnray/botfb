package cmd

import (
	"botfb/internal/consts"
	"botfb/internal/controller"
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Name:        consts.AppName,
		Usage:       consts.AppName,
		Brief:       "start fb review bot server",
		Description: "this is the command entry for starting your process",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			startServer(ctx)
			return
		},
	}
	Version = &gcmd.Command{
		Name:        "version",
		Brief:       "view version",
		Description: "this is the command entry for view version",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Printf("current version: %s\n", consts.BotVersion)
			return nil
		},
	}
)

func startServer(ctx context.Context) {
	s := g.Server()
	group := s.Group("/")
	c := &controller.MsgReceive{}
	group.POST("/webhooks", c.HandleMessage)
	group.GET("/webhooks", c.VerifyHandler)
	go eventLoop(ctx)
	s.Run()
}

func eventLoop(ctx context.Context) {
	g.Log().Info(ctx, "run event loop")
	conn, err := g.Redis().Conn(ctx)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	defer conn.Close(ctx)
	_, err = conn.Subscribe(ctx, consts.EventTopic)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	appHandle := controller.NewEventHandle(ctx)
	for {
		msg, err := conn.ReceiveMessage(ctx)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		g.Log().Info(ctx, msg.Payload)
		if err := appHandle.Dispatch(ctx, msg.Payload); err != nil {
			g.Log().Warning(ctx, err)
		}
	}
}
