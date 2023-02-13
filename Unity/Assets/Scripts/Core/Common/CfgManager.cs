using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;

using YooAsset;

using HeartStone.Config;
using UnityGeneralFramework.Common;
using UnityGeneralFramework.HotUpdateLogic;

namespace HeartStone { 
    public class CfgManager : MonoSingleton<CfgManager> {
        // ---------------------------------------------------------------
        public static Cfg.Tables GetTables() {
            if (Instance == null) {
                Debug.LogWarning("no tables");
                return null;
            }

            if (Instance._tables == null) {
                Debug.LogWarning("no tables");
                return null;
            }
            return Instance._tables;
        }
        // ---------------------------------------------------------------
        private Cfg.Tables _tables;

        public new void OnInit() {
            // ¼ÓÔØËùÓÐcfg
            _tables = new Cfg.Tables(fileName => {
                return JSON.Parse(              
                    YooAssetProxy.LoadRawFileFileTextSync(GameEnum.CfgPath + "/" + fileName + ".json")) ;
            } );
        }
    }
}

