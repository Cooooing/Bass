package catalog

import (
	bbsmodule "bbs/module"
	commonmodule "common/pkg/module"
	contentmodule "content/module"
	economymodule "economy/module"
	gametownmodule "game_town/module"
	immodule "im/module"
	notifymodule "notify/module"
	platformmodule "platform/module"
	pushhubmodule "push_hub/module"
	pushnodemodule "push_node/module"
	schedulermodule "scheduler/module"
	usermodule "user/module"
)

// Descriptors 声明单体当前装配的业务模块。
var Descriptors = []commonmodule.Descriptor{
	Named("bbs", bbsmodule.Descriptor()),
	Named("content", contentmodule.Descriptor()),
	Named("economy", economymodule.Descriptor()),
	Named("game_town", gametownmodule.Descriptor()),
	Named("im", immodule.Descriptor()),
	Named("notify", notifymodule.Descriptor()),
	Named("platform", platformmodule.Descriptor()),
	Named("push_hub", pushhubmodule.Descriptor()),
	Named("push_node", pushnodemodule.Descriptor()),
	Named("scheduler", schedulermodule.Descriptor()),
	Named("user", usermodule.Descriptor()),
}

// Named 用配置模块名覆盖模块描述中的默认名称。
func Named(name string, descriptor commonmodule.Descriptor) commonmodule.Descriptor {
	descriptor.Name = name
	return descriptor
}
