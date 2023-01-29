#!/bin/zsh

GEN_CLIENT=Luban.ClientServer/Luban.ClientServer.dll
CONF_ROOT=..

dotnet ${GEN_CLIENT} -j cfg --\
 -d ${CONF_ROOT}/Excel/Defines/__root__.xml \
 --input_data_dir ${CONF_ROOT}/Excel/Datas \
 --output_code_dir ../Assets/HFRes/Config/Gen \
 --output_data_dir ../Assets/HFRes/Config/json \
 --gen_types code_cs_bin,data_bin \
 --external:selectors unity_cs \
 -s all 
