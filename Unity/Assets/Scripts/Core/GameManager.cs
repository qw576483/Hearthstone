using System;
using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;
using YooAsset;

using UnityGeneralFramework.Common;
using UnityGeneralFramework.HotUpdateLogic;
using HeartStone.Net;

namespace HeartStone {
    public class GameManager : MonoSingleton<GameManager> {
        public CfgManager cfgManager;
        public NetManager netManager;

        public override void AwakeSingleton() {
            base.AwakeSingleton();
            DontDestroyOnLoad(gameObject);
        }

        void Start() {
#if UNITY_EDITOR
            StartCoroutine(YooAssetProxy.InitPackage(() => {
                StartCoroutine( OnLoadModule() );
                return true;
            }));
#else
            StartCoroutine( OnLoadModule() );
#endif
        }

        public IEnumerator OnLoadModule() {
            yield return new WaitForSeconds(1.0f);

            CfgManager.Instance.OnInit();
        }
    }
}
