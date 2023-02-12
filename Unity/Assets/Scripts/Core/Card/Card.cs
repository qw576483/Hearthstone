using System.Collections;
using System.Collections.Generic;
using ThisGame.Config;
using ThisGame.Data;
using UnityEngine;
using UnityGeneralFramework.HotUpdateLogic;
using YooAsset;

namespace ThisGame.Card {
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

        public MeshRenderer iconRender;

        void Awake() {
            iconRender = transform.Find("ImgIcon").GetComponent<MeshRenderer>();
        }

        /// <summary>
        /// ����Card
        /// </summary>
        /// <param name="ID"></param>
        public void OnResetCard(int id) {
            this.ID = id;

            cfgCards = CfgManager.GetTables().TbCard.Get(id);
            cfgCardAttrs = CfgManager.GetTables().TbCardAttrs.Get(id);

            Material material = new Material(Shader.Find("Diffuse"));
            material.mainTexture = YooAssetProxy.LoadAssetSync<Texture>(GamePathConfig.CardTexture + cfgCards.ImageName);
            material.SetTextureScale("_MainTex", new Vector2(0.6f, 0.4f));
            material.SetTextureOffset("_MainTex", new Vector2(0.2f, 0.43f));

            iconRender.material = material;
        }
    }
}
