package room

import (
	"../battle"
	"errors"
	"time"
)


type RoomBuildContext struct {
	Lifecycle		int;
	SuddenDeath		int;
	WinScore		int;
}
type BattleRoomBuilder struct {
	r *Room1v1;
}

func NewBattleRoomBuilder(r *Room1v1)(*BattleRoomBuilder){
	return &BattleRoomBuilder{r};
}
func (me* BattleRoomBuilder)WithContext(context *RoomBuildContext)(*BattleRoomBuilder){
	me.r.lifecycle=time.Second*(time.Duration(context.Lifecycle));
	me.r.sudden_death=time.Second*(time.Duration(context.SuddenDeath));
	me.r.win_score=context.WinScore;
	return me;
}
func (me* BattleRoomBuilder)WithPlayers(i_player_getter ...interface{
	GetPlayerID()uint32;
	GetPlayerName()string;
	GetUnits()[]battle.Unit;
})(*BattleRoomBuilder){
	if len(i_player_getter)!=2{
		panic(errors.New("1v1 room need 2 players"));
		return me;
	}
	me.r.p1=&room_player{
		i_player_getter[0].GetPlayerID(),
		i_player_getter[0].GetPlayerName(),
		nil,
		nil,
	}
	me.r.p2=&room_player{
		i_player_getter[1].GetPlayerID(),
		i_player_getter[1].GetPlayerName(),
		nil,
		nil,
	}
	for i,_:=range i_player_getter[0].GetUnits(){
		me.r.the_battle.CreateUnitDo(func(unit *battle.Unit) {
			unit.SetAll(&i_player_getter[0].GetUnits()[i]);
			me.r.the_battle.AddMainBaseID(unit.ID);
			/*
			if unit.Type==200031{
				unit.Score=300;
			}else{
				unit.Score=100;
			}
			*/
		})
	}
	for i,_:=range i_player_getter[1].GetUnits(){
		me.r.the_battle.CreateUnitDo(func(unit *battle.Unit) {
			unit.SetAll(&i_player_getter[1].GetUnits()[i]);
			me.r.the_battle.AddMainBaseID(unit.ID);
			if unit.Type==200031{
				unit.Score=300;
			}else{
				unit.Score=100;
			}
		})
	}
	return me;
}
func BuildRoom1v1(
	context *RoomBuildContext,
	plrs ...interface{
		GetPlayerID()uint32;
		GetPlayerName()string;
		GetUnits()[]battle.Unit;
	},
)(*Room1v1,error){
	r:=NewBattleRoomBuilder(&Room1v1{
		new_base_room(battle.NewBattle1v1()),
		time.Second*120,
		1,
		nil,
		nil,
	}).WithContext(context).WithPlayers(plrs[0],plrs[1]).r;
	return r,nil;
}