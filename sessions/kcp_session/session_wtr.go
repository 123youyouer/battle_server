package kcp_session

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)
type SendUtilErrorContext struct {
	s *Session;
	msg_puller func([]byte)int
	err_handle func(err error);
}
func (context *SendUtilErrorContext)WithSession(s*Session)*SendUtilErrorContext{
	context.s=s;
	return context;
}
func (context *SendUtilErrorContext)WithMsgPuller(f func([]byte)int)*SendUtilErrorContext{
	context.msg_puller=f;
	return context;
}
func (context *SendUtilErrorContext)WithErrHandle(f func(err error))*SendUtilErrorContext{
	context.err_handle=f;
	return context;
}
func (context *SendUtilErrorContext)SendUtilError(){
	buf:=make([]byte,1024);
	err:= func()(e error) {
		defer func(){
			if err:=recover();err!=nil{
				e=errors.New(fmt.Sprint(err));
			}
		}();
		for {
			if l:=context.msg_puller(buf);l>0{
				if _,e:=context.s.con.Write(buf[:l]);e!=nil{
					return e;
				}
			}
		}
		return nil;
	}();
	logrus.Error(err);
}