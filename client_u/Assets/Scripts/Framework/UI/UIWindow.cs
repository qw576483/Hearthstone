using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

namespace ThisGame.Framework.UI {
    /// <summary>
    /// 窗口基类
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
        /// 显示窗口。便捷
        /// </summary>
        public void Show() { SetVisible(true);}
        /// <summary>
        /// 隐藏窗口。便捷
        /// </summary>
        public void Hide() { SetVisible(false);}

        /// <summary>
        /// 设置窗口可见
        /// </summary>
        /// <param name="visible">是否可见</param>
        /// <param name="delay">延迟时间</param>
        public void SetVisible(bool visible, float delay = 0) {
            // delay 默认值是0，虽然是0，但因为是协程，所以不会立即执行，而是会在下一帧执行
            // 如果对 delay 为0 时，下一帧执行不满足需求，那我们再改成当 delay 为0时，不走协程
            StartCoroutine(SetVisibleDelay(visible, delay));
        }

        /// <summary>
        /// 窗口显隐协程
        /// </summary>
        /// <param name="state">是否可见</param>
        /// <param name="delay">延迟时间</param>
        /// <returns></returns>
        private IEnumerator SetVisibleDelay(bool visible, float delay) {
            yield return new WaitForSeconds(delay);
            // Canvas Group
            canvasGroup.alpha = visible ? 1 : 0;

            //是否要用？
            //gameObject.active = visible;
        }
    }
}
