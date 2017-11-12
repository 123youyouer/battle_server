package restful

import (
	"github.com/emicklei/go-restful"
	"../../world"
)

func NewRoomWS(w *world.World){
	ws:=new(restful.WebService);
	ws.
	Path("/room").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON);
	ws.Route(
		ws.	GET("/").
			To(
			func(req *restful.Request, res *restful.Response){
				list_rooms(req,res,w);
			}));

	ws.Route(
		ws.	GET("/count").
			To(
			func(req *restful.Request, res *restful.Response){
				count_room(req,res,w);
			}));

	ws.Route(
		ws.	GET("/{room_id}").
			To(
			func(req *restful.Request, res *restful.Response){
				get_room(req,res,w);
			}));

	ws.Route(
		ws.	PUT("/").
			To(
			func(req *restful.Request, res *restful.Response){
				new_room(req,res,w);
			}));
	ws.Route(
		ws. GET("/result/{room_guid}").
			To(
			func(req *restful.Request, res *restful.Response){
				get_room_result(req,res);
			}));

	restful.Add(ws);
}