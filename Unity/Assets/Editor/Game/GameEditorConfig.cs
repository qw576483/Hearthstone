using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEditor;
using UnityEngine;

using UnityGeneralFramework.HotUpdateLogic;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// Editor ����
    /// </summary>
    public static class GameEditorConfig{
        public static string ExcelBatPath = Path.Combine(Application.dataPath, "../Tools/Luban/gen_code_unity_json.bat");

        #region hot update path
        //�ȸ�����
        public static string HUConfigPath = Path.Combine(Application.dataPath, "Resources/HUConfig.json");

        //�ȸ���ԴAssetsλ��
        public static string HUCodeResPath = "/HURes/Code";

        //YooAsset�����Դ����λ��
        public static string HUResCopyPath = Path.Combine(Application.dataPath, "../Bundles/");
        //"Assets/StreamingAssets/BuildinFiles";

        // �汾��¼
        public static string HURecordPath = Path.Combine(Application.dataPath, "../../BuildRecord/");
        #endregion

        #region local server path
        //����python������λ��
        public static string LocalServerPath = Path.Combine(Application.dataPath, "../../Server/");
        //py3
        public static string LocalServerBatPath = Path.Combine(LocalServerPath, "start.bat");
        //py2
        public static string LocalServerBatPy2Path = Path.Combine(LocalServerPath, "start_py2.bat");

        //���ط������ȸ�λ��
        public static string LocalServerHotUpdatePath = Path.Combine(LocalServerPath, "HotUpdate"); 
        #endregion

        /// <summary>
        /// ��ȡ���ط������ĸ���λ��
        /// </summary>
        /// <returns></returns>
        public static string GetServerHotUpdatePath(BuildTarget target) {
            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            JSONNode config = JSONNode.Parse(txt.text);
            string gameVersion = config["game_version"];

            return $"{LocalServerHotUpdatePath}/{target}/{gameVersion}";
        }

        /// <summary>
        /// ��ȡ����build��Դλ��
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
        /// ��ȡ�汾��¼λ��
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

