package battle

import (
	"errors"
	"fmt"
	"../utils"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)
func (context *Battle)each_unit_attack_done(who uint32,rdr *packet_decoder,wtr *packet_encoder)(){
	effect:=rdr.read_uint32();
	wtr.write_uint32(effect);
	count:=(int)(rdr.read_unit_count());
	wtr.write_unit_count((uint8)(count));
	for i:=0;i<count;i++{
		power:=(rdr.read_unit_attack_power())
		id:=rdr.read_unit_id();
		wtr.write_unit_id(id);
		u2:=context.FindUnit(id);
		if u2!=nil{
			if u2.HP>power{
				u2.HP-=power;
			}else{
				u2.HP=0;
			}
			u2.Killer=who;
			wtr.write_uint16(u2.HP);
		}else{
			wtr.write_uint16(0);
		}
	}
}
func (context *Battle)UnitAttackDone(who uint32,data []byte)(i interface{}){
	res:=(utils.I_RES)(nil);
	defer func(){
		if e:=recover();e!=nil{
			if res!=nil{
				res.Return();
			}
			i=&BattlePanicError{errors.New(fmt.Sprint(e))};
			logrus.Error(e);
			logrus.Error(fmt.Sprintf("%s",debug.Stack()));
		}
	}()
	res=context.kcp_res_pool.Pop().(utils.I_RES);
	res.SetUID(0);
	res.SetBroadcast(true);
	wtr:=&packet_encoder{
		res.GetWriteBuffer(),
		0,
	}
	rdr:=&packet_decoder{
		data:data,
		pos:0,
	}
	ph1:=wtr.get_uint16_placeholder();
	ph2:=wtr.get_uint08_placeholder();
	count:=(int)(rdr.read_unit_count());
	wtr.write_unit_count((uint8)(count));
	for i:=0;i<count;i++{
		context.each_unit_attack_done(who,rdr,wtr);
	}
	ph2(utils.CMD_attack_done);
	ph1((uint16)(wtr.pos)-2);
	return res;
}
