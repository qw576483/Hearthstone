using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace ThisGame.Framework {
    /// <summary>
    /// 变换组件助手类
    /// </summary>
    public static class TransformHelper {
        /// <summary>
        /// 未知层级，查找后代指定名称的组件
        /// 就是递归而已。。。
        /// </summary>
        /// <param name="currentTF">当前transform</param>
        /// <param name="name">子物体名称</param>
        /// <returns></returns>
        public static Transform FindChildByName(this Transform currentTF, string childName) {
            Transform childTF = currentTF.Find(childName);
            if(childTF != null) return childTF;

            for(int i = 0; i < currentTF.childCount; i++) {
                childTF = FindChildByName(currentTF.GetChild(i), childName);
                if(childTF != null) return childTF;
            }

            return null;
        }
    }

}

