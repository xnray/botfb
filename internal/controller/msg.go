package controller

import (
	"botfb/internal/consts"
	"botfb/internal/model/entity"
	"botfb/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type MsgReceive struct {
}

func (s *MsgReceive) HandleMessage(r *ghttp.Request) {
	body := r.GetBody()
	g.Log().Info(nil, string(body))
	// TODO parse receive message
	var data entity.Review
	data.CustomerId = "recipientID"
	// save message
	if err := service.Review().Create(nil, data); err != nil {
		r.Response.Status = 400
		return
	}
	r.Response.Status = 200
}

func (s *MsgReceive) VerifyHandler(r *ghttp.Request) {
	mode := r.FormValue("hub.mode")
	token := r.FormValue("hub.verify_token")
	challenge := r.FormValue("hub.challenge")
	g.Log().Infof(nil, "mode=%s, token=%s, challenge=%d", mode, token, challenge)
	// Check if a token and mode is in the query string of the request
	// Check the mode and token sent is correct
	if mode == consts.HubMode && token == consts.VerifyToken {
		// Respond with the challenge token from the request
		g.Log().Info(nil, "WEBHOOK_VERIFIED")
		//res.status(200).send(challenge)
		r.Response.Status = 200
		r.Response.Write(challenge)
	} else {
		// Respond with '403 Forbidden' if verify tokens do not match
		g.Log().Warning(nil, " verify tokens do not match")
		r.Response.Status = 403
	}
}
