using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using UnityGeneralFramework.Common;

namespace HeartStone.Net {
    /// <summary>
    /// NetManager
    /// ȫ��Ψһ����
    /// </summary>
    public class NetManager : MonoSingleton<NetManager> {

        public override void AwakeSingleton() {
            base.AwakeSingleton();
            GameManager.Instance.netManager = this;
        }

    }
}