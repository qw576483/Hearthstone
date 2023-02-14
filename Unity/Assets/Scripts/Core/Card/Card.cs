using System.Collections;
using System.Collections.Generic;
using HeartStone.Config;
using HeartStone.Data;
using UnityEngine;
using UnityGeneralFramework.HotUpdateLogic;
using YooAsset;

namespace HeartStone.Card {
    public class Card : MonoBehaviour {
        //配表数据
        public Cfg.Cards.Card cfgCards;
        public Cfg.Cards.CardAttrs cfgCardAttrs;

        //卡片类型
        public Cfg.Cards.ECardType CardType { get { return cfgCards.Type; } }
        //种族
        public Cfg.Cards.ECardRace Race { get { return cfgCards.Race; } }
        //职业
        public Cfg.Cards.ECardClass Class { get { return cfgCards.Classs; } }
        //品质
        public Cfg.Cards.ECardQuality Quality { get { return cfgCards.Quality; } }
        //套牌类型
        public Cfg.Cards.ECardSet Set { get { return cfgCards.Set; } }

        //当前数据
        public CardData cardData;
        //ID
        public int ID;

        void Awake() {
        }

        /// <summary>
        /// 重置Card
        /// </summary>
        /// <param name="ID"></param>
        public void OnResetCard(int id) {
            this.ID = id;

            cfgCards = CfgManager.GetTables().TbCard.Get(id);
            cfgCardAttrs = CfgManager.GetTables().TbCardAttrs.Get(id);

            //根据数据加载皮肤
            GetComponent<CardSkinLoader>().OnLoadSkin(id);
        }
    }
}
