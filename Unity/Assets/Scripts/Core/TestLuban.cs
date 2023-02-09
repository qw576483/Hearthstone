using SimpleJSON;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;
using System;

using cfg.item;

public class TestLuban : MonoBehaviour
{
    private void Awake() {
        var tables = new cfg.Tables(file =>
                JSON.Parse(File.ReadAllText("Assets/HFRes/Config/json" + "/" + file + ".json")
        ));

        cfg.item.Item itemInfo = tables.TbItem.Get(10000);
        //Debug.Log(" 1111 " + itemInfo.ToString() );
    }


    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
