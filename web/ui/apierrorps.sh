#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
# echo $SCRIPT_DIR
cd src/openapi


# 遍历所有文件，检查是否存在 @ts-nocheck
for file in $(find . -name "*.ts")
do
  if grep -q "@ts-nocheck" $file; then
    echo "File $file exists and has @ts-nocheck"
  else
    # 在文件开头添加 @ts-nocheck
    sed -i '1s;^;// @ts-nocheck\n;' $file
  fi
done
