package main


import (
	"os"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli"
	"../server/restful"
	"../server/kcp_server"
	"../server/udp_server"
	"../world"
	"net/http"
	"time"
	"../server/service_discovery"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU());
	logrus.StandardLogger().Formatter.(*logrus.TextFormatter).TimestampFormat=time.StampMilli;
	app:=&cli.App{
		Name:"battle server",
		Usage:"frame sync server for battle",
		Version:"0.0.1",
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name:"bind",
				Value:"",
				Usage:"service bind address",
			},
			&cli.StringFlag{
				Name:"kcp",
				Value:"9090",
				Usage:"listen kcp",
			},
			&cli.StringFlag{
				Name:"udp",
				Value:"9091",
				Usage:"listen kcp",
			},
			&cli.StringFlag{
				Name:"rpc",
				Value:"9092",
				Usage:"listen tcp",
			},
			&cli.StringFlag{
				Name:"consul",
				Value:"10.0.0.101:8500",
				Usage:"consul address",
			},
		},
	};

	app.Action=func(c *cli.Context) error{
		w:=world.NewWorld();
		kcp_server.StartGateway(":"+c.String("kcp"),func(uid,rid uint32,session *kcp_server.KcpSession){
			defer func(){
				recover();
			}()
			if r:=w.FindRoom(rid);r!=nil{
				r.OnKcpSession(uid,session);
			}else{
				logrus.Error("can not find room ",rid," at session",session.RemoteAddr);
				session.Close(false);
			}
		});
		udp_server.StartGateway(":"+c.String("udp"));
		restful.NewRoomWS(w);
		restful.NewConsulCheck();
		service_discovery.RegisteServiceToConsul(
			c.String("consul"),
			c.String("bind")+":"+c.String("kcp")+","+c.String("udp")+","+c.String("rpc"),
			//c.String("bind")+":"+c.String("rpc")+"/consul/check",
			"10.0.0.6"+":"+c.String("rpc")+"/consul/check",
			);
		http.ListenAndServe(":"+c.String("rpc"),nil);
		return nil;
	}
	logrus.Info("server started at ",time.Now());
	app.Run(os.Args);
	logrus.Info("server stoped at ",time.Now());
}
