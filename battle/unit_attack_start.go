package battle

import (
	"fmt"
	"errors"
	"../utils"
)

func (context *Battle)UnitAttackStart(data []byte)(i interface{}){
	defer func(){
		if e:=recover();e!=nil{
			i=errors.New(fmt.Sprint(e));
		}
	}()
	res:=context.kcp_res_pool.Pop().(*utils.KcpRes);
	res.UID=0;
	res.Broadcast=true;
	wtr:=&packet_encoder{
		res.BDY,
		0,
	}
	ph1:=wtr.get_uint16_placeholder();
	ph2:=wtr.get_uint08_placeholder();
	wtr.write_bytes(data);
	ph2(utils.CMD_attack_start);
	ph1(uint16(wtr.pos)-2);
	return res;
}