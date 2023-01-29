set CONF_ROOT=..\..

set GEN_CLIENT=Luban.ClientServer\Luban.ClientServer.exe
set OUTPUT_DIR = %CONF_ROOT%\Assets\HFRes\Config

%GEN_CLIENT% -j cfg --^
 -d %CONF_ROOT%\..\Common\Excel\Defines\__root__.xml ^
 --input_data_dir %CONF_ROOT%\..\Common\Excel\Datas ^
 --output_code_dir ..\..\Assets\Scripts\HotUpdate\Config\ ^
 --output_data_dir ..\..\Assets\HURes\Config\ ^
 --gen_types code_cs_unity_json,data_json ^
 -s all 

pause