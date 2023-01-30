using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using HybridCLR;
using YooAsset;

namespace UnityGeneralFramework.HotUpdateLogic {
    public class HybridCLRProxy : MonoBehaviour {

        #region demo ��� loadDll.cs ��������ʹ��

        /// <summary>
        /// ��ʼ����
        /// </summary>
        public void StartGame() {
            Debug.Log("HybridCLR ��ʼ");
            // 1.���ز���Ԫ���ݴ���
            LoadMetadataForAOTAssemblies();
            // 2.�����ȸ�����
            LoadHotUpdateAssemblies();
            // 3.���½����������ȸ���ز���
        }

        /// <summary>
        /// Ϊaot assembly����ԭʼmetadata�� ��������aot�����ȸ��¶��С�
        /// һ�����غ����AOT���ͺ�����Ӧnativeʵ�ֲ����ڣ����Զ��滻Ϊ����ģʽִ��
        /// </summary>
        private void LoadMetadataForAOTAssemblies() {
            Debug.Log("��ʼ���ز���Ԫ����");
            /// ע�⣬����Ԫ�����Ǹ�AOT dll����Ԫ���ݣ������Ǹ��ȸ���dll����Ԫ���ݡ�
            /// �ȸ���dll��ȱԪ���ݣ�����Ҫ���䣬�������LoadMetadataForAOTAssembly�᷵�ش���

            // TODO:
            HomologousImageMode mode = HomologousImageMode.SuperSet;
            foreach (var aotDllName in HUConfig.AOTMetaAssemblyNames) {
                byte[] dllBytes = YooAssets.LoadRawFileSync($"Assets/HURes/Code/{aotDllName}").GetRawFileData();
                // ����assembly��Ӧ��dll�����Զ�Ϊ��hook��
                // һ��aot���ͺ�����native���������ڣ��ý������汾����
                LoadImageErrorCode err = RuntimeApi.LoadMetadataForAOTAssembly(dllBytes, mode);
                Debug.Log($"LoadMetadataForAOTAssembly:{aotDllName}. mode:{mode} ret:{err}");
            }
        }

        /// <summary>
        /// �����ȸ� assembly
        /// </summary>
        private void LoadHotUpdateAssemblies() {
            Debug.Log("��ʼ�����ȸ�����");
#if !UNITY_EDITOR
        //System.Reflection.Assembly.Load( YooAssets.LoadRawFileSync("Assets/HURes/Code/HU.Code.dll").GetRawFileData() );
#endif
            //AssetBundle prefabAb = AssetBundle.LoadFromMemory(GetAssetData("prefabs"));
            //GameObject testPrefab = Instantiate(prefabAb.LoadAsset<GameObject>("HotUpdatePrefab.prefab"));
            System.Reflection.Assembly.Load(YooAssets.LoadRawFileSync("Assets/HURes/Code/HU.Code.dll").GetRawFileData());
            Debug.Log("�������");
        }
        #endregion
    }
}
