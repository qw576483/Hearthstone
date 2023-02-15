using System.Collections;
using System.Collections.Generic;
using UniFramework.Event;
using UnityEngine;

namespace HeartStone.Net {
    /// <summary>
    /// 网络事件
    /// </summary>
    public class NetEventDefine {
        /// <summary>
        /// 补丁包初始化失败
        /// </summary>
        public class WebScoketEvent : IEventMessage {
            /// <summary>
            /// 很多事件不需要带其他信息，集合到一个事件里就行了
            /// </summary>
            public enum EventName { Open, Close, }

            /// <summary>
            /// 子事件名称
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
