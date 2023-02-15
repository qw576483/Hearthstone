using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

namespace HeartStone.Net {
    /// <summary>
    /// �������ݸ�ʽ
    /// ǰ׺S����ʾSend
    /// ǰ׺R����ʾReceive
    /// SR�㶮��
    /// </summary>
    public static class NetData {
        /// <summary>
        /// ͨ�ý�������
        /// </summary>
        public class ParaseData { 
            public string Name { get; set; }
            public string Data { get; set; }
        }
    }


    /// <summary>
    /// ֱ������һЩԤ��õ�����
    /// </summary>
    public static class NetDataPre {
        /// <summary>
        /// ����
        /// </summary>
        public static Dictionary<string, Dictionary<string, string>> SRHeartBeatPre = 
            new Dictionary<string, Dictionary<string, string>>() {
                {"Hello", new Dictionary<string, string>(){ {"Name", ""} } }
            };
        
        public static string sHeartBeat = JsonConvert.SerializeObject(SRHeartBeatPre);
    
    }
}