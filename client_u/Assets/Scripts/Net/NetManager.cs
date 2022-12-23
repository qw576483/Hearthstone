using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net;
using System.Net.Sockets;
using System.Net.WebSockets;

using ThisGame.Framework;

namespace ThisGame {
    public class NetManager : MonoSingleton<NetManager> {
        [Header("游戏App")]
        public WebSocket GameSocket;
        [Header("房间战斗")]
        public WebSocket BattleScoket;

        [Header("房间url")]
        public string BattleUrl = "ws://192.168.1.206:6070";

        // Start is called before the first frame update
        void Start() {

        }

        // Update is called once per frame
        void Update() {

        }

        /// <summary>
        /// 连接房间
        /// </summary>

        public void ConnectRoom() {

        }

        /// <summary>
        /// 断开房间
        /// </summary>
        public void DisconnectRoom() {

        }
    }
}

