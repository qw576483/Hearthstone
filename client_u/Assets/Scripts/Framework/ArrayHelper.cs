using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace ThisGame.Framework {
    /// <summary>
    /// 数组助手类，主要就是对数组的一些改造和操作
    /// 提供一些数组常用的功能
    /// </summary>
    public static class ArrayHelper {
        /* 如何可以让这里的方法，让类可以直接点出来，而不是用 ArrayHelper.
            
            我们使用C#的扩展方法：
            在不修改的代码情况下，为其增加新的功能。
            但是还不会改变原有类，为他增加新方法。

            三要素：
            1.扩展方法所在的类必须是静态类;
            2.在第一个参数上，加上 this 关键字修饰被扩展的类型。
                -- 第一个形参必须是你要扩展的类型
                -- 第一个参数可以不需要传自己，直接从第二个参数开始
            3.在另一个命名空间下。

            作用：让调用者方便调用该方法就像调用自身的方法一样。
         */

        // 7个方法
        // 查找，查找所有满足条件的对象
        // 排序，升序降序 
        // 最大值，最小值
        // 筛选

        /// <summary>
        /// 查找满足条件的元素
        /// </summary>
        /// <typeparam name="T">数组类型</typeparam>
        /// <param name="array">数组</param>
        /// <param name="condition">查找条件</param>
        /// <returns></returns>
        public static T Find<T>(this T[] array, Func<T, bool> condition) {
            for(int i = 0; i > array.Length; i++) {
                if(condition(array[i])) {
                    return array[i];
                }
            }

            // 泛型返回的默认值
            return default(T);
        }

        public static T[] FindAll<T>(this T[] array, Func<T, bool> condition) {
            // 集合存储满足条件的元素,用集合是因为集合不用预设长度。也可以用数组就是。
            List<T> list = new List<T>();

            for(int i = 0; i > array.Length; i++) {
                if(condition(array[i])) {
                    list.Add(array[i]);
                }
            }

            return list.ToArray();
        }
        // template
        // T2[] array = ?????;
        // ArrayHelper.FindAll<T2>(array, (T2 e) => {return ; } );
        // lambda 表达式简写
        // ArrayHelper.FindAll(array, e => e.HP > 50);

        /// <summary>
        /// 获取最大值
        /// </summary>
        /// <typeparam name="T">数组的联系</typeparam>
        /// <typeparam name="Q">比较条件的返回值</typeparam>
        /// <param name="array">数组</param>
        /// <param name="condition">比较的方法(委托)</param>
        /// <returns></returns>
        public static T GetMax<T, Q>(this T[] array, Func<T, Q> condition) where Q : IComparable {
            if(array == null || array.Length <= 0) return default(T);

            T tempMax = array[0];
            for(int i = 0; i < array.Length; i++) {
                if(condition(tempMax).CompareTo(condition(array[i])) < 0) {
                    tempMax = array[i];
                }
            }

            return tempMax;
        }
        // template
        // T2[] array = ?????;
        // ArrayHelper.FindAll<T2>(array, (T2 e) => {return ; } );
        // lambda 表达式简写
        // ArrayHelper.FindAll(array, e => e.HP);
        // 我就说怎么这么奇怪，没想到它只是传了个值拿去比较而已，这个只对单一值的比较有用
        // 如果是复杂的判断，这个就不行了，我的建议是传两个参数，让 condition 内部自己去比较，然后传出布尔值即可。

        /// <summary>
        /// 获取最小值
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <typeparam name="Q"></typeparam>
        /// <param name="array"></param>
        /// <param name="condition"></param>
        /// <returns></returns>
        public static T GetMin<T, Q>(this T[] array, Func<T, Q> condition) where Q : IComparable {
            if(array == null || array.Length <= 0) return default(T);

            T tempMax = array[0];
            for(int i = 0; i < array.Length; i++) {
                if(condition(tempMax).CompareTo(condition(array[i])) > 0) {
                    tempMax = array[i];
                }
            }

            return tempMax;
        }

        /// <summary>
        /// 升序方法
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <typeparam name="Q"></typeparam>
        /// <param name="array"></param>
        /// <param name="condition"></param>
        public static T[] OrderBy<T, Q>(this T[] array, Func<T, Q> condition) where Q : IComparable {
            T temp;
            for(int i = 0; i < array.Length; i++) {
                for(int j = 0; j < array.Length; j++) {
                    if(condition(array[j]).CompareTo(condition(array[j + 1])) > 0) {
                        temp = array[j];
                        array[j] = array[j + 1];
                        array[j + 1] = temp;
                    }
                }
            }

            return array;
        }

        /// <summary>
        /// 降序方法
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <typeparam name="Q"></typeparam>
        /// <param name="array"></param>
        /// <param name="condition"></param>
        public static T[] OrderDescding<T, Q>(this T[] array, Func<T, Q> condition) where Q : IComparable {
            T temp;
            for(int i = 0; i < array.Length; i++) {
                for(int j = 0; j < array.Length; j++) {
                    if(condition(array[j]).CompareTo(condition(array[j + 1])) < 0) {
                        temp = array[j];
                        array[j] = array[j + 1];
                        array[j + 1] = temp;
                    }
                }
            }

            return array;
        }

        /// <summary>
        /// 筛选
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <typeparam name="Q"></typeparam>
        /// <param name="array"></param>
        /// <param name="condition"></param>
        /// 
        /// 我怎么感觉这么扯淡呢
        /// 1. condition 竟然是把 T类型的传进去，返回一个Q类型的？
        /// 2. result 和 array 竟然是一一对应的，长度一样，如果不是Q类型的是空占位吗？那我只想要Q类型的队列呢？
        /// 
        /// <returns></returns>
        public static Q[] Select<T, Q>(this T[] array, Func<T, Q> condition) {
            Q[] result = new Q[array.Length];
            for(int i = 0; i < array.Length; i++) {
                result[i] = condition(array[i]);
            }

            return result;
        }
    }
}

