using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

namespace HeartStone.Net {
    /// <summary>
    /// 传输数据格式
    /// 前缀S，表示Send
    /// 前缀R，表示Receive
    /// SR你懂的
    /// </summary>
    public static class NetData {
        /// <summary>
        /// 通用解析数据
        /// </summary>
        public class ParaseData { 
            public string Name { get; set; }
            public string Data { get; set; }
        }
    }


    /// <summary>
    /// 直接设置一些预设好的数据
    /// </summary>
    public static class NetDataPre {
        /// <summary>
        /// 心跳
        /// </summary>
        public static Dictionary<string, Dictionary<string, string>> SRHeartBeatPre = 
            new Dictionary<string, Dictionary<string, string>>() {
                {"Hello", new Dictionary<string, string>(){ {"Name", ""} } }
            };
        
        public static string sHeartBeat = JsonConvert.SerializeObject(SRHeartBeatPre);
    
    }
}