using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using UnityGeneralFramework.HotUpdateLogic;

using HeartStone.Config;

namespace HeartStone.Card {
    /// <summary>
    /// ��ƬƤ��
    /// </summary>
    public class CardSkinLoader : MonoBehaviour {
        //�����ƣ�����
        public Material hero, azs, ams, adly, asq, afs, alr, asm, ass, adz;
        //???
        public Material mzs, mms, mzl, mdly, msq, mfs, mlr, msm, mss, mdz;
        //ϡ�ж�
        public Material common, rare, epic, legendary;

        private Renderer front, frontkuan, quality, icon;
        private UILabel labName, labCost, labDesc, labAtk, labHP;

        void Awake() {
            //3dModel
            icon = transform.Find("icon").GetComponent<Renderer>();
            front = transform.Find("front").GetComponent<Renderer>();
            frontkuan = transform.Find("frontkuan").GetComponent<Renderer>();
            quality = transform.Find("quality").GetComponent<Renderer>();
            //lab
            labName = transform.Find("name").GetComponent<UILabel>();
            labDesc = transform.Find("description").Find("description").GetComponent<UILabel>();
            labAtk = transform.Find("attack").GetComponent<UILabel>();
            labHP = transform.Find("health").GetComponent<UILabel>();
        }

        public void OnLoadSkin(int ID) {
            Cfg.Cards.Card cfgCards = CfgManager.GetTables().TbCard.Get(ID);
            Cfg.Cards.CardAttrs cfgCardAttrs = CfgManager.GetTables().TbCardAttrs.Get(ID);

            Material material = new Material(Shader.Find("Diffuse"));
            material.mainTexture = YooAssetProxy.LoadAssetSync<Texture>(GamePathConfig.CardTexture + cfgCards.ImageName);
            material.SetTextureScale("_MainTex", new Vector2(0.6f, 0.4f));
            material.SetTextureOffset("_MainTex", new Vector2(0.2f, 0.43f));

            icon.material = material;

            if(cfgCards.Cnname == "") {
                labName.text = cfgCards.Name;
            } else {
                labName.text = cfgCards.Cnname;
            }
            labCost.text = cfgCardAttrs.Cost.ToString();


            if(cfgCards.Cndescription == "") {
                labDesc.text = cfgCards.Description;
            } else {
                labDesc.text = cfgCards.Cndescription;
            }
            if(cfgCards.Type != Cfg.Cards.ECardType.HEROSKILL)//������Ǽ��ܿ�
            {
                labAtk.text = cfgCardAttrs.Attack.ToString();
                labHP.text = cfgCardAttrs.Health.ToString();
            }

            LoadClassSkin(cfgCards);
            LoadQualitySkin(cfgCards);
            LoadRaceSkin(cfgCards);
        }

        /// <summary>
        /// ְҵ
        /// </summary>
        /// <param name="cfgCard"></param>
        void LoadClassSkin(Cfg.Cards.Card cfgCard) {
            //�����Ӣ��
            if(cfgCard.Type == Cfg.Cards.ECardType.HERO) {
                front.material = hero;
                return;
            }

            //ֻ�м���  ���˿���Ƥ��
            if(cfgCard.Type != Cfg.Cards.ECardType.ATKMAGIC 
            && cfgCard.Type != Cfg.Cards.ECardType.VALET) {
                return;
            }

            Material zs, ms, zl, dly, sq, fs, lr, sm, ss, dz;

            if(cfgCard.Type == Cfg.Cards.ECardType.VALET) {
                zs = mzs;
                ms = mms;
                zl = mzl;
                dly = mdly;
                sq = msq;
                fs = mfs;
                lr = mlr;
                sm = msm;
                ss = mss;
                dz = mdz;
            } else {
                zl = adz;
                zs = azs;
                ms = ams;
                dly = adly;
                sq = asq;
                fs = afs;
                lr = alr;
                sm = asm;
                ss = ass;
                dz = adz;
            }

            switch(cfgCard.Classs) {
                case Cfg.Cards.ECardClass.ANY:
                    front.material = zl;
                    break;

                case Cfg.Cards.ECardClass.DRUID:
                    front.material = dly;
                    break;

                case Cfg.Cards.ECardClass.HUNTER:
                    front.material = lr;
                    break;

                case Cfg.Cards.ECardClass.MAGE:
                    front.material = fs;
                    break;

                case Cfg.Cards.ECardClass.PALADIN:
                    front.material = sq;
                    break;

                case Cfg.Cards.ECardClass.PRIEST:
                    front.material = ms;
                    break;

                case Cfg.Cards.ECardClass.ROGUE:
                    front.material = dz;
                    break;

                case Cfg.Cards.ECardClass.SHAMA:
                    front.material = sm;
                    break;

                case Cfg.Cards.ECardClass.WARLOCK:
                    front.material = ss;
                    break;

                case Cfg.Cards.ECardClass.WARRIOR:
                    front.material = zs;
                    break;
            }

            frontkuan = front;
        }

        /// <summary>
        /// Ʒ��Ƥ��
        /// </summary>
        void LoadQualitySkin(Cfg.Cards.Card cfgCard) {

            if(cfgCard.Set == Cfg.Cards.ECardSet.BASIC)//��������û��Ʒ��ģ�͵�
            {
                quality.gameObject.SetActive(false);
            }
            switch(cfgCard.Quality) {
                case Cfg.Cards.ECardQuality.FREE:
                    quality.gameObject.SetActive(false);
                    break;

                case Cfg.Cards.ECardQuality.Common:
                    quality.material = common;
                    break;

                case Cfg.Cards.ECardQuality.EPIC:
                    quality.material = epic;
                    break;

                case Cfg.Cards.ECardQuality.LENGENDARY:
                    quality.material = legendary;
                    break;

                case Cfg.Cards.ECardQuality.RARE:
                    quality.material = rare;
                    break;
            }
        }
        /// <summary>
        /// ����Ƥ��
        /// </summary>
        void LoadRaceSkin(Cfg.Cards.Card cfgCard) {
            //Ӣ��
            if(cfgCard.Type != Cfg.Cards.ECardType.VALET 
            && cfgCard.Type != Cfg.Cards.ECardType.HERO)//�����ʹ�ûrace
            {
                return;
            }
            UILabel racetext = transform.Find("race").GetComponent<UILabel>();

            switch(cfgCard.Race) {
                case Cfg.Cards.ECardRace.ANY:
                    racetext.text = "";
                    transform.Find("racedi").gameObject.SetActive(false);
                    break;

                case Cfg.Cards.ECardRace.BEAST:
                    racetext.text = "Ұ��";
                    break;

                case Cfg.Cards.ECardRace.DEMON:
                    racetext.text = "��ħ";
                    break;

                case Cfg.Cards.ECardRace.DRAGON:
                    racetext.text = "����";
                    break;

                case Cfg.Cards.ECardRace.MURLOC:
                    racetext.text = "����";
                    break;

                case Cfg.Cards.ECardRace.PIRATE:
                    racetext.text = "����";
                    break;

                case Cfg.Cards.ECardRace.TOTEM:
                    racetext.text = "ͼ��";
                    break;
            }
        }
    }
}

