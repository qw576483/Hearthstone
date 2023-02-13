using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;
using YooAsset;

using UnityGeneralFramework.Common;
using UnityGeneralFramework.HotUpdateLogic;
using System;

namespace HeartStone {
    public class GameManager : MonoSingleton<GameManager> {
        [SerializeField]
		private Card.Card _card;

        public override void OnInit() {
            base.OnInit();
            DontDestroyOnLoad(gameObject);
#if UNITY_EDITOR
            StartCoroutine( YooAssetProxy.InitPackage( () => {
                CfgManager.Instance.OnInit();
                _card.OnResetCard(1);
                return true;
            } ) );
#endif
        }

        private bool onLoadDown() {
            throw new NotImplementedException();
        }

        // Start is called before the first frame update
        void Start() {

        }

        // Update is called once per frame
        void Update() {

        }
    }
}
