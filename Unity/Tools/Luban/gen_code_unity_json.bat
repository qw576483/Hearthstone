set CONF_ROOT=..\..

set GEN_CLIENT=Luban.ClientServer\Luban.ClientServer.exe

%GEN_CLIENT% -j cfg --^
 -d %CONF_ROOT%\..\Common\Excel\Defines\__root__.xml ^
 --input_data_dir %CONF_ROOT%\..\Common\Excel\Datas ^
 --output_code_dir ..\..\Assets\Scripts\Core\Cfg\ ^
 --output_data_dir ..\..\Assets\HURes\Cfg\ ^
 --gen_types code_cs_unity_json,data_json ^
 -s all 

pause