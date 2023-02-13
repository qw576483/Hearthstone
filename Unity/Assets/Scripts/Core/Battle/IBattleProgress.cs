using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Battle {
    /// <summary>
    /// 战斗流程控制
    /// </summary>
    public interface IBattleProgress {
        void OnPEnter();
        void OnPUpdate();
        void OnPExit();
    }
}

