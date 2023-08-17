package msg

//import (
//	"fmt"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/paked/messenger"
//	"time"
//)
//
//func NewMsg() {
//	client := messenger.New(messenger.Options{
//		Verify:      false,
//		AppSecret:   "69eefe38875f9b9babaaa17ed5b13b02",
//		VerifyToken: "botfbkkk",
//		//Token:       "",
//	})
//	// Setup a handler to be triggered when a message is received
//	client.HandleMessage(func(m messenger.Message, r *messenger.Response) {
//		fmt.Printf("%v (Sent, %v)\n", m.Text, m.Time.Format(time.UnixDate))
//
//		p, err := client.ProfileByID(m.Sender.ID, []string{"name", "first_name", "last_name", "profile_pic"})
//		if err != nil {
//			fmt.Println("Something went wrong!", err)
//		}
//
//		r.Text(fmt.Sprintf("Hello, %v!", p.FirstName), messenger.ResponseType)
//	})
//	//addr := fmt.Sprintf("%s:%d", *host, *port)
//	//log.Println("Serving messenger bot on", addr)
//	//addr, err := gcfg.Instance().Get(nil, "server.address")
//	//if err != nil {
//	//	g.Log().Error(nil, err)
//	//	return
//	//}
//	s := g.Server()
//	s.BindHandler("/webhooks", client.Handler())
//	s.Run()
//	//log.Fatal(http.ListenAndServe(addr.String(), client.Handler()))
//}
