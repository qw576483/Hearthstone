package iface

type ICHCommon interface {
	GetId() int

	GetFather() ICHCommon           // 获得父级
	GetFatherCard() ICard           // 获得父卡牌
	GetFatherHero() IHero           // 获得父英雄
	SetFather(interface{})          // 设置父卡牌
	GetSubCards() []ICard           // 获得子卡牌
	SetSubCards([]ICard)            // 设置子卡牌
	AddSubCards(ICard, interface{}) // 添加子卡牌
	RemoveSubCards(ICard)           // 删除子卡牌
}
