package typestruct

import (
	"math"
	"net"
	"sync"

	. "github.com/KouKouChan/CSO2-Server/blademaster/Exp"
)

type (
	User struct {
		//个人信息
		Userid               uint32 `json:"-"`
		UserName             []byte
		IngameName           []byte
		Password             []byte
		Gm                   uint8
		Level                uint16
		Rank                 uint8
		RankFrame            uint8
		Points               uint64
		CurrentExp           uint64
		MaxExp               uint64
		PlayedMatches        uint32
		Wins                 uint32
		Kills                uint32
		Headshots            uint32
		Deaths               uint32
		Assists              uint32
		Accuracy             uint16
		SecondsPlayed        uint32
		NetCafeName          []byte
		Cash                 uint32
		ClanName             []byte
		ClanMark             uint32
		WorldRank            uint32
		Campaign             uint16
		Mpoints              uint32
		TitleId              uint16
		UnlockedTitles       []byte
		Signature            []byte
		BestGamemode         uint32
		BestMap              uint32
		UnlockedAchievements []byte
		Avatar               uint16
		UnlockedAvatars      []byte
		VipLevel             uint8
		VipXp                uint32
		SkillHumanCurXp      uint64
		SkillHumanMaxXp      uint64
		SkillHumanPoints     uint8
		SkillZombieCurXp     uint64
		SkillZombieMaxXp     uint64
		SkillZombiePoints    uint8
		UserMail             []byte
		//连接
		CurrentConnection net.Conn `json:"-"`
		//频道房间信息
		CurrentChannelServerIndex uint8       `json:"-"`
		CurrentChannelIndex       uint8       `json:"-"`
		CurrentRoomId             uint16      `json:"-"`
		CurrentTeam               uint8       `json:"-"`
		Currentstatus             uint8       `json:"-"`
		CurrentIsIngame           bool        `json:"-"`
		CurrentSequence           *uint8      `json:"-"`
		CurrentKillNum            uint16      `json:"-"`
		CurrentDeathNum           uint16      `json:"-"`
		CurrentAssistNum          uint16      `json:"-"`
		NetInfo                   UserNetInfo `json:"-"`
		//仓库信息
		Inventory   UserInventory
		WeaponKills map[uint32]uint32

		UserMutex *sync.Mutex `json:"-"`
	}

	UserNetInfo struct {
		ExternalIpAddress  uint32
		ExternalClientPort uint16
		ExternalServerPort uint16
		ExternalTvPort     uint16

		LocalIpAddress  uint32
		LocalClientPort uint16
		LocalServerPort uint16
		LocalTvPort     uint16
	}
)

const (
	//user status
	UserNotReady = 0
	UserIngame   = 1
	UserReady    = 2

	//阵营
	UserForceUnknown          = 0
	UserForceTerrorist        = 1
	UserForceCounterTerrorist = 2
)

func (u User) IsVIP() bool {
	if u.VipLevel <= 0 {
		return false
	}
	return true
}

func (u *User) SetVIP() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.VipLevel = 1
}

func (u *User) SetGM() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Gm = 1
}

func (u *User) SetID(id uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Userid = id
}

func (u *User) SetNewMutex() {
	if u == nil {
		return
	}
	var mutex sync.Mutex
	u.UserMutex = &mutex
}

func (u *User) SetUserName(loginName, username []byte) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.UserName = loginName
	u.IngameName = username
}

func (u *User) SetUserChannelServer(id uint8) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentChannelServerIndex = id
}

func (u *User) SetUserChannel(id uint8) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentChannelIndex = id
}

func (u *User) SetUserRoom(id uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentRoomId = id
}

func (u *User) QuitChannel() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentChannelIndex = 0
}

func (u *User) QuitRoom() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentRoomId = 0
	u.CurrentTeam = UserForceUnknown
	u.Currentstatus = UserNotReady
	u.CurrentIsIngame = false
}

func (u *User) JoinRoom(id uint16, team uint8) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentRoomId = id
	u.CurrentTeam = team
	u.Currentstatus = UserNotReady
	u.CurrentIsIngame = false
}

