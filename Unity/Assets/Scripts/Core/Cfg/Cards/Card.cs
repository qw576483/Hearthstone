//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------
using Bright.Serialization;
using System.Collections.Generic;
using SimpleJSON;



namespace Cfg.Cards
{ 

public sealed partial class Card :  Bright.Config.BeanBase 
{
    public Card(JSONNode _json) 
    {
        { if(!_json["id"].IsNumber) { throw new SerializationException(); }  Id = _json["id"]; }
        { if(!_json["imageName"].IsString) { throw new SerializationException(); }  ImageName = _json["imageName"]; }
        { if(!_json["quality"].IsNumber) { throw new SerializationException(); }  Quality = (Cards.ECardQuality)_json["quality"].AsInt; }
        { if(!_json["set"].IsNumber) { throw new SerializationException(); }  Set = (Cards.ECardSet)_json["set"].AsInt; }
        { if(!_json["type"].IsNumber) { throw new SerializationException(); }  Type = (Cards.ECardType)_json["type"].AsInt; }
        { if(!_json["race"].IsNumber) { throw new SerializationException(); }  Race = (Cards.ECardRace)_json["race"].AsInt; }
        { if(!_json["classs"].IsNumber) { throw new SerializationException(); }  Classs = (Cards.ECardClass)_json["classs"].AsInt; }
        { if(!_json["name"].IsString) { throw new SerializationException(); }  Name = _json["name"]; }
        { if(!_json["cnname"].IsString) { throw new SerializationException(); }  Cnname = _json["cnname"]; }
        { if(!_json["cndescription"].IsString) { throw new SerializationException(); }  Cndescription = _json["cndescription"]; }
        { if(!_json["description"].IsString) { throw new SerializationException(); }  Description = _json["description"]; }
        PostInit();
    }

    public Card(int id, string imageName, Cards.ECardQuality quality, Cards.ECardSet set, Cards.ECardType type, Cards.ECardRace race, Cards.ECardClass classs, string name, string cnname, string cndescription, string description ) 
    {
        this.Id = id;
        this.ImageName = imageName;
        this.Quality = quality;
        this.Set = set;
        this.Type = type;
        this.Race = race;
        this.Classs = classs;
        this.Name = name;
        this.Cnname = cnname;
        this.Cndescription = cndescription;
        this.Description = description;
        PostInit();
    }

    public static Card DeserializeCard(JSONNode _json)
    {
        return new Cards.Card(_json);
    }

    /// <summary>
    /// 这是id
    /// </summary>
    public int Id { get; private set; }
    /// <summary>
    /// Texture name
    /// </summary>
    public string ImageName { get; private set; }
    /// <summary>
    /// 品质
    /// </summary>
    public Cards.ECardQuality Quality { get; private set; }
    /// <summary>
    /// 名字
    /// </summary>
    public Cards.ECardSet Set { get; private set; }
    /// <summary>
    /// 类型
    /// </summary>
    public Cards.ECardType Type { get; private set; }
    public Cards.ECardRace Race { get; private set; }
    public Cards.ECardClass Classs { get; private set; }
    /// <summary>
    /// 名称
    /// </summary>
    public string Name { get; private set; }
    /// <summary>
    /// 名称
    /// </summary>
    public string Cnname { get; private set; }
    /// <summary>
    /// 描述
    /// </summary>
    public string Cndescription { get; private set; }
    /// <summary>
    /// 描述
    /// </summary>
    public string Description { get; private set; }

    public const int __ID__ = -801029381;
    public override int GetTypeId() => __ID__;

    public  void Resolve(Dictionary<string, object> _tables)
    {
        PostResolve();
    }

    public  void TranslateText(System.Func<string, string, string> translator)
    {
    }

    public override string ToString()
    {
        return "{ "
        + "Id:" + Id + ","
        + "ImageName:" + ImageName + ","
        + "Quality:" + Quality + ","
        + "Set:" + Set + ","
        + "Type:" + Type + ","
        + "Race:" + Race + ","
        + "Classs:" + Classs + ","
        + "Name:" + Name + ","
        + "Cnname:" + Cnname + ","
        + "Cndescription:" + Cndescription + ","
        + "Description:" + Description + ","
        + "}";
    }
    
    partial void PostInit();
    partial void PostResolve();
}
}
