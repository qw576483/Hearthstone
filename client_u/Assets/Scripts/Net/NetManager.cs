using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net;
using System.Net.Sockets;
using System.Net.WebSockets;

using ThisGame.Framework;

namespace ThisGame {
    public class NetManager : MonoSingleton<NetManager> {
        [Header("��ϷApp")]
        public WebSocket GameSocket;
        [Header("����ս��")]
        public WebSocket BattleScoket;

        [Header("����url")]
        public string BattleUrl = "ws://192.168.1.206:6070";

        // Start is called before the first frame update
        void Start() {

        }

        // Update is called once per frame
        void Update() {

        }

        /// <summary>
        /// ���ӷ���
        /// </summary>

        public void ConnectRoom() {

        }

        /// <summary>
        /// �Ͽ�����
        /// </summary>
        public void DisconnectRoom() {

        }
    }
}

