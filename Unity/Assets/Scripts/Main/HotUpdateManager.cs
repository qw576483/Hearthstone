using UnityEngine;
using YooAsset;

namespace UnityGeneralFramework.HotUpdateLogic {
    /// <summary>
    /// 热更新管理
    /// </summary>
    public class HotUpdateManager : MonoSingleton<HotUpdateManager> {

        public YooAssetProxy HUYooAsset;
        public HybridCLRProxy HUHybridCLR;

        public override void OnInit() {
            HUYooAsset = transform.Find("YooAsset").GetComponent<YooAssetProxy>();
            HUHybridCLR = transform.Find("HybirdCLR").GetComponent<HybridCLRProxy>();
        }

        public void OnDownLoadDone() {
            HUHybridCLR.StartGame();
            YooAssets.LoadSceneAsync("login");
        }
    }

}