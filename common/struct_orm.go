package common

type Basic struct {
	GameVersion   int `gorm:"primary_key"`
	CustomVersion int
	MapIndex      int
	MonsterIndex  int
	NPCIndex      int
	QuestIndex    int
	GameShopIndex int
	ConquestIndex int
	RespawnIndex  int
}

type GameShopItem struct {
	GameShopItemIndex int `gorm:"primary_key"`
	ItemIndex         int
	GoldPrice         int
	CreditPrice       int
	Count             int
	Class             string
	Category          string
	Stock             int
	IStock            int
	Deal              int
	TopItem           int
	//CreateDate
}

type ItemInfo struct {
	ItemIndex      int32 `gorm:"primary_key"`
	Name           string
	Type           ItemType
	Grade          ItemGrade
	RequiredType   RequiredType
	RequiredClass  RequiredClass
	RequiredGender RequiredGender
	ItemSet        ItemSet
	Shape          int16
	Weight         uint8
	Light          uint8
	RequiredAmount uint8
	Image          uint16
	Durability     uint16
	StackSize      uint32
	Price          uint32
	MinAC          uint8
	MaxAC          uint8
	MinMAC         uint8
	MaxMAC         uint8
	MinDC          uint8
	MaxDC          uint8
	MinMC          uint8
	MaxMC          uint8
	MinSC          uint8
	MaxSC          uint8
	HP             uint16
	MP             uint16
	Accuracy       uint8
	Agility        uint8
	Luck           int8
	AttackSpeed    int8
	StartItem      bool
	BagWeight      uint8
	HandWeight     uint8
	WearWeight     uint8
	Effect         uint8
	Strong         uint8
	MagicResist    uint8
	PoisonResist   uint8
	HealthRecovery uint8
	SpellRecovery  uint8
	PoisonRecovery uint8
	HpRate         uint8 // C# HRate
	MpRate         uint8 // C# MRate
	CriticalRate   uint8
	CriticalDamage uint8
	Bools          uint8
	MaxAcRate      uint8
	MaxMacRate     uint8
	Holy           uint8
	Freezing       uint8
	PoisonAttack   uint8
	Bind           int16
	Reflect        uint8
	HpDrainRate    uint8
	UniqueItem     int16
	RandomStatsId  uint8
	CanFastRun     bool
	CanAwakening   bool
	ToolTip        string
}

type MagicInfo struct {
	Name            string
	Spell           int
	BaseCost        int
	LevelCost       int
	Icon            int
	Level1          int
	Level2          int
	Level3          int
	Need1           int
	Need2           int
	Need3           int
	DelayBase       int
	DelayReduction  int
	PowerBase       int
	PowerBonus      int
	MPowerBase      int
	MPowerBonus     int
	MagicRange      int
	MultiplierBase  float32
	MultiplierBonus float32
}

type MapInfo struct {
	MapIndex        int    `gorm:"primary_key"`
	Filename        string `gorm:"Column:file_name"`
	Title           string
	MiniMap         int
	BigMap          int
	Music           int
	Light           int
	MapDarkLight    int
	MineIndex       int
	NoTeleport      int
	NoReconnect     int
	NoRandom        int
	NoEscape        int
	NoRecall        int
	NoDrug          int
	NoPosition      int
	NoFight         int
	NoThrowItem     int
	NoDropPlayer    int
	NoDropMonster   int
	NoNames         int
	NoMount         int
	NeedBridle      int
	Fight           int
	Fire            int
	Lightning       int
	NoTownTeleport  int
	NoReincarnation int
	NoReconnectMap  string
	FireDamage      int
	LightningDamage int
}

//type MineZone struct {
//	MapIndex  int `gorm:"primary_key"`
//	Mine      int
//	LocationX int `gorm:"Column:location_x"`
//	LocationY int `gorm:"Column:location_y"`
//	Size      int
//}

type MonsterInfo struct {
	MonsterIndex int `gorm:"primary_key"`
	Name         string
	Image        int
	AI           int `gorm:"Column:ai"`
	Effect       int
	Level        int
	ViewRange    int
	CoolEye      int
	HP           int `gorm:"Column:hp"`
	MinAC        int
	MaxAC        int
	MinMAC       int
	MaxMAC       int
	MinDC        int
	MaxDC        int
	MinMC        int
	MaxMC        int
	MinSC        int
	MaxSC        int
	Accuracy     int
	Agility      int
	Light        int
	AttackSpeed  int
	MoveSpeed    int
	Experience   int
	CanPush      int
	CanTame      int
	AutoRev      int
	Undead       int
}

type MovementInfo struct {
	MapIndex      int `gorm:"primary_key"`
	SourceX       int `gorm:"Column:source_x"`
	SourceY       int `gorm:"Column:source_y"`
	DestinationX  int `gorm:"Column:destination_x"`
	DestinationY  int `gorm:"Column:destination_y"`
	NeedHole      int
	NeedMove      int
	ConquestIndex int
}
type NpcInfo struct {
	MapIndex      int
	NpcIndex      int    `gorm:"primary_key"`
	Filename      string `gorm:"Column:file_name"`
	Name          string
	LocationX     int `gorm:"Column:location_x"`
	LocationY     int `gorm:"Column:location_y"`
	Rate          int
	Image         int
	TimeVisible   int
	HourStart     int
	MinuteStart   int
	HourEnd       int
	MinuteEnd     int
	MinLev        int
	MaxLev        int
	DayOfWeek     string
	ClassRequired string
	FlagNeeded    int
	Conquest      int
}

type QuestInfo struct {
	QuestIndex       int `gorm:"primary_key"`
	Name             string
	QuestGroup       string
	Filename         string `gorm:"Column:file_name"`
	RequiredMinLevel int
	RequiredMaxLevel int
	RequiredQuest    int
	RequiredClass    int
	QuestType        int
	GotoMessage      string
	KillMessage      string
	ItemMessage      string
	FlagMessage      string
}

type RespawnInfo struct {
	MapIndex        int
	MonsterIndex    int
	LocationX       int
	LocationY       int
	Count           int
	Spread          int
	Delay           int
	RandomDelay     int
	Direction       int
	RoutePath       string
	RespawnIndex    int
	SaveRespawnTime int
	RespawnTicks    int
}

type SafeZoneInfo struct {
	MapIndex   int
	LocationX  int
	LocationY  int
	Size       int
	StartPoint int
}
