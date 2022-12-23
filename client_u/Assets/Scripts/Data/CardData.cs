using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace ThisGame.Battle.Card {
    /// <summary>
    /// 卡片数据
    /// </summary>
    public class CardData {
        [Header("生命值")]
        public int HP = 0;
        [Header("最大生命值")]
        public int HPMax = 0;
        [Header("攻击力")]
        public int Atk = 0;
        [Header("攻击力最大值")]
        public int AtkMax = 0;
    }
}

