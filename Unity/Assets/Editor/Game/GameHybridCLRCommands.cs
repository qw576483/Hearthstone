using System.Collections;
using System.Collections.Generic;
using System.IO;
using HybridCLR.Editor;
using HybridCLR.Editor.Commands;
using UnityEditor;
using UnityEngine;

using UnityGeneralFramework.HotUpdateLogic;

namespace UnityGeneralFramework.Editor {
    public static class GameHybridCLRCommands {
        // 3. 编译需求平台的 热更代码dll
        [MenuItem("Game/HybridCLR/CompileTargetHotUpdateDll")]
        public static void CompileTargetHotUpdateDll() {
            BuildTarget target = EditorUserBuildSettings.activeBuildTarget;
            CompileDllCommand.CompileDll(target);
        }


        // 4.拷贝编译好的 热更新代码dll 到 YooAsset 的代码资源更新目录下
        [MenuItem("Game/HybridCLR/CopyHotUpdateDllToRes")]
        public static void CopyHotUpdateDllToRes() {
            BuildTarget target = EditorUserBuildSettings.activeBuildTarget;
            //BuildAssetBundleByTarget(target);
            CopyHotUpdateAssembliesToRes(target);
        }

        /// <summary>
        /// 复制编译完的dll到热更新资源目录下
        /// </summary>
        public static void CopyHotUpdateAssembliesToRes(BuildTarget target) {
            string hotfixDllSrcDir = SettingsUtil.GetHotUpdateDllsOutputDirByTarget(target);
            string hotfixAssembliesDstDir = Application.dataPath + GameEditorConfig.HUCodeResPath;

            foreach (var dll in SettingsUtil.HotUpdateAssemblyFiles) {
                string dllPath = $"{hotfixDllSrcDir}/{dll}";
                string dllBytesPath = $"{hotfixAssembliesDstDir}/{dll}.bytes";
                File.Copy(dllPath, dllBytesPath, true);
                Debug.Log($"[CopyHotUpdateAssembliesToRes] copy hotfix dll {dllPath} -> {dllBytesPath}");
            }
        }
    }
}