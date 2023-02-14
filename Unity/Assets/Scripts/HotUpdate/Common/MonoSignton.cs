using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace UnityGeneralFramework.Common {
    /// <summary>
    /// Mono��������
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class MonoSingleton<T> : MonoBehaviour where T : MonoSingleton<T> {
        // T2 ��ʾ��������
        private static T instance;
        public static T Instance {
            get {
                if (instance == null) {
                    instance = FindObjectOfType<T>();
                    if (instance == null) {
                        //�����ű�����
                        instance = new GameObject("Singleton of " + typeof(T)).AddComponent<T>();
                    } else {
                        instance.AwakeSingleton();
                    }
                }
                return instance;
            }
        }

        protected void Awake() {
            if (instance == null) {
                instance = this as T;
                AwakeSingleton();
            }
        }

        /// <summary>
        /// ������Awake
        /// </summary>
        public virtual void AwakeSingleton() {}

        /// <summary>
        /// ��ʼ��
        /// </summary>
        public virtual void OnInit() { }
       

        /*
            * ��ע��
            * 1.�����ԣ������д���Ψһ�Ķ��󣬼����øö���̳е�ǰ��
            * 2.������ã�
            *  -- �̳�ʱ���봫����������
            *  -- ������ű����������У�ͨ���������ͷ���Instance����
            */
    }
}
