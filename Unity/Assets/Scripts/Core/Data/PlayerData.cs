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
        //�ǳ�
        public string Name = "???";
    }

    /// <summary>
    /// ���ս������
    /// </summary>
    [SerializeField]
    public class PlayerBattleData {
        //ְҵ
        public Cfg.Cards.ECardClass cls = Cfg.Cards.ECardClass.UNKNOWN;
        //�ƿ�
        public List<CardData> cardDatas;
        //kv�ƿ�
        public Dictionary<int, CardData> cardDataMap;
    }
}

