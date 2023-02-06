﻿using System;
using System.IO;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UniFramework.Machine;
using UniFramework.Module;
using YooAsset;

using UnityGeneralFramework.HotUpdateLogic;

/// <summary>
/// 初始化资源包
/// </summary>
internal class FsmInitialize : IStateNode
{
	private StateMachine _machine;

	void IStateNode.OnCreate(StateMachine machine)
	{
		_machine = machine;
	}
	void IStateNode.OnEnter()
	{
		PatchEventDefine.PatchStatesChange.SendEventMessage("初始化资源包！");
		UniModule.StartCoroutine(InitPackage());
	}
	void IStateNode.OnUpdate()
	{
	}
	void IStateNode.OnExit()
	{
	}

	private IEnumerator InitPackage()
	{
		yield return new WaitForSeconds(1f);

		var playMode = PatchManager.Instance.PlayMode;

		// 创建默认的资源包
		string packageName = HUConfig.Instance.defPackageName;
		var package = YooAssets.TryGetAssetsPackage(packageName);
		if (package == null)
		{
			package = YooAssets.CreateAssetsPackage(packageName);
			YooAssets.SetDefaultAssetsPackage(package);
		}

		// 编辑器下的模拟模式
		InitializationOperation initializationOperation = null;
		if (playMode == EPlayMode.EditorSimulateMode)
		{
			var createParameters = new EditorSimulateModeParameters();
			createParameters.SimulatePatchManifestPath = EditorSimulateModeHelper.SimulateBuild(packageName);
			initializationOperation = package.InitializeAsync(createParameters);
		}

		// 单机运行模式
		if (playMode == EPlayMode.OfflinePlayMode)
		{
			var createParameters = new OfflinePlayModeParameters();
			createParameters.DecryptionServices = new GameDecryptionServices();
			initializationOperation = package.InitializeAsync(createParameters);
		}

		// 联机运行模式
		if (playMode == EPlayMode.HostPlayMode)
		{
			var createParameters = new HostPlayModeParameters();
			createParameters.DecryptionServices = new GameDecryptionServices();
			createParameters.QueryServices = new GameQueryServices();
			createParameters.DefaultHostServer = GetHostServerURL();
			createParameters.FallbackHostServer = GetHostServerURL();
			Debug.Log("createParameters " + createParameters);
			initializationOperation = package.InitializeAsync(createParameters);
		}

		yield return initializationOperation;
		if (package.InitializeStatus == EOperationStatus.Succeed)
		{
			_machine.ChangeState<FsmUpdateVersion>();
		}
		else
		{
			Debug.LogWarning($"{initializationOperation.Error}");
			PatchEventDefine.InitializeFailed.SendEventMessage();
		}
	}

	/// <summary>
	/// 获取资源服务器地址
	/// </summary>
	private string GetHostServerURL()
	{

		string hostServerIP = HUConfig.Instance.URL1;
		string gameVersion = HUConfig.Instance.gameVersion;

#if UNITY_EDITOR
		if (UnityEditor.EditorUserBuildSettings.activeBuildTarget == UnityEditor.BuildTarget.Android)
			return $"{hostServerIP}/hotUpdate/android/{gameVersion}";
		else if (UnityEditor.EditorUserBuildSettings.activeBuildTarget == UnityEditor.BuildTarget.iOS)
			return $"{hostServerIP}/hotUpdate/ios/{gameVersion}";
		else if (UnityEditor.EditorUserBuildSettings.activeBuildTarget == UnityEditor.BuildTarget.WebGL)
			return $"{hostServerIP}/hotUpdate/WebGL/{gameVersion}";
		else
			return $"{hostServerIP}/hotUpdate/{UnityEditor.EditorUserBuildSettings.activeBuildTarget}/{gameVersion}";
#else
		if (Application.platform == RuntimePlatform.Android)
			return $"{hostServerIP}/hotUpdate/android/{gameVersion}";
		else if (Application.platform == RuntimePlatform.IPhonePlayer)
			return $"{hostServerIP}/hotUpdate/ios/{gameVersion}";
		else if (Application.platform == RuntimePlatform.WebGLPlayer)
			return $"{hostServerIP}/hotUpdate/WebGL/{gameVersion}";
		else
			return $"{hostServerIP}/hotUpdate/StandaloneWindows64/{gameVersion}";
#endif
	}

	/// <summary>
	/// 内置文件查询服务类
	/// </summary>
	private class GameQueryServices : IQueryServices
	{
		public bool QueryStreamingAssets(string fileName)
		{
			// 注意：使用了BetterStreamingAssets插件，使用前需要初始化该插件！
			string buildinFolderName = YooAssets.GetStreamingAssetBuildinFolderName();
			return BetterStreamingAssets.FileExists($"{buildinFolderName}/{fileName}");
		}
	}

	/// <summary>
	/// 资源文件解密服务类
	/// </summary>
	private class GameDecryptionServices : IDecryptionServices
	{
		public ulong LoadFromFileOffset(DecryptFileInfo fileInfo)
		{
			return 32;
		}

		public byte[] LoadFromMemory(DecryptFileInfo fileInfo)
		{
			throw new NotImplementedException();
		}

		public FileStream LoadFromStream(DecryptFileInfo fileInfo)
		{
			BundleStream bundleStream = new BundleStream(fileInfo.FilePath, FileMode.Open);
			return bundleStream;
		}

		public uint GetManagedReadBufferSize()
		{
			return 1024;
		}
	}
}