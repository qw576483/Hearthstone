using System.Collections;
using System.Collections.Generic;
using System.Net.WebSockets;
using UnityEngine;

using UnityGeneralFramework.Common;

namespace HeartStone.Net {
    /// <summary>
    /// NetManager
    /// 全局唯一管理
    /// </summary>
    public class NetManager : MonoSingleton<NetManager> {
        public ClientWebSocketProxy battleConn;

        public override void AwakeSingleton() {
            base.AwakeSingleton();
            GameManager.Instance.netManager = this;
        }
        

        public void Connect2Server() {
            if(battleConn) {
                if(battleConn.CurrState == WebSocketState.Closed) {
                    Destroy(battleConn.gameObject);
                } else {
                    Debug.LogWarning("WebSocket is exist!!!");
                    return;
                }
            }

            GameObject obj = new GameObject();
            obj.transform.parent = transform;
            obj.transform.name = "GameWebSocket";

            battleConn = obj.AddComponent<ClientWebSocketProxy>();
            battleConn.ConnectAuthReceive();
        }
    }
}