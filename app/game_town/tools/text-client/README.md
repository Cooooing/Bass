# Game Town Text Client

该客户端用于本地开发和联调，支持两种交互模式：

- `tui`：在真实终端中使用 Bubble Tea 界面。
- `console`：在 IDEA Run/Services 等非 TTY 控制台中逐行读取命令。
- `auto`：默认模式，自动根据标准输入是否为 TTY 选择界面。

## 启动

直连 gRPC：

```bash
go run ./tools/text-client -addr 127.0.0.1:9000
```

在 IDEA 服务窗口中可显式指定逐行模式：

```bash
go run ./tools/text-client -mode console -addr 127.0.0.1:9000
```

通过 Consul 发现服务：

```bash
go run ./tools/text-client -consul-addr 127.0.0.1:8500
```

启动后输入 `/help` 查看命令。Console 模式需要输入完整命令并按回车。

## 代码结构

- `main.go`：启动参数和交互模式选择。
- `client.go`：gRPC 直连与 Consul 服务发现。
- `tui.go`：Bubble Tea 界面。
- `console.go`：非 TTY 逐行控制台。
- `watcher.go`：Event Watch 断线重连和 sequence 续传。
- `command_*.go`：按业务领域拆分的命令处理。

## 基本流程

使用 /config list、/world list 查看包含 ID 的资源列表；客户端重启后可通过 /player use <player_id> 和 /world use <world_id> 恢复上下文。

/look 会展示全部场景、场景编码、NPC 数量及全部 NPC 所在位置。移动使用 /move <location_code>，只有与玩家处于同一场景的 NPC 才能 /talk。

~~~
/config add local-qwen3-1.7b ollama http://192.168.100.10:31434 qwen3:1.7b
/config list
/register town_player_20260720_01
/status
/world create 1 在九州灵脉逐渐枯竭的仙侠世界中，维系天地秩序的上古天门突然崩裂，无数秘境碎片坠入人间。各大宗门为争夺残存灵脉彼此征伐，妖族趁机越过边境，沉寂千年的魔道也开始复苏。玩家是一名身世不明的散修，体内封印着一缕不属于这个时代的剑魂。随着修行深入，玩家将结识宗门剑修、游历医仙、妖族圣女和神秘卦师，在仙城、古战场、云海剑宗和幽冥裂谷之间探索，寻找天门崩裂的真相，并决定重建天道、开辟新的修行纪元，还是让旧有秩序彻底覆灭。

/status
/world list
/look
/talk <当前场景npc_id> 你好，你知道最近的失踪事件吗？
/talk <当前场景npc_id> 这些神秘符文以前出现过吗？
/act 我仔细检查当前场景是否存在失踪者留下的痕迹
/events
/move <目标location_code>
/look
/talk <新场景npc_id> 我想了解这里发生过什么
/act 我观察新场景中的人物和异常线索
/events
/quit
~~~