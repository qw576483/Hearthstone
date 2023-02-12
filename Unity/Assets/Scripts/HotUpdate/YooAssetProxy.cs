using UnityEngine;
using UniFramework.Event;
using UniFramework.Module;

using YooAsset;
using System.Collections;
using System;
using Newtonsoft.Json;

namespace UnityGeneralFramework.HotUpdateLogic {

    public class YooAssetProxy : MonoBehaviour {
        // 资源系统运行模式
        public EPlayMode PlayMode = EPlayMode.EditorSimulateMode;
        void Awake() {
            Debug.Log($"资源系统运行模式：{PlayMode}");
            Application.targetFrameRate = 60;
            Application.runInBackground = true;
        }
        void Start() {
            // 初始化BetterStreaming
            BetterStreamingAssets.Initialize();

            // 初始化事件系统
            UniEvent.Initalize();

            // 初始化管理系统
            UniModule.Initialize();

            // 初始化资源系统
            YooAssets.Initialize();
            YooAssets.SetOperationSystemMaxTimeSlice(30);

            // 创建补丁管理器
            UniModule.CreateModule<PatchManager>();

            // 开始补丁更新流程
            PatchManager.Instance.Run(PlayMode);
        }

        #region 静态区
        /// <summary>
        /// 编辑界面不从热更界面走，就会导致YooAsset没有初始化
        /// 这时候就调用这里就行
        /// </summary>
        /// <param name="cb"></param>
        /// <returns></returns>
        public static IEnumerator InitPackage(Func<bool> cb) {
#if UNITY_EDITOR
            YooAssets.Initialize();

            yield return new WaitForSeconds(0.5f);

            TextAsset txt = Resources.Load<TextAsset>("HUConfig");
            HUConfigJson config = JsonConvert.DeserializeObject<HUConfigJson>(txt.text);

            string packageName = config.default_package_name;
            var package = YooAssets.TryGetAssetsPackage(packageName);
            if (package == null) {
                package = YooAssets.CreateAssetsPackage(packageName);
                YooAssets.SetDefaultAssetsPackage(package);
            }

            // 编辑器下的模拟模式
            InitializationOperation initializationOperation = null;
            var createParameters = new EditorSimulateModeParameters();
            createParameters.SimulatePatchManifestPath = EditorSimulateModeHelper.SimulateBuild(packageName);
            initializationOperation = package.InitializeAsync(createParameters);

            yield return initializationOperation;
            if (package.InitializeStatus == EOperationStatus.Succeed) {
                if (cb != null) {
                    cb();
                }
            } else {
                Debug.LogWarning($"{initializationOperation.Error}");
            }
#endif
        }

        /// <summary>
        /// LoadAssetSync
        /// 同步加载Asset
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="path"></param>
        /// <returns></returns>
        public static T LoadAssetSync<T>(string path) where T:UnityEngine.Object {
            return YooAssets.LoadAssetSync<T>(path).AssetObject as T;
        }

        /// <summary>
        /// LoadRawFileFileTextSync
        /// 同步加载RawFile，获取文件string
        /// </summary>
        /// <param name="path"></param>
        /// <returns></returns>
        public static string LoadRawFileFileTextSync(string path) {
            return YooAssets.LoadRawFileSync(path).GetRawFileText();
        }
        /// <summary>
        /// LoadRawFileDataTextSync
        /// 同步加载RawFile，获取文件byte[]
        /// </summary>
        /// <param name="path"></param>
        /// <returns></returns>
        public static Byte[] LoadRawFileDataTextSync(string path) {
            return YooAssets.LoadRawFileSync(path).GetRawFileData();
        }

        #endregion
    }
}