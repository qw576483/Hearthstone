using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using UnityEditor;
using UnityEngine;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// 其他工具
    /// </summary>
    public static class GameToolCommands {
        [MenuItem("Game/Tools/OutputExcels")]
        public static void OutputExcels() {
#if UNITY_EDITOR_WIN
            Debug.Log(GameEditorConfig.ExcelBatPath);
            Editor.Common.SomeCommon.RunBat(GameEditorConfig.ExcelBatPath);
#endif
        }
}
}