func (u *User) SetUserTeam(team uint8) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentTeam = team
}

func (u *User) SetUserStatus(status uint8) {
	if u == nil {
		return
	}
	if status <= 2 &&
		status >= 0 {
		u.UserMutex.Lock()
		defer u.UserMutex.Unlock()
		u.Currentstatus = status
	}
}

//GetUserChannelServerID 获取用户所在分区服务器ID
func (u User) GetUserChannelServerID() uint8 {
	if u.Userid <= 0 {
		return 0
	}
	return u.CurrentChannelServerIndex
}

//获取用户所在频道ID
func (u User) GetUserChannelID() uint8 {
	if u.Userid <= 0 {
		return 0
	}
	return u.CurrentChannelIndex
}

//获取用户所在房间ID
func (u User) GetUserRoomID() uint16 {
	if u.Userid <= 0 {
		return 0
	}
	return u.CurrentRoomId
}

func (u User) GetUserTeam() uint8 {
	return u.CurrentTeam
}

func (u User) IsUserReady() bool {
	return u.Currentstatus == UserReady
}

func (u *User) SetUserIngame(ingame bool) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentIsIngame = ingame
	if ingame {
		u.Currentstatus = UserIngame
	} else {
		u.Currentstatus = UserNotReady
	}

}

func (u *User) UpdateHolepunch(portId uint16, localPort uint16, externalPort uint16) uint16 {
	if u == nil {
		return 0xFFFF
	}
	switch portId {
	case UDPTypeClient:
		u.UserMutex.Lock()
		defer u.UserMutex.Unlock()
		u.NetInfo.LocalClientPort = localPort
		u.NetInfo.ExternalClientPort = externalPort
		return 0
	case UDPTypeServer:
		u.UserMutex.Lock()
		defer u.UserMutex.Unlock()
		u.NetInfo.LocalServerPort = localPort
		u.NetInfo.ExternalServerPort = externalPort
		return 1
	case UDPTypeSourceTV:
		u.UserMutex.Lock()
		defer u.UserMutex.Unlock()
		u.NetInfo.LocalTvPort = localPort
		u.NetInfo.ExternalTvPort = externalPort
		return 2
	default:
		return 0xFFFF
	}
}

func (u *User) CountKillNum(num uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentKillNum = num
}

func (u *User) CountDeadNum(num uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentDeathNum = num
}
func (u *User) CountAssistNum() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentAssistNum++
}
func (u *User) ResetAssistNum() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentAssistNum = 0
}
func (u *User) ResetKillNum() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.CurrentKillNum = 0
}

func (u *User) ResetDeadNum() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	(*u).CurrentDeathNum = 0
}

func (u *User) SetSignature(sig []byte) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	(*u).Signature = sig
}

func (u *User) SetAvatar(id uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Avatar = id
}

func (u *User) SetTitle(id uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.TitleId = id
}

func (u *User) UpdateCampaign(id uint16) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Campaign = u.Campaign | id
}

func (u *User) SetBuyMenu(menu UserBuyMenu) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Inventory.BuyMenu = menu
}

