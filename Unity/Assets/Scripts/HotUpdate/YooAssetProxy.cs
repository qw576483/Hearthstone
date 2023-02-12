using UnityEngine;
using UniFramework.Event;
using UniFramework.Module;

using YooAsset;
using System.Collections;
using System;
using Newtonsoft.Json;

namespace UnityGeneralFramework.HotUpdateLogic {

    public class YooAssetProxy : MonoBehaviour {
        // ��Դϵͳ����ģʽ
        public EPlayMode PlayMode = EPlayMode.EditorSimulateMode;
        void Awake() {
            Debug.Log($"��Դϵͳ����ģʽ��{PlayMode}");
            Application.targetFrameRate = 60;
            Application.runInBackground = true;
        }
        void Start() {
            // ��ʼ��BetterStreaming
            BetterStreamingAssets.Initialize();

            // ��ʼ���¼�ϵͳ
            UniEvent.Initalize();

            // ��ʼ������ϵͳ
            UniModule.Initialize();

            // ��ʼ����Դϵͳ
            YooAssets.Initialize();
            YooAssets.SetOperationSystemMaxTimeSlice(30);

            // ��������������
            UniModule.CreateModule<PatchManager>();

            // ��ʼ������������
            PatchManager.Instance.Run(PlayMode);
        }

        #region ��̬��
        /// <summary>
        /// �༭���治���ȸ������ߣ��ͻᵼ��YooAssetû�г�ʼ��
        /// ��ʱ��͵����������
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

            // �༭���µ�ģ��ģʽ
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
        /// ͬ������Asset
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="path"></param>
        /// <returns></returns>
        public static T LoadAssetSync<T>(string path) where T:UnityEngine.Object {
            return YooAssets.LoadAssetSync<T>(path).AssetObject as T;
        }

        /// <summary>
        /// LoadRawFileFileTextSync
        /// ͬ������RawFile����ȡ�ļ�string
        /// </summary>
        /// <param name="path"></param>
        /// <returns></returns>
        public static string LoadRawFileFileTextSync(string path) {
            return YooAssets.LoadRawFileSync(path).GetRawFileText();
        }
        /// <summary>
        /// LoadRawFileDataTextSync
        /// ͬ������RawFile����ȡ�ļ�byte[]
        /// </summary>
        /// <param name="path"></param>
        /// <returns></returns>
        public static Byte[] LoadRawFileDataTextSync(string path) {
            return YooAssets.LoadRawFileSync(path).GetRawFileData();
        }

        #endregion
    }
}