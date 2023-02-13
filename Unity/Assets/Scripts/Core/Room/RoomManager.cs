using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone {
    /// <summary>
    /// �������
    /// </summary>
    public class RoomManager : MonoBehaviour {
        private static RoomManager _instanceRoom;
        public static RoomManager InstanceRoom {
            get { return _instanceRoom; }
        }

        void Awake() {
            if (_instanceRoom) {
                Destroy(gameObject);
                return;
            }

            _instanceRoom = this;
        }

        // TODO:����״̬����״̬����
        public void RoomInit() {}
        public void RoomReset() {}
        public void RoomClose() {}

    }
}
