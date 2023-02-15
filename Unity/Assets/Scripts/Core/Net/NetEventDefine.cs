using System.Collections;
using System.Collections.Generic;
using UniFramework.Event;
using UnityEngine;

namespace HeartStone.Net {
    /// <summary>
    /// �����¼�
    /// </summary>
    public class NetEventDefine {
        /// <summary>
        /// ��������ʼ��ʧ��
        /// </summary>
        public class WebScoketEvent : IEventMessage {
            /// <summary>
            /// �ܶ��¼�����Ҫ��������Ϣ�����ϵ�һ���¼��������
            /// </summary>
            public enum EventName { Open, Close, }

            /// <summary>
            /// ���¼�����
            /// </summary>
            public EventName Name;
            public static void SendEventMessage(EventName name) {
                var msg = new WebScoketEvent();
                msg.Name = name;
                UniEvent.SendMessage(msg);
            }
        }

    }
}