func GetNewUser() User {
	var mutex sync.Mutex
	return User{
		0,
		[]byte{},        //loginname
		[]byte{},        //username,looks can change it to another name
		[]byte{},        //passwd
		0,               //Gm
		1,               //level
		0,               //rank
		0,               //rankframe
		10000,           //points
		0,               //curEXP
		LevelExp[0],     //maxEXP
		0,               //playermatchs
		0,               //wins
		0,               //kills
		0,               //headshots
		0,               //deaths
		0,               //assists
		0,               // accuracy
		0,               // secondsPlayed
		NewNullString(), // netCafeName
		0,               // cash
		NewNullString(), // clanName
		0,               // clanMark
		0,               // worldRank
		0,               //campaign
		0,               // mpoints
		0,               // titleId
		[]uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // unlockedTitles
		NewNullString(), // signature
		0,               // bestGamemode
		0,               // bestMap
		[]uint8{0x00, 0x00, 0x18, 0x08, 0x00, 0x00, 0x00, 0x00, 0x42, 0x02,
			0x18, 0xC0, 0x09, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0xC0, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0xC8, 0xB7, 0x08, 0x00, 0x00, 0x04, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // unlockedAchievements
		1006, // avatar
		[]uint8{0x00, 0x00, 0x18, 0x08, 0x00, 0x00, 0x00, 0x00, 0x42, 0x02,
			0x18, 0xC0, 0x09, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0xC0, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0xC8, 0xB7, 0x08, 0x00, 0x00, 0x04, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x00, 0x00}, // unlockedAvatars
		0,        //viplevel
		0,        //vipXp
		0,        //skillHumanCurXp
		0x19AC,   //skillHumanMaxXp
		0,        //skillHumanPoints
		0,        //skillZombieCurXp
		0x16F6,   //skillZombieMaxXp
		0,        //skillZombiePoints
		[]byte{}, //mail
		nil,      //connection
		1,        //serverid
		0,        //channelid
		0,        //roomid
		0,        //currentTeam
		0,        //currentstatus
		false,    //currentIsIngame
		nil,      //sequence
		0,
		0,
		0,
		UserNetInfo{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		CreateNewUserInventory(), //仓库
		map[uint32]uint32{},
		&mutex,
	}
}

func (u *User) PunishPoints() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	u.Points = u.Points - 1000
	if u.Points < 0 {
		u.Points = 0
	}
}

func (u *User) GetPoints(num uint64) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if math.MaxUint64-num < u.Points {
		u.Points = math.MaxUint64
	} else {
		u.Points += num
	}
}

func (u *User) GetExp(num uint64) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	for num > 0 {
		if u.MaxExp-u.CurrentExp <= num { //能升级
			num -= u.MaxExp - u.CurrentExp
			u.CurrentExp = 0
			u.Level++
			if u.Level >= MAXLEVEL {
				u.Level = MAXLEVEL
				num = 0
			} else {
				u.MaxExp = LevelExp[u.Level-1]
			}
		} else if u.Level < MAXLEVEL { //不能升级
			u.CurrentExp += num
			num = 0
		} else { //已经满级
			num = 0
		}
	}
}

func (u *User) GetKills(num uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if math.MaxUint32-u.Kills <= num { //能升级
		u.Kills = math.MaxUint32
	} else { //不能升级
		u.Kills += num
	}
}

func (u *User) GetDeathes(num uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if math.MaxUint32-u.Deaths <= num { //能升级
		u.Deaths = math.MaxUint32
	} else { //不能升级
		u.Deaths += num
	}
}

func (u *User) GetAssists(num uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if math.MaxUint32-u.Assists <= num { //能升级
		u.Assists = math.MaxUint32
	} else { //不能升级
		u.Assists += num
	}
}

func (u *User) AddMatches() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if u.PlayedMatches < math.MaxUint32 {
		u.PlayedMatches++
	}
}

func (u *User) AddWins() {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if u.Wins < math.MaxUint32 {
		u.Wins++
	}
}

func (u *User) CountWeaponKill(itemid uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	if _, ok := u.WeaponKills[itemid]; ok {
		if u.WeaponKills[itemid] < math.MaxUint32 {
			u.WeaponKills[itemid]++
		}
	} else {
		u.WeaponKills[itemid] = 1
	}
}

func (u *User) AddItem(itemid uint32) {
	if u == nil {
		return
	}
	u.UserMutex.Lock()
	defer u.UserMutex.Unlock()
	for k, v := range u.Inventory.Items {
		if v.Id == itemid {
			if u.Inventory.Items[k].Count < math.MaxUint16 {
				u.Inventory.Items[k].Count++
			}
			return
		}
	}
	u.Inventory.Items = append(u.Inventory.Items, UserInventoryItem{itemid, 1})
	if u.Inventory.NumOfItem < math.MaxUint16 {
		u.Inventory.NumOfItem++
	}
}
