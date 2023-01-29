using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using UnityEngine;

namespace UnityGeneralFramework.Editor.Common {
    /// <summary>
    /// 暂时未归类，通用的方法
    /// </summary>
    public static class SomeCommon {

        /// <summary>
        /// windows 下运行 bat 文件
        /// </summary>
        /// <param name="batPath"></param>
        public static void RunBat(string batPath) {
            Process pro = new Process();

            FileInfo file = new FileInfo(batPath);
            pro.StartInfo.WorkingDirectory = file.Directory.FullName;
            pro.StartInfo.FileName = batPath;
            pro.StartInfo.CreateNoWindow = false;
            pro.Start();
            pro.WaitForExit();
        }

        /// <summary>
        /// 拷贝全部文件
        /// </summary>
        /// <param name="srcPath">源路径</param>
        /// <param name="aimPath">目标路径</param>
        public static void CopyDir(string srcPath, string aimPath) {
            // 检查目标目录是否以目录分割字符结束如果不是则添加
            if(aimPath[aimPath.Length - 1] != System.IO.Path.DirectorySeparatorChar) {
                aimPath += System.IO.Path.DirectorySeparatorChar;
            }
            // 判断目标目录是否存在如果不存在则新建
            if(!System.IO.Directory.Exists(aimPath)) {
                System.IO.Directory.CreateDirectory(aimPath);
            }
            // 得到源目录的文件列表，该里面是包含文件以及目录路径的一个数组
            // 如果你指向copy目标文件下面的文件而不包含目录请使用下面的方法
            // string[] fileList = Directory.GetFiles（srcPath）；

            string[] fileList = System.IO.Directory.GetFileSystemEntries(srcPath);

            // 遍历所有的文件和目录
            foreach(string file in fileList) {
                // 先当作目录处理如果存在这个目录就递归Copy该目录下面的文件
                if(System.IO.Directory.Exists(file)) {
                    CopyDir(file, aimPath + System.IO.Path.GetFileName(file));
                }
                // 否则直接Copy文件
                else {
                    System.IO.File.Copy(file, aimPath + System.IO.Path.GetFileName(file), true);
                }
            }
        }

        /// <summary>
        /// 拷贝筛选文件
        /// </summary>
        /// <param name="srcPath">源路径</param>
        /// <param name="aimPath">目标路径</param>
        /// <param name="check">筛选函数</param>

        public static void CopyDir(string srcPath, string aimPath, Func<string, bool> check) {
            // 检查目标目录是否以目录分割字符结束如果不是则添加
            if(aimPath[aimPath.Length - 1] != System.IO.Path.DirectorySeparatorChar) {
                aimPath += System.IO.Path.DirectorySeparatorChar;
            }
            // 判断目标目录是否存在如果不存在则新建
            if(!System.IO.Directory.Exists(aimPath)) {
                System.IO.Directory.CreateDirectory(aimPath);
            }

            // 得到源目录的文件列表，该里面是包含文件以及目录路径的一个数组
            // 如果你指向copy目标文件下面的文件而不包含目录请使用下面的方法
            // string[] fileList = Directory.GetFiles（srcPath）；
                
            string[] fileList = System.IO.Directory.GetFileSystemEntries(srcPath);

            // 遍历所有的文件和目录
            foreach(string file in fileList) {
                // 先当作目录处理如果存在这个目录就递归Copy该目录下面的文件
                if(System.IO.Directory.Exists(file) ) {
                    CopyDir(file, aimPath + System.IO.Path.GetFileName(file), check);
                }
                // 否则直接Copy文件
                else {
                    if( check( System.IO.Path.GetFileName(file) ) ) { 
                        System.IO.File.Copy(file, aimPath + System.IO.Path.GetFileName(file), true);
                    }
                }
            }
        }

    }
}
