using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEditor;
using UnityEngine;

using UnityGeneralFramework.HotUpdateLogic;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// Editor 配置
    /// </summary>
    public static class GameEditorConfig{
        public static string ExcelBatPath = Path.Combine(Application.dataPath, "../Tools/Luban/gen_code_unity_json.bat");

        #region hot update path
        //热更配置
        public static string HUConfigPath = Path.Combine(Application.dataPath, "Resources/HUConfig.json");

        //热更资源Assets位置
        public static string HUCodeResPath = "/HURes/Code";

        //YooAsset打包资源拷贝位置
        public static string HUResCopyPath = Path.Combine(Application.dataPath, "../Bundles/");
        //"Assets/StreamingAssets/BuildinFiles";

        // 版本记录
        public static string HURecordPath = Path.Combine(Application.dataPath, "../../BuildRecord/");
        #endregion

        #region local server path
        //本地python服务器位置
        public static string LocalServerPath = Path.Combine(Application.dataPath, "../../Server/");
        //py3
        public static string LocalServerBatPath = Path.Combine(LocalServerPath, "start.bat");
        //py2
        public static string LocalServerBatPy2Path = Path.Combine(LocalServerPath, "start_py2.bat");

        //本地服务器热更位置
        public static string LocalServerHotUpdatePath = Path.Combine(LocalServerPath, "HotUpdate"); 
        #endregion

        /// <summary>
        /// 获取本地服务器的更新位置
        /// </summary>
        /// <returns></returns>
        public static string GetServerHotUpdatePath(BuildTarget target) {
            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            JSONNode config = JSONNode.Parse(txt.text);
            string gameVersion = config["game_version"];

            return $"{LocalServerHotUpdatePath}/{target}/{gameVersion}";
        }

        /// <summary>
        /// 获取本地build资源位置
        /// </summary>
        /// <returns></returns>
        public static string GetLocalBuildPath(BuildTarget target) {
            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            JSONNode config = JSONNode.Parse(txt.text);

            string version = config["hot_update_version"];
            string PackageName = config["default_package_name"];

            return $"{HUResCopyPath}/{PackageName}/{target}/{version}";
        }

        /// <summary>
        /// 获取版本记录位置
        /// </summary>
        /// <param name="target"></param>
        /// <returns></returns>

        public static string GetBuildRecordPath(BuildTarget target) {
            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            JSONNode config = JSONNode.Parse(txt.text);

            string gameVersion = config["game_version"];
            string PackageName = config["default_package_name"];

            return $"{HURecordPath}/{target}/{gameVersion}";
        }
    }
}

