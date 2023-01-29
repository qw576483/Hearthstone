using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using UnityEngine;

namespace UnityGeneralFramework.Editor.Common {
    /// <summary>
    /// ��ʱδ���࣬ͨ�õķ���
    /// </summary>
    public static class SomeCommon {

        /// <summary>
        /// windows ������ bat �ļ�
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
        /// ����ȫ���ļ�
        /// </summary>
        /// <param name="srcPath">Դ·��</param>
        /// <param name="aimPath">Ŀ��·��</param>
        public static void CopyDir(string srcPath, string aimPath) {
            // ���Ŀ��Ŀ¼�Ƿ���Ŀ¼�ָ��ַ�����������������
            if(aimPath[aimPath.Length - 1] != System.IO.Path.DirectorySeparatorChar) {
                aimPath += System.IO.Path.DirectorySeparatorChar;
            }
            // �ж�Ŀ��Ŀ¼�Ƿ����������������½�
            if(!System.IO.Directory.Exists(aimPath)) {
                System.IO.Directory.CreateDirectory(aimPath);
            }
            // �õ�ԴĿ¼���ļ��б��������ǰ����ļ��Լ�Ŀ¼·����һ������
            // �����ָ��copyĿ���ļ�������ļ���������Ŀ¼��ʹ������ķ���
            // string[] fileList = Directory.GetFiles��srcPath����

            string[] fileList = System.IO.Directory.GetFileSystemEntries(srcPath);

            // �������е��ļ���Ŀ¼
            foreach(string file in fileList) {
                // �ȵ���Ŀ¼��������������Ŀ¼�͵ݹ�Copy��Ŀ¼������ļ�
                if(System.IO.Directory.Exists(file)) {
                    CopyDir(file, aimPath + System.IO.Path.GetFileName(file));
                }
                // ����ֱ��Copy�ļ�
                else {
                    System.IO.File.Copy(file, aimPath + System.IO.Path.GetFileName(file), true);
                }
            }
        }

        /// <summary>
        /// ����ɸѡ�ļ�
        /// </summary>
        /// <param name="srcPath">Դ·��</param>
        /// <param name="aimPath">Ŀ��·��</param>
        /// <param name="check">ɸѡ����</param>

        public static void CopyDir(string srcPath, string aimPath, Func<string, bool> check) {
            // ���Ŀ��Ŀ¼�Ƿ���Ŀ¼�ָ��ַ�����������������
            if(aimPath[aimPath.Length - 1] != System.IO.Path.DirectorySeparatorChar) {
                aimPath += System.IO.Path.DirectorySeparatorChar;
            }
            // �ж�Ŀ��Ŀ¼�Ƿ����������������½�
            if(!System.IO.Directory.Exists(aimPath)) {
                System.IO.Directory.CreateDirectory(aimPath);
            }

            // �õ�ԴĿ¼���ļ��б��������ǰ����ļ��Լ�Ŀ¼·����һ������
            // �����ָ��copyĿ���ļ�������ļ���������Ŀ¼��ʹ������ķ���
            // string[] fileList = Directory.GetFiles��srcPath����
                
            string[] fileList = System.IO.Directory.GetFileSystemEntries(srcPath);

            // �������е��ļ���Ŀ¼
            foreach(string file in fileList) {
                // �ȵ���Ŀ¼��������������Ŀ¼�͵ݹ�Copy��Ŀ¼������ļ�
                if(System.IO.Directory.Exists(file) ) {
                    CopyDir(file, aimPath + System.IO.Path.GetFileName(file), check);
                }
                // ����ֱ��Copy�ļ�
                else {
                    if( check( System.IO.Path.GetFileName(file) ) ) { 
                        System.IO.File.Copy(file, aimPath + System.IO.Path.GetFileName(file), true);
                    }
                }
            }
        }

    }
}
