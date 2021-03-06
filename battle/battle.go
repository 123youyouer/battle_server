package battle

import (
	"../utils"
	"container/list"
	//"github.com/sirupsen/logrus"
)

const unit_id_offset	=1000;
const max_unit_count	=2000;
type BattlePanicError struct{
	E error;
}
type Battle struct {
	kcp_res_pool *utils.MemoryPool;
	udp_res_pool *utils.MemoryPool;
	all_units []*Unit;
	living_units *list.List;
	main_base_list *list.List;
}
func (context *Battle)AddMainBaseID(id uint16){
	context.main_base_list.PushBack(id);
}
func (context *Battle)GetFreeID()uint16{
	for i,u:=range context.all_units{
		if u==nil{
			return uint16(1000+i);
		}
	}
	return 0;
}
func (context *Battle)AllUnit()[]*Unit{
	return context.all_units;
}
func (context *Battle)FindUnit(id uint16)*Unit{
	return context.all_units[id-1000];
}
func (context *Battle)NewUnit(id uint16)*Unit{
	context.all_units[id-1000]=NewUnit(id);
	//logrus.Error("context.living_units.PushBack : ",id);
	context.living_units.PushBack(id);
	return context.all_units[id-1000];
}
func (context *Battle)ForEachUnitDo(f func(*Unit)(bool)){
	for _,u:=range context.all_units{
		if u==nil{
			continue;
		}
		if !f(u){
			return ;
		}
	}
}
func (context *Battle)FindUnitDo(id uint16,f func(*Unit)){
	f(context.FindUnit(id));
}
func (context *Battle)CreateUnitDo(f func(*Unit)){
	f(context.NewUnit(context.GetFreeID()));
}
func NewBattle1v1()*Battle{
	return &Battle{
		utils.NewKcpResPool(8),
		utils.NewUdpResPool(8),
		make([]*Unit,max_unit_count),
		list.New(),
		list.New(),
	};
}