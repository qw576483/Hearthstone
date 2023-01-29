using HybridCLR.Editor;
using HybridCLR.Editor.Commands;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEditor;
using UnityEngine;

using UnityGeneralFramework.HotUpdateLogic;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// 重新打包命令
    /// </summary>
    public static class GameRebuildCommands {
        /* 1.使用 HybridCLR 工具 Generate/All。
         * 目的是:
         * 重新生成 link.xml 和 AOTGenericReferences.cs
         * 重新生成 AOT 补充元数据的 dll（再Asset同级目录下的HybridCLRData/AssembliesPostIl2CppStrip下）。
         */
        [MenuItem("Game/Rebuild/1.HybridCLR->Generate->All")]
        public static void UseHybridToolGenerateAll() {
            PrebuildCommand.GenerateAll();
        }

        // 2.拷贝AOT补充元数据到热更新文件夹下
        [MenuItem("Game/Rebuild/2.CopyAOTMetaDllsToRes")]
        public static void CopyAOTMetaDllsToRes() {
            BuildTarget target = EditorUserBuildSettings.activeBuildTarget;

            string temp = SettingsUtil.AssembliesPostIl2CppStripDir;

            string hotfixDllSrcDir = SettingsUtil.GetAssembliesPostIl2CppStripDir(target);
            string hotfixAssembliesDstDir = Application.dataPath + GameEditorConfig.HUCodeResPath;

            foreach (string dll in HUConfig.AOTMetaAssemblyNames) {
                string dllPath = $"{hotfixDllSrcDir}/{dll}";
                string dllBytesPath = $"{hotfixAssembliesDstDir}/{dll}.bytes";
                File.Copy(dllPath, dllBytesPath, true);
                Debug.Log($"[CopyHotUpdateAssembliesToRes] copy hotfix dll {dllPath} -> {dllBytesPath}");
            }
        }

        // 3. 编译需求平台的 热更代码dll
        [MenuItem("Game/Rebuild/3.CompileTargetHotUpdateDll")]
        public static void CompileTargetHotUpdateDll() {
            GameHybridCLRCommands.CompileTargetHotUpdateDll();
        }

        // 4.拷贝编译好的 热更新代码dll 到 YooAsset 的代码资源更新目录下
        [MenuItem("Game/Rebuild/4.CopyHotUpdateDllToRes")]
        public static void CopyHotUpdateDllToRes() {
            GameHybridCLRCommands.CopyHotUpdateDllToRes();
        }

        // 5.打开 YooAsset Collector 检查设置
        [MenuItem("Game/Rebuild/5.OpenYooAssetCollector")]
        public static void OpenYooAssetCollector() {
            YooAsset.Editor.AssetBundleCollectorWindow.ShowExample();
        }

        // 6.打开 HUConfig.json 修改版本
        [MenuItem("Game/Rebuild/6.ChangeHUVersion")]
        public static void ChangeHUVersion() {
            //Debug.Log(GameEditorConfig.HUConfigPath);
            System.Diagnostics.Process.Start(GameEditorConfig.HUConfigPath);
        }

        // 6.打开 YooAsset Builder，检查设置，修改 build Mode 为 force rebuild, OK则开始构建
        [MenuItem("Game/Rebuild/7.OpenYooAssetBuilder")]
        public static void OpenYooAssetBuilder() {
            YooAsset.Editor.AssetBundleBuilderWindow.ShowExample();
        }

        // 7.拷贝资源到服务器热更目录
        [MenuItem("Game/Rebuild/8.CopyResToServer")]
        public static void CopyResToServer() {
            GameServerCommands.CopyResToServer();
        }

        // 8.记录这次的版本包内容
        [MenuItem("Game/Rebuild/9.RecordThisVersionRes")]
        public static void RecordThisVersionRes() {
            BuildTarget target = EditorUserBuildSettings.activeBuildTarget;

            string serverHotUpdateDir = GameEditorConfig.GetServerHotUpdatePath(target);

            //Debug.Log("serverHotUpdateDir " + serverHotUpdateDir);
            //Debug.Log("YooAssetOutputDir " + GameEditorConfig.HUResCopyPath);

            Common.SomeCommon.CopyDir(
                GameEditorConfig.GetLocalBuildPath(target),
                GameEditorConfig.GetBuildRecordPath(target),
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

