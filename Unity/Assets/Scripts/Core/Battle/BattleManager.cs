using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace UnityGeneralFramework.Common{
    public class BattleManager : MonoSingleton<BattleManager> {
        public BattleController controller;
        public bool isInBattle;

        public void OnBattleBegin() {
            isInBattle = true;
        }
    }
}
