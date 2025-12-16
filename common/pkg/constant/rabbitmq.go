package constant

import "github.com/rabbitmq/amqp091-go"

// -- 类型别名定义

type ExchangeName string

func (e ExchangeName) String() string {
	return string(e)
}

type QueueName string

func (q QueueName) String() string {
	return string(q)
}

type QueueBindName string

func (q QueueBindName) String() string {
	return string(q)
}

type RoutingKey string

func (r RoutingKey) String() string {
	return string(r)
}

// -- 初始化参数结构体声明

type ExchangeDeclare struct {
	Name       ExchangeName  // 交换机名称
	Kind       string        // 类型: "direct", "fanout", "topic", "headers"
	Durable    bool          // 持久化：RabbitMQ 重启后保留
	AutoDelete bool          // 当没有绑定队列时是否自动删除
	Internal   bool          // 是否为内部使用（true 则不允许应用直接发送）
	NoWait     bool          // 是否不等待服务器确认
	Args       amqp091.Table // 额外参数（如备用交换机）
}

type QueueDeclare struct {
	Name       QueueName     // 队列名称
	Durable    bool          // 持久化
	AutoDelete bool          // 无消费者时删除
	Exclusive  bool          // 是否排他，仅创建队列的消费者才能访问
	NoWait     bool          // 是否不等待服务器确认
	Args       amqp091.Table // 额外参数
}

type QueueBind struct {
	Name     QueueName     // 队列名称
	Key      RoutingKey    // 路由键
	Exchange ExchangeName  // 交换机名称
	NoWait   bool          // 是否不等待服务器确认
	Args     amqp091.Table // 额外参数
}

// -- 枚举常量定义

// 交换机枚举
const (
	ExchangeUser    ExchangeName = "user.topic.exchange"
	ExchangeUserDlx ExchangeName = "user.dlx.exchange"

	ExchangeContent    ExchangeName = "content.topic.exchange"
	ExchangeContentDlx ExchangeName = "content.dlx.exchange"

	ExchangeNotify    ExchangeName = "notify.topic.exchange"
	ExchangeNotifyDlx ExchangeName = "notify.dlx.exchange"

	ExchangeEconomy    ExchangeName = "economy.topic.exchange"
	ExchangeEconomyDlx ExchangeName = "economy.dlx.exchange"
)

// 队列枚举
const (
	// 队列模块

	QueueUser    QueueName = "queue.user"
	QueueContent QueueName = "queue.content"
	QueueNotify  QueueName = "queue.notify"
	QueueEconomy QueueName = "queue.economy"

	// 死信队列

	QueueUserDlx    QueueName = "queue.user.dlx"
	QueueContentDlx QueueName = "queue.content.dlx"
	QueueNotifyDlx  QueueName = "queue.notify.dlx"
	QueueEconomyDlx QueueName = "queue.economy.dlx"
)

// 路由键枚举
const (
	// 用户模块

	RoutingKeyUser    RoutingKey = "user.#"
	RoutingKeyUserDlx RoutingKey = "user.dlx"

	RoutingKeyUserFollow   RoutingKey = "user.user.follow"
	RoutingKeyUserUnfollow RoutingKey = "user.user.unfollow"
	RoutingKeyUserBlock    RoutingKey = "user.user.block"
	RoutingKeyUserUnblock  RoutingKey = "user.user.unblock"

	// 内容模块

	RoutingKeyContent    RoutingKey = "content.#"
	RoutingKeyContentDlx RoutingKey = "content.dlx"

	RoutingKeyContentArticlePublish RoutingKey = "content.article.publish"
	RoutingKeyContentArticleThank   RoutingKey = "content.article.thank"
	RoutingKeyContentArticleLike    RoutingKey = "content.article.like"
	RoutingKeyContentArticleCollect RoutingKey = "content.article.collect"
	RoutingKeyContentArticleWatch   RoutingKey = "content.article.watch"
	RoutingKeyContentArticleAt      RoutingKey = "content.article.at"

	RoutingKeyContentCommentPublish       RoutingKey = "content.comment.publish"
	RoutingKeyContentCommentThank         RoutingKey = "content.comment.thank"
	RoutingKeyContentCommentLike          RoutingKey = "content.comment.like"
	RoutingKeyContentCommentCollect       RoutingKey = "content.comment.collect"
	RoutingKeyContentCommentAt            RoutingKey = "content.comment.at"
	RoutingKeyContentArticleVote          RoutingKey = "content.article.vote"
	RoutingKeyContentArticleLottery       RoutingKey = "content.article.lottery"
	RoutingKeyContentArticleLotteryWinner RoutingKey = "content.article.lottery_winner"

	// 通知模块

	RoutingKeyNotify    RoutingKey = "notify.#"
	RoutingKeyNotifyDlx RoutingKey = "notify.dlx"

	// 经济模块

	RoutingKeyEconomy    RoutingKey = "economy.#"
	RoutingKeyEconomyDlx RoutingKey = "economy.dlx"
)

