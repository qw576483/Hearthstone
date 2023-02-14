using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using Newtonsoft.Json;

namespace UnityGeneralFramework.HotUpdateLogic {
   /// <summary>
   /// 热更json文件解析
   /// </summary>
    public class HUConfigJson {
        public string default_package_name;
        public string game_version;
        public string hot_update_version;
    }

    /// <summary>
    /// 热更配置都在这
    /// </summary>
    public class HUConfig : Common.MonoSingleton<HUConfig> {
        //配置
        public string windowsUrl1 = "http://127.0.0.1";
        public string windowsUrl2 = "http://127.0.0.1";

        public string androidUrl1 = "";
        public string androidUrl2 = "";

        public string iosUrl1 = "";
        public string iosUrl2 = "";

        public string URL1 = "";
        public string URL2 = "";
        public int port = 0;
        public string gameVersion = "v1.0";
        public string defPackageName = "";

        public override void AwakeSingleton() {
            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            HUConfigJson config = JsonConvert.DeserializeObject<HUConfigJson>(txt.text);

            defPackageName = config.default_package_name;
            gameVersion = config.game_version;

            //editor windows
#if UNITY_EDITOR || UNITY_WINDOW
            URL1 = windowsUrl1;
            URL2 = windowsUrl2;

            //android
#elif UNITY_ANDROID
        URL1 = androidUrl1;
        URL2 = androidUrl2;

        //iOS
#elif UNITY_IOS
        
        URL1 = iosUrl1;
        URL2 = iosUrl2;
#else
        URL1 = windowsUrl1;
        URL2 = windowsUrl2;
#endif
        }

        #region HybridCLR
        //AOT 补充元数据 dll      
        public static List<string> AOTMetaAssemblyNames { get; } = new List<string>() {
            "mscorlib.dll",
            "System.dll",
            "System.Core.dll",
        };
        #endregion
    }

}