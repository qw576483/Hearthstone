using System.Collections.Generic;
using UniFramework.Event;
using UnityEngine;

using UnityGeneralFramework.Common;

namespace UnityGeneralFramework.HotUpdateLogic {
    /// <summary>
    /// �ȸ��¹���
    /// </summary>
    public class HotUpdateManager : MonoSingleton<HotUpdateManager> {

        public YooAssetProxy HUYooAsset;
        public HybridCLRProxy HUHybridCLR;

        private readonly EventGroup _eventGroup = new EventGroup();

        public override void OnInit() {
            HUYooAsset = transform.Find("YooAsset").GetComponent<YooAssetProxy>();
            HUHybridCLR = transform.Find("HybirdCLR").GetComponent<HybridCLRProxy>();

            _eventGroup.AddListener<PatchEventDefine.PatchDone>(o => {
                //���ش���
                HUHybridCLR.LoadCode();
            });

            _eventGroup.AddListener<PatchEventDefine.LoadCodeDone>(o => {
                //��ʼ��Ϸ�������¼����
                
            });
        }

    }

}