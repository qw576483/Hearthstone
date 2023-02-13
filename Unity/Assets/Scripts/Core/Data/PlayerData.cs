using HeartStone.Data;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Player {
    /// <summary>
    /// Player Data
    /// </summary>
    [SerializeField]
    public class PlayerData {
        //ID
        public long ID = -1;
        //昵称
        public string Name = "???";
    }

    /// <summary>
    /// 玩家战斗数据
    /// </summary>
    [SerializeField]
    public class PlayerBattleData {
        //职业
        public Cfg.Cards.ECardClass cls = Cfg.Cards.ECardClass.UNKNOWN;
        //牌库
        public List<CardData> cardDatas;
        //kv牌库
        public Dictionary<int, CardData> cardDataMap;
    }
}

