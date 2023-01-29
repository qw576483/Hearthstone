using HybridCLR.Editor;
using HybridCLR.Editor.Commands;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEditor;
using UnityEngine;

namespace UnityGeneralFramework.Editor {
    /// <summary>
    /// ���´������
    /// </summary>
    public static class GameHotUpdateCommands {
        // 1. ��������ƽ̨�� �ȸ�����dll
        [MenuItem("Game/HotUpdate/1.CompileTargetHotUpdateDll")]
        public static void CompileTargetHotUpdateDll() {
            GameHybridCLRCommands.CompileTargetHotUpdateDll();
        }

        // 2.��������õ� �ȸ��´���dll �� YooAsset �Ĵ�����Դ����Ŀ¼��
        [MenuItem("Game/HotUpdate/2.CopyHotUpdateDllToRes")]
        public static void CopyHotUpdateDllToRes() {
            GameHybridCLRCommands.CopyHotUpdateDllToRes();
        }

        // 3.�� YooAsset Collector �������
        [MenuItem("Game/HotUpdate/3.OpenYooAssetCollector")]
        public static void OpenYooAssetCollector() {
            YooAsset.Editor.AssetBundleCollectorWindow.ShowExample();
        }

        // 4.�� HUConfig.json �޸İ汾
        [MenuItem("Game/HotUpdate/4.ChangeHUVersion")]
        public static void ChangeHUVersion() {
            System.Diagnostics.Process.Start(GameEditorConfig.HUConfigPath);
        }

        // 5.�� YooAsset Builder��������ã��޸� build Mode Ϊ incremental rebuild, OK��ʼ����

        [MenuItem("Game/HotUpdate/5.OpenYooAssetBuilder")]
        public static void OpenYooAssetBuilder() {
            YooAsset.Editor.AssetBundleBuilderWindow.ShowExample();
        }

        // 6.������Դ���������ȸ�Ŀ¼
        [MenuItem("Game/HotUpdate/6.CopyResToServer")]
        public static void CopyResToServer() {
            GameServerCommands.CopyResToServer();
        }
    }
}

