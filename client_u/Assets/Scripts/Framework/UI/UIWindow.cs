using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

namespace ThisGame.Framework.UI {
    /// <summary>
    /// ���ڻ���
    /// </summary>
    [RequireComponent(typeof(CanvasGroup))]
    public class UIWindow : MonoBehaviour {
        public CanvasGroup canvasGroup;
        public Image imgBg;

        void Awake() {
            canvasGroup = GetComponent<CanvasGroup>();
            imgBg = GetComponent<Image>();
        }

        /// <summary>
        /// ��ʾ���ڡ����
        /// </summary>
        public void Show() { SetVisible(true);}
        /// <summary>
        /// ���ش��ڡ����
        /// </summary>
        public void Hide() { SetVisible(false);}

        /// <summary>
        /// ���ô��ڿɼ�
        /// </summary>
        /// <param name="visible">�Ƿ�ɼ�</param>
        /// <param name="delay">�ӳ�ʱ��</param>
        public void SetVisible(bool visible, float delay = 0) {
            // delay Ĭ��ֵ��0����Ȼ��0������Ϊ��Э�̣����Բ�������ִ�У����ǻ�����һִ֡��
            // ����� delay Ϊ0 ʱ����һִ֡�в����������������ٸĳɵ� delay Ϊ0ʱ������Э��
            StartCoroutine(SetVisibleDelay(visible, delay));
        }

        /// <summary>
        /// ��������Э��
        /// </summary>
        /// <param name="state">�Ƿ�ɼ�</param>
        /// <param name="delay">�ӳ�ʱ��</param>
        /// <returns></returns>
        private IEnumerator SetVisibleDelay(bool visible, float delay) {
            yield return new WaitForSeconds(delay);
            // Canvas Group
            canvasGroup.alpha = visible ? 1 : 0;

            //�Ƿ�Ҫ�ã�
            //gameObject.active = visible;
        }
    }
}
