using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UniFramework.Machine;
using UniFramework.Module;

using UnityGeneralFramework.HotUpdateLogic;

namespace UnityGeneralFramework.HotUpdateLogic {
    /// <summary>
    /// 流程更新完毕
    /// </summary>
    internal class FsmPatchDone : IStateNode {
        void IStateNode.OnCreate(StateMachine machine) {
        }
        void IStateNode.OnEnter() {
            PatchEventDefine.PatchStatesChange.SendEventMessage("热更新完毕！");
            PatchEventDefine.PatchDone.SendEventMessage("资源包下载完成！");
        }
        void IStateNode.OnUpdate() {
        }
        void IStateNode.OnExit() {
        }
    }
}