// 配置映射表

// ExchangeMap 交换机配置
var ExchangeMap = map[ExchangeName]ExchangeDeclare{
	ExchangeUser:       {Name: ExchangeUser, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeUserDlx:    {Name: ExchangeUserDlx, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeContent:    {Name: ExchangeContent, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeContentDlx: {Name: ExchangeContentDlx, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeNotify:     {Name: ExchangeNotify, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeNotifyDlx:  {Name: ExchangeNotifyDlx, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeEconomy:    {Name: ExchangeEconomy, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
	ExchangeEconomyDlx: {Name: ExchangeEconomyDlx, Kind: "topic", Durable: true, AutoDelete: false, Internal: false, NoWait: false, Args: nil},
}

// QueueMap 队列配置
var QueueMap = map[QueueName]QueueDeclare{
	QueueUser:    {Name: QueueUser, Durable: true, Args: amqp091.Table{"x-dead-letter-exchange": ExchangeUserDlx.String(), "x-dead-letter-routing-key": RoutingKeyUserDlx.String()}},
	QueueContent: {Name: QueueContent, Durable: true, Args: amqp091.Table{"x-dead-letter-exchange": ExchangeContentDlx.String(), "x-dead-letter-routing-key": RoutingKeyContentDlx.String()}},
	QueueNotify:  {Name: QueueNotify, Durable: true, Args: amqp091.Table{"x-dead-letter-exchange": ExchangeNotifyDlx.String(), "x-dead-letter-routing-key": RoutingKeyNotifyDlx.String()}},
	QueueEconomy: {Name: QueueEconomy, Durable: true, Args: amqp091.Table{"x-dead-letter-exchange": ExchangeEconomyDlx.String(), "x-dead-letter-routing-key": RoutingKeyEconomyDlx.String()}},

	QueueUserDlx:    {Name: QueueUserDlx, Durable: true},
	QueueContentDlx: {Name: QueueContentDlx, Durable: true},
	QueueNotifyDlx:  {Name: QueueNotifyDlx, Durable: true},
	QueueEconomyDlx: {Name: QueueEconomyDlx, Durable: true},
}

// QueueBindMap 队列绑定配置
var QueueBindMap = map[QueueBindName]QueueBind{
	// User模块
	//QueueBindName("bind.queue.user.notify"): {GetName: QueueUser, Key: RoutingKey("user.#"), Exchange: ExchangeUser},

	QueueBindName("bind.dlx.queue.user->user"): {Name: QueueUserDlx, Key: RoutingKeyUserDlx, Exchange: ExchangeUserDlx},

	// Content模块
	//QueueBindName("bind.queue.content"): {GetName: QueueContent, Key: RoutingKey("content.#"), Exchange: ExchangeContent},

	QueueBindName("bind.dlx.queue.content->content"): {Name: QueueContentDlx, Key: RoutingKeyContentDlx, Exchange: ExchangeContentDlx},

	// Notify模块
	QueueBindName("bind.queue.user->notify"):    {Name: QueueNotify, Key: RoutingKeyUser, Exchange: ExchangeUser},
	QueueBindName("bind.queue.content->notify"): {Name: QueueNotify, Key: RoutingKeyContent, Exchange: ExchangeContent},

	QueueBindName("bind.dlx.queue.notify->notify"): {Name: QueueNotifyDlx, Key: RoutingKeyNotifyDlx, Exchange: ExchangeNotifyDlx},

	// Economy模块
	//QueueBindName("bind.queue.economy"): {GetName: QueueEconomy, Key: RoutingKey("economy.#"), Exchange: ExchangeEconomy},

	QueueBindName("bind.dlx.queue.economy->economy"): {Name: QueueEconomyDlx, Key: RoutingKeyEconomyDlx, Exchange: ExchangeEconomyDlx},
}
