using System;
using System.IO;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UniFramework.Machine;
using UniFramework.Module;

/// <summary>
/// 流程准备工作
/// </summary>
internal class FsmPatchPrepare : IStateNode
{
	private StateMachine _machine;

	void IStateNode.OnCreate(StateMachine machine)
	{
		_machine = machine;
	}
	void IStateNode.OnEnter()
	{
		// 加载更新面板
		var go = Resources.Load<GameObject>("PatchWindow");
		go = GameObject.Instantiate(go);
		go.transform.parent = GameObject.Find("Canvas").transform;
		go.transform.localScale = new Vector3(1.0f, 1.0f, 1.0f);
		((RectTransform)go.transform).anchorMax = new Vector2(0.5f, 0.5f);
		((RectTransform)go.transform).anchorMin = new Vector2(0.5f, 0.5f);

		_machine.ChangeState<FsmInitialize>();
	}
	void IStateNode.OnUpdate()
	{
	}
	void IStateNode.OnExit()
	{
	}
}