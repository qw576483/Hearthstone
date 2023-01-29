using System.Collections;
using System.Collections.Generic;
using System.IO;
using HybridCLR.Editor;
using HybridCLR.Editor.Commands;
using UnityEditor;
using UnityEngine;

namespace UnityGeneralFramework.Editor {
    public static class GameServerCommands {
        /// <summary>
        /// windows下使用简易的python http 服务器
        /// </summary>
        [MenuItem("Game/LocalServer/OpenSimpleServer")]
        public static void OpenSimpleServer() {
#if UNITY_EDITOR_WIN
            //Debug.Log( Path.Combine(Application.dataPath, "../") );
            Editor.Common.SomeCommon.RunBat(GameEditorConfig.LocalServerBatPath);
#endif
        }

        /// <summary>
        /// windows下使用简易的python2 simple http 服务器
        /// </summary>
        [MenuItem("Game/LocalServer/OpenSimpleServerPy2")]
        public static void OpenSimpleServerPy2() {
#if UNITY_EDITOR_WIN
            //Debug.Log( Path.Combine(Application.dataPath, "../") );
            Editor.Common.SomeCommon.RunBat(GameEditorConfig.LocalServerBatPy2Path);
#endif
        }

        // 7.拷贝资源到服务器热更目录
        [MenuItem("Game/LocalServer/CopyResToServer")]
        public static void CopyResToServer() {
            BuildTarget target = EditorUserBuildSettings.activeBuildTarget;

            string serverHotUpdateDir = GameEditorConfig.GetServerHotUpdatePath(target);

            //Debug.Log("serverHotUpdateDir " + serverHotUpdateDir);
            //Debug.Log("YooAssetOutputDir " + GameEditorConfig.HUResCopyPath);

            Common.SomeCommon.CopyDir(
                GameEditorConfig.GetLocalBuildPath(target),
                GameEditorConfig.GetServerHotUpdatePath(target),
                o => {
                    Debug.Log(o);
                    if(o.IndexOf(".meta") == -1) {
                        return
                           o.IndexOf(".rawfile") != -1 || o.IndexOf(".bundle") != -1
                        || o.IndexOf(".version") != -1 || o.IndexOf(".bytes") != -1
                        || o.IndexOf(".hash") != -1 || o.IndexOf(".json") != -1;
                    }
                    return false;
                });
        }

    }
}