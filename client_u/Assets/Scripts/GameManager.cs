using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using ThisGame.Framework;

namespace ThisGame {
    public class GameManager : MonoSingleton<GameManager> {
        // Start is called before the first frame update
        void Start() {
            DontDestroyOnLoad(gameObject);
        }

        // Update is called once per frame
        void Update() {

        }
    }
}
