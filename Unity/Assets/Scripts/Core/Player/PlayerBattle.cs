using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Player {
    /// <summary>
    /// Player
    /// </summary>
    public class PlayerBattle : MonoBehaviour {
        private PlayerData playerData;
        private PlayerBattleData battleData;

        public void OnReset(PlayerBattleData data) {
            this.battleData = data;
        }
    }
}

