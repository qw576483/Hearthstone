using UnityEngine;
using UniFramework.Event;
using UniFramework.Module;

using YooAsset;

namespace UnityGeneralFramework.HotUpdateLogic {

    public class YooAssetProxy : MonoBehaviour {
        /// <summary>
        /// ��Դϵͳ����ģʽ
        /// </summary>
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
    }
}