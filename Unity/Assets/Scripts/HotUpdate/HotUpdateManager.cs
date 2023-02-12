using System.Collections.Generic;
using UniFramework.Event;
using UnityEngine;

using UnityGeneralFramework.Common;

namespace UnityGeneralFramework.HotUpdateLogic {
    /// <summary>
    /// 热更新管理
    /// </summary>
    public class HotUpdateManager : MonoSingleton<HotUpdateManager> {

        public YooAssetProxy HUYooAsset;
        public HybridCLRProxy HUHybridCLR;

        private readonly EventGroup _eventGroup = new EventGroup();

        public override void OnInit() {
            HUYooAsset = transform.Find("YooAsset").GetComponent<YooAssetProxy>();
            HUHybridCLR = transform.Find("HybirdCLR").GetComponent<HybridCLRProxy>();

            _eventGroup.AddListener<PatchEventDefine.PatchDone>(o => {
                //加载代码
                HUHybridCLR.LoadCode();
            });

            _eventGroup.AddListener<PatchEventDefine.LoadCodeDone>(o => {
                //开始游戏，进入登录界面
                
            });
        }

    }

}