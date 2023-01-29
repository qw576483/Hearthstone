using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class MonoSingleton<T> : MonoBehaviour where T : MonoSingleton<T> {
    // T2 表示子类类型
    private static T instance;
    public static T Instance {
        get {
            if (instance == null) {
                instance = FindObjectOfType<T>();
                if (instance == null) {
                    //创建脚本对象
                    instance = new GameObject("Singleton of " + typeof(T)).AddComponent<T>();
                } else {
                    instance.OnInit();
                }
            }
            return instance;
        }
    }

    protected void Awake() {
        if (instance == null) {
            instance = this as T;
            OnInit();
        }
    }

    public virtual void OnInit() {
            
    }

    /*
        * 备注：
        * 1.适用性：场景中存在唯一的对象，即可让该对象继承当前类
        * 2.如何适用：
        *  -- 继承时必须传递子类类型
        *  -- 在任意脚本生命周期中，通过子类类型访问Instance属性
        */
}
