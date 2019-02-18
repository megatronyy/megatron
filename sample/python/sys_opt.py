"""
python处理文本文件数据
"""

#! /usr/bin/python3
#

import linecache
import sys


def handleTxt(filepath, line_number=0):
    theline = linecache.getline(filepath, line_number)
    if theline is None and len(theline) > 0:
        print(theline)
    else:
        print("文件不能为空")


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("请输入相应参数")
        sys.exit()

    filepath = sys.argv[1]
    line_number = sys.argv[2]

    handleTxt(filepath, line_number)
