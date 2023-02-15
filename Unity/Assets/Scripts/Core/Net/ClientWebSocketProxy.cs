using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using UnityEngine;

namespace HeartStone.Net {
    /// <summary>WebSocket����</summary>
    public class ClientWebSocketProxy : MonoBehaviour{
        protected Uri m_uri = null;

        [SerializeField]
        private WebSocketState ShowInEditor = WebSocketState.None;
        //����Json
        public string HeartBeat = "{""Hello"": [{""Name"": ""}]}";

        public WebSocketState CurrState {
            get {
                if(m_client == null)
                    return WebSocketState.None;

                return m_client.State;
            }
        }

        /// <summary>WebSocket�ͻ��˶���</summary>
        protected ClientWebSocket m_client = null;
        protected CancellationToken m_cToken;
        
        /// <summary>����WebSocket���ص���Ϣ����</summary>
        protected WebSocketReceiveResult m_websocketReceiveResult = null;
        
        /// <summary>byte���飬���ڽ���WebSocket���ص�����</summary>
        protected byte[] m_byteArrBuffer = null;
        
        /// <summary>����WebSocket���ص��ַ�������</summary>
        protected string m_result = null;

        /// <summary>�Ƿ�ѭ�������Ӵ��ڴ�״̬��</summary>
        protected bool Loop { get { return m_client.State == WebSocketState.Open; } }

        void Awake() {
            m_uri = new Uri(NetEnum.WebSocketUrl + ":" + NetEnum.WebSocketPort);
            m_client = new ClientWebSocket();
            m_cToken = new CancellationToken();
        }

        /// <summary>��ȡ��������byte�����</summary>
        /// <param name="arr">byte��������</param>
        /// <returns>���byte�����</returns>
        protected ArraySegment<byte> GetBuffer(byte[] arr) {
            return new ArraySegment<byte>(arr);
        }

        /// <summary>��ȡ��������byte�����</summary>
        /// <param name="str">�ַ�������</param>
        /// <returns>���byte�����</returns>
        protected ArraySegment<byte> GetBuffer(string str) {
            return GetBuffer(Encoding.UTF8.GetBytes(str));
        }

        /// <summary>������Ϣ</summary>
        /// <returns>����ֵΪWebSocketReceiveResult��Task</returns>
        protected async Task<WebSocketReceiveResult> ReceiveMessage() {
            m_byteArrBuffer = new byte[1024];
            WebSocketReceiveResult wsrResult = await m_client.ReceiveAsync(GetBuffer(m_byteArrBuffer), new CancellationToken() );//��������
            
            //Debug.Log(wsrResult.Count + "---" + wsrResult.EndOfMessage + "---" + wsrResult.MessageType);
            
            m_result += Encoding.UTF8.GetString(m_byteArrBuffer, 0, wsrResult.Count);
            return wsrResult;
        }
        
        /// <summary>�������</summary>
        protected virtual void ParseResult() {
            ShowInEditor = m_client.State;
            Debug.Log("���ݽ��:" + m_result);
        }

        /// <summary>���籨��</summary>
        /// <param name="ex">������Ϣ</param>
        protected virtual void WebSocketError(Exception ex) {
            ShowInEditor = m_client.State;
            Debug.LogError(ex.Message + "\n" + ex.StackTrace + "\n" + ex.Source + "\n" + ex.HelpLink);
        }

        /// <summary>���ӡ�����</summary>
        public async void ConnectAuthReceive() {
            try {
                await m_client.ConnectAsync(m_uri, m_cToken);//����

                ShowInEditor = m_client.State;

                while(Loop) {//����������Ϣ
                    m_websocketReceiveResult = await ReceiveMessage();

                    if(m_websocketReceiveResult.EndOfMessage) {//������һ��������Ϣ������
                                                                //Debug.Log("����һ����Ϣ��" + m_result);
                        if(string.IsNullOrEmpty(m_result)) {// 
                            break;
                        }
                        ParseResult();
                    }
                }
            } catch(Exception ex) {
                WebSocketError(ex);
            }
        }
        /// <summary>��������</summary>
        /// <param name="text">������Ϣ����</param>
        public async Task Send(string text) {
            if(m_client.State == WebSocketState.None) { Debug.Log("δ�������ӣ�"); return; }
            await m_client.SendAsync(GetBuffer(text), WebSocketMessageType.Text, true, m_cToken);//��������
        }

        /// <summary>�ر�</summary>
        public async void Close() {
            if(m_client.State == WebSocketState.None) { Debug.Log("δ�������ӣ�"); return; }
            await m_client.CloseAsync(WebSocketCloseStatus.NormalClosure, "����հ�", m_cToken);
        }
        
        /// <summary>��ֹ</summary>
        public void Abort() {
            if(m_client.State == WebSocketState.None) { Debug.Log("δ�������ӣ�"); return; }
            m_client.Abort();
        }

        /// <summary>
        /// ��������
        /// </summary>
        /// <returns></returns>
        public IEnumerator SendHeartBeatMsg() {
            yield return new WaitForSeconds(5.0f);
        }
     }
}