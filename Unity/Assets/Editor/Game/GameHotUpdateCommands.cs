using HybridCLR.Editor;
using HybridCLR.Editor.Commands;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEditor;
using UnityEngine;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// 重新打包命令
    /// </summary>
    public static class GameHotUpdateCommands {
        // 1. 编译需求平台的 热更代码dll
        [MenuItem("Game/HotUpdate/1.CompileTargetHotUpdateDll")]
        public static void CompileTargetHotUpdateDll() {
            GameHybridCLRCommands.CompileTargetHotUpdateDll();
        }

        // 2.拷贝编译好的 热更新代码dll 到 YooAsset 的代码资源更新目录下
        [MenuItem("Game/HotUpdate/2.CopyHotUpdateDllToRes")]
        public static void CopyHotUpdateDllToRes() {
            GameHybridCLRCommands.CopyHotUpdateDllToRes();
        }

        // 3.打开 YooAsset Collector 检查设置
        [MenuItem("Game/HotUpdate/3.OpenYooAssetCollector")]
        public static void OpenYooAssetCollector() {
            YooAsset.Editor.AssetBundleCollectorWindow.ShowExample();
        }

        // 4.打开 HUConfig.json 修改版本
        [MenuItem("Game/HotUpdate/4.ChangeHUVersion")]
        public static void ChangeHUVersion() {
            System.Diagnostics.Process.Start(GameEditorConfig.HUConfigPath);
        }

        // 5.打开 YooAsset Builder，检查设置，修改 build Mode 为 incremental rebuild, OK则开始构建

        [MenuItem("Game/HotUpdate/5.OpenYooAssetBuilder")]
        public static void OpenYooAssetBuilder() {
            YooAsset.Editor.AssetBundleBuilderWindow.ShowExample();
        }

        // 6.拷贝资源到服务器热更目录
        [MenuItem("Game/HotUpdate/6.CopyResToServer")]
        public static void CopyResToServer() {
            GameServerCommands.CopyResToServer();
        }
    }
}

