using System.Collections;
using System.Collections.Generic;
using HeartStone.Config;
using HeartStone.Data;
using UnityEngine;
using UnityGeneralFramework.HotUpdateLogic;
using YooAsset;

namespace HeartStone.Card {
    public class Card : MonoBehaviour {
        //�������
        public Cfg.Cards.Card cfgCards;
        public Cfg.Cards.CardAttrs cfgCardAttrs;

        //��Ƭ����
        public Cfg.Cards.ECardType CardType { get { return cfgCards.Type; } }
        //����
        public Cfg.Cards.ECardRace Race { get { return cfgCards.Race; } }
        //ְҵ
        public Cfg.Cards.ECardClass Class { get { return cfgCards.Classs; } }
        //Ʒ��
        public Cfg.Cards.ECardQuality Quality { get { return cfgCards.Quality; } }
        //��������
        public Cfg.Cards.ECardSet Set { get { return cfgCards.Set; } }

        //��ǰ����
        public CardData cardData;
        //ID
        public int ID;

        void Awake() {
        }

        /// <summary>
        /// ����Card
        /// </summary>
        /// <param name="ID"></param>
        public void OnResetCard(int id) {
            this.ID = id;

            cfgCards = CfgManager.GetTables().TbCard.Get(id);
            cfgCardAttrs = CfgManager.GetTables().TbCardAttrs.Get(id);

            //�������ݼ���Ƥ��
            GetComponent<CardSkinLoader>().OnLoadSkin(id);
        }
    }
}
