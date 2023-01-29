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
    /// ���´������
    /// </summary>
    public static class GameRebuildCommands {
        /* 1.ʹ�� HybridCLR ���� Generate/All��
         * Ŀ����:
         * �������� link.xml �� AOTGenericReferences.cs
         * �������� AOT ����Ԫ���ݵ� dll����Assetͬ��Ŀ¼�µ�HybridCLRData/AssembliesPostIl2CppStrip�£���
         */
        [MenuItem("Game/Rebuild/1.HybridCLR->Generate->All")]
        public static void UseHybridToolGenerateAll() {
            PrebuildCommand.GenerateAll();
        }

        // 2.����AOT����Ԫ���ݵ��ȸ����ļ�����
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

        // 3. ��������ƽ̨�� �ȸ�����dll
        [MenuItem("Game/Rebuild/3.CompileTargetHotUpdateDll")]
        public static void CompileTargetHotUpdateDll() {
            GameHybridCLRCommands.CompileTargetHotUpdateDll();
        }

        // 4.��������õ� �ȸ��´���dll �� YooAsset �Ĵ�����Դ����Ŀ¼��
        [MenuItem("Game/Rebuild/4.CopyHotUpdateDllToRes")]
        public static void CopyHotUpdateDllToRes() {
            GameHybridCLRCommands.CopyHotUpdateDllToRes();
        }

        // 5.�� YooAsset Collector �������
        [MenuItem("Game/Rebuild/5.OpenYooAssetCollector")]
        public static void OpenYooAssetCollector() {
            YooAsset.Editor.AssetBundleCollectorWindow.ShowExample();
        }

        // 6.�� HUConfig.json �޸İ汾
        [MenuItem("Game/Rebuild/6.ChangeHUVersion")]
        public static void ChangeHUVersion() {
            //Debug.Log(GameEditorConfig.HUConfigPath);
            System.Diagnostics.Process.Start(GameEditorConfig.HUConfigPath);
        }

        // 6.�� YooAsset Builder��������ã��޸� build Mode Ϊ force rebuild, OK��ʼ����
        [MenuItem("Game/Rebuild/7.OpenYooAssetBuilder")]
        public static void OpenYooAssetBuilder() {
            YooAsset.Editor.AssetBundleBuilderWindow.ShowExample();
        }

        // 7.������Դ���������ȸ�Ŀ¼
        [MenuItem("Game/Rebuild/8.CopyResToServer")]
        public static void CopyResToServer() {
            GameServerCommands.CopyResToServer();
        }

        // 8.��¼��εİ汾������
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

