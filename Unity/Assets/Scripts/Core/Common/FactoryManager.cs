using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace HeartStone.Common {
    public class FactoryManager : MonoBehaviour {
        public GameObject ValetCard;
        public GameObject MagicCard;
        public GameObject WeaponCard;

        /// <summary>
        /// ����һ���̴ӿ�
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateValetCard(int id) {
            return Instantiate(ValetCard);
        }

        /// <summary>
        /// ����һ��ħ����
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateMagicCard(int id) {
            return Instantiate(MagicCard);
        }

        /// <summary>
        /// ����һ��������
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public GameObject CreateWeaponCard(int id) {
            return Instantiate(WeaponCard);
        }
    }
}
