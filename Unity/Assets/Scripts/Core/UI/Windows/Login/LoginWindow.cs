using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using YooAsset;

namespace ThisGame.UI.Window {
    /// <summary>
    /// µÇÂ¼½çÃæ
    /// </summary>
    public class LoginWindow : MonoBehaviour {
        public Button BtnLogin;

        // Start is called before the first frame update
        void Start() {
            BtnLogin.onClick.AddListener(OnLoginClick);
        }

        // Update is called once per frame
        void Update() {

        }

        void OnLoginClick() {
            YooAssets.LoadSceneAsync("dz");
        }
    }
}

