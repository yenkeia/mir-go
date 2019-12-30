package main

import (
	"fmt"
	"github.com/yenkeia/mirgo/common"
)

type NPC struct {
	MapObject
	Info *common.NpcInfo // 仅在与数据库交互时使用
}

func (n *NPC) Point() common.Point {
	return *n.CurrentLocation
}

func NewNPC(m *Map, ni *common.NpcInfo) *NPC {
	return &NPC{
		MapObject: MapObject{
			ID:               m.Env.NewObjectID(),
			Name:             ni.Name,
			Map:              m,
			CurrentLocation:  common.NewPoint(ni.LocationX, ni.LocationY),
			CurrentDirection: common.MirDirectionDown,
		},
		Info: ni,
	}
}

func (n *NPC) String() string {
	return fmt.Sprintf("NPC Coordinate: %s, ID: %d, name: %s\n", n.Point().Coordinate(), n.ID, n.Name)
}
