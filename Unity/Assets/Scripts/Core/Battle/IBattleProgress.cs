using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Battle {
    /// <summary>
    /// ս�����̿���
    /// </summary>
    public interface IBattleProgress {
        void OnPEnter();
        void OnPUpdate();
        void OnPExit();
    }
}

