using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Common {
    public class FactoryManager : MonoBehaviour {
        public GameObject ValetCard;
        public GameObject MagicCard;
        public GameObject WeaponCard;

        /// <summary>
        /// 创建一张侍从卡
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateValetCard(int id) {
            return Instantiate(ValetCard);
        }

        /// <summary>
        /// 创建一张魔法卡
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateMagicCard(int id) {
            return Instantiate(MagicCard);
        }

        /// <summary>
        /// 创建一张武器卡
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateWeaponCard(int id) {
            return Instantiate(WeaponCard);
        }
    }
}
