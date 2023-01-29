public class AOTGenericReferences : UnityEngine.MonoBehaviour
{

	// {{ constraint implement type
	// }} 

	// {{ AOT generic type
	//System.Collections.Generic.Dictionary`2<System.Object,System.Object>
	//System.Collections.Generic.Dictionary`2<System.Int32,System.Object>
	//System.Collections.Generic.IEnumerable`1<System.Object>
	//System.Collections.Generic.IEnumerator`1<System.Object>
	//System.Collections.Generic.List`1<System.Object>
	//System.Collections.Generic.List`1/Enumerator<System.Object>
	//System.Func`2<System.Object,System.Object>
	//System.Nullable`1<System.Int32>
	// }}

	public void RefMethods()
	{
		// System.String Bright.Common.StringUtil::CollectionToString<System.Object>(System.Collections.Generic.IEnumerable`1<System.Object>)
	}
}