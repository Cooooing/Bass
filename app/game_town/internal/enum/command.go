package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1"
)

// CommandType 是文字命令类型的内部持久化枚举。
type CommandType string

const (
	CommandTypeHelp              CommandType = "help"
	CommandTypeRegister          CommandType = "register"
	CommandTypeWorldCreate       CommandType = "world_create"
	CommandTypeWorldJoin         CommandType = "world_join"
	CommandTypeWorldList         CommandType = "world_list"
	CommandTypeLook              CommandType = "look"
	CommandTypeMove              CommandType = "move"
	CommandTypeTalk              CommandType = "talk"
	CommandTypeDo                CommandType = "do"
	CommandTypeStatus            CommandType = "status"
	CommandTypeEvents            CommandType = "events"
	CommandTypeNpcs              CommandType = "npcs"
	CommandTypeMemory            CommandType = "memory"
	CommandTypeTick              CommandType = "tick"
	CommandTypeAgentConfigCreate CommandType = "agent_config_create"
	CommandTypeAgentConfigList   CommandType = "agent_config_list"
)

// CommandTypeMap 维护内部持久化值与 proto 枚举之间的映射。
var CommandTypeMap = commonenum.NewMapping[CommandType, v1.GameTownCommandType](map[CommandType]commonenum.Entry[CommandType, v1.GameTownCommandType]{
	CommandTypeHelp:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_HELP},
	CommandTypeRegister:          {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_REGISTER},
	CommandTypeWorldCreate:       {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_WORLD_CREATE},
	CommandTypeWorldJoin:         {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_WORLD_JOIN},
	CommandTypeWorldList:         {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_WORLD_LIST},
	CommandTypeLook:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_LOOK},
	CommandTypeMove:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_MOVE},
	CommandTypeTalk:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_TALK},
	CommandTypeDo:                {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_DO},
	CommandTypeStatus:            {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_STATUS},
	CommandTypeEvents:            {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_EVENTS},
	CommandTypeNpcs:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_NPCS},
	CommandTypeMemory:            {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_MEMORY},
	CommandTypeTick:              {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_TICK},
	CommandTypeAgentConfigCreate: {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_AGENT_CONFIG_CREATE},
	CommandTypeAgentConfigList:   {Proto: v1.GameTownCommandType_GAME_TOWN_COMMAND_TYPE_AGENT_CONFIG_LIST},
})

// CommandStatus 是命令处理状态的内部持久化枚举。
type CommandStatus string

const (
	CommandStatusReceived  CommandStatus = "received"
	CommandStatusSucceeded CommandStatus = "succeeded"
	CommandStatusFailed    CommandStatus = "failed"
)

// CommandStatusMap 维护内部持久化值与 proto 枚举之间的映射。
var CommandStatusMap = commonenum.NewMapping[CommandStatus, v1.GameTownCommandStatus](map[CommandStatus]commonenum.Entry[CommandStatus, v1.GameTownCommandStatus]{
	CommandStatusReceived:  {Proto: v1.GameTownCommandStatus_GAME_TOWN_COMMAND_STATUS_RECEIVED},
	CommandStatusSucceeded: {Proto: v1.GameTownCommandStatus_GAME_TOWN_COMMAND_STATUS_SUCCEEDED},
	CommandStatusFailed:    {Proto: v1.GameTownCommandStatus_GAME_TOWN_COMMAND_STATUS_FAILED},
})
