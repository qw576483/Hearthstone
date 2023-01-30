using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using HybridCLR;
using YooAsset;

namespace UnityGeneralFramework.HotUpdateLogic {
    public class HybridCLRProxy : MonoBehaviour {

        #region demo 里的 loadDll.cs 拷贝过来使用

        /// <summary>
        /// 开始加载
        /// </summary>
        public void StartGame() {
            Debug.Log("HybridCLR 开始");
            // 1.加载补充元数据代码
            LoadMetadataForAOTAssemblies();
            // 2.加载热更代码
            LoadHotUpdateAssemblies();
            // 3.更新结束，调用热更相关测试
        }

        /// <summary>
        /// 为aot assembly加载原始metadata， 这个代码放aot或者热更新都行。
        /// 一旦加载后，如果AOT泛型函数对应native实现不存在，则自动替换为解释模式执行
        /// </summary>
        private void LoadMetadataForAOTAssemblies() {
            Debug.Log("开始加载补充元数据");
            /// 注意，补充元数据是给AOT dll补充元数据，而不是给热更新dll补充元数据。
            /// 热更新dll不缺元数据，不需要补充，如果调用LoadMetadataForAOTAssembly会返回错误

            // TODO:
            HomologousImageMode mode = HomologousImageMode.SuperSet;
            foreach (var aotDllName in HUConfig.AOTMetaAssemblyNames) {
                byte[] dllBytes = YooAssets.LoadRawFileSync($"Assets/HURes/Code/{aotDllName}").GetRawFileData();
                // 加载assembly对应的dll，会自动为它hook。
                // 一旦aot泛型函数的native函数不存在，用解释器版本代码
                LoadImageErrorCode err = RuntimeApi.LoadMetadataForAOTAssembly(dllBytes, mode);
                Debug.Log($"LoadMetadataForAOTAssembly:{aotDllName}. mode:{mode} ret:{err}");
            }
        }

        /// <summary>
        /// 加载热更 assembly
        /// </summary>
        private void LoadHotUpdateAssemblies() {
            Debug.Log("开始加载热更代码");
#if !UNITY_EDITOR
        //System.Reflection.Assembly.Load( YooAssets.LoadRawFileSync("Assets/HURes/Code/HU.Code.dll").GetRawFileData() );
#endif
            //AssetBundle prefabAb = AssetBundle.LoadFromMemory(GetAssetData("prefabs"));
            //GameObject testPrefab = Instantiate(prefabAb.LoadAsset<GameObject>("HotUpdatePrefab.prefab"));
            System.Reflection.Assembly.Load(YooAssets.LoadRawFileSync("Assets/HURes/Code/HU.Code.dll").GetRawFileData());
            Debug.Log("加载完成");
        }
        #endregion
    }
}
