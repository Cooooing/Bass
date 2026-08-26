package enum

import "common/proto/gen/common/enums"

// LoginRealm 表示认证会话所属入口域
type LoginRealm string

const (
	LoginRealmBBS      LoginRealm = "bbs"
	LoginRealmBBSAdmin LoginRealm = "bbs_admin"
	LoginRealmGameIdle LoginRealm = "game_idle"
)

// LoginRealmMap 将内部登录域映射到 proto 枚举
var LoginRealmMap = NewMapping[LoginRealm, enums.LoginRealm](map[LoginRealm]Entry[LoginRealm, enums.LoginRealm]{
	LoginRealmBBS:      {Proto: enums.LoginRealm_LOGIN_REALM_BBS},
	LoginRealmBBSAdmin: {Proto: enums.LoginRealm_LOGIN_REALM_BBS_ADMIN},
	LoginRealmGameIdle: {Proto: enums.LoginRealm_LOGIN_REALM_GAME_IDLE},
})

func (e LoginRealm) String() string {
	return string(e)
}
