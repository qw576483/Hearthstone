# hs
炉石传说 , golang实现

Hearthstone by golang

## 使用

server:`go run main.go`

client:打开2个`client/client.html`

## 概览

![](./example/overview/1.png)

![](./example/overview/2.png)

## 已实现

### 流程

- 选择预选卡 cpre
- 出卡 r
- 卡牌攻击 a
- 玩家攻击 ha
- 玩家技能 s
- 回合结束 e

### 特质示例

- 冲锋(id:1 石牙野猪)
- 战吼(id:3 寒光智者)
- 亡语(id:4 麦田傀儡)
- 突袭(id:8 螃蟹骑士)
- 风怒(id:8 螃蟹骑士)
- 连击(id:9 毁灭之刃)
- 无法攻击(id:11 上古看守者)
- 嘲讽(id:12 持盾卫士)
- 圣盾(id:13 银色侍从)
- 潜行(id:14 耐心的刺客)
- 剧毒(id:14 耐心的刺客)
- 抉择(id:32 丛林守护者)
- 免疫(id:35 狂野怒火)

### 机制示例

#### 卡牌类型
- 随从卡(id:3 寒光智者)
- 英雄技能(id:26 盗贼基础技能)
- 武器卡(id:9 毁灭之刃)
- 法术卡(id:29 奉献) , 奥秘(id:34 忏悔)
- buff卡(id:21 我的回合结束驱散)

#### 固有事件
- 初始化时(id:21 buff 我的回合结束驱散)
- 释放时(id:0 幸运币)
- 步入战场时(id:6 工程车)
- 离开战场时(id:6 工程车)
- 死亡时(id:4 麦田傀儡)
- 受到伤害后(id:37 古拉巴什狂暴者)
- 生命值改变后(id:38 阿曼尼狂战士)
- 获得自己的费用时(id:38 熔核巨人)

#### 注册事件
- 回合开始时(id:6 工程车)
- 回合结束时(id:7 铸剑师)
- 其他卡牌死亡时(id:10 食腐土狼)
- 其他卡牌步入战场时(id:16 飞刀杂耍者)
- 其他卡牌获取自己的攻击力时(id:17 火舌图腾)
- 其他卡牌获取自己的费用时(id:18 小个子召唤师)
- 其他卡牌获取自己的生命值时(id:19 暴风城勇士)
- 其他卡牌释放前(id:31 游学者周卓)
- 其他卡牌释放后(id:34 忏悔)
- 其他卡牌攻击前(id:40 冰冻陷阱)

#### 选取目标
- 随机场上所有单位(id:15 疯狂投弹者)
- 所有敌人(id:29 奉献)
- 随机敌人(id:6 工程车)
- 随机另一个友方随从(id:7 铸剑师)
- 随机随从(id:44 希尔瓦娜斯·风行者)
 
#### 攻和血
- 伤血对卡牌(id:6 工程车)
- 伤血对英雄(id:6 工程车)
- 固化属性,交换伤害和血量(id:2 疯狂的炼金师)

#### 法力水晶
- 本回合获得水晶(id:0 幸运币)
- 过载(id:36 闪电箭)

#### 特质
- 增加特质(id:25 银色保卫者)

#### BUFF(攻击力,生命值,特质,法术伤害)

`如果没有回合限制或者附加机制,实际上没有必要加buff`

- 我的回合结束驱散(id:21)
- 我的回合开始驱散(id:22)
- 永久生效(id:23)
- 我的回合结束时消散和消灭宿主（挂载英雄上就会消灭英雄！）(id:64)
- 我的回合开始时消散和消灭宿主（挂载英雄上就会消灭英雄！）(id:65)
- 添加buff,本回合内增加攻击(id:24 叫嚣的中士)
- 添加buff,本回合内增加特质(id:35 狂野怒火)
- 添加buff,永久增加攻击力(id:38 阿曼尼狂战士)

#### 操作卡牌
- 抽牌(id:3 寒光智者)
- 复制卡牌(id:31 游学者周卓)
- 卡牌置入手牌(id:31 游学者周卓)
- 召唤新随从到战场(id:4 麦田傀儡)
- 从战场中回到手牌(id:33 年轻的酒仙)
- 变形战场上的卡牌(id:43 工匠大师欧沃斯巴克)
- 夺取对手战场卡牌(id:44 希尔瓦娜斯·风行者)
- 丢弃卡牌(id:61 死亡之翼)

#### 特殊机制
- 沉默(id:28 铁喙猫头鹰)
- 增加法术伤害,实际上配置config就行(id:30 狗头人地卜师)
- 抉择两种或多种不同效果(id:32 丛林守护者)

## todo
- 实现经典卡池全部机制
- 实现全部经典卡牌
- 实现笨蛋ai
- 整理代码

## 卡牌实现

定义
- `logic/config/card.go`

实现
- `logic/cards/card1-100.go`
- `logic/cards/card_point.go`
