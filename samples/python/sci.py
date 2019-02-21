import numpy as np
import pandas as pd
from pandas import Series, DataFrame


obj = Series([2,4,7,8,9,3])
"""obj=
0    2
1    4
2    7
3    8
4    9
5    3
dtype: int64
"""
print(obj.values)    #return:array([2, 4, 7, 8, 9, 3], dtype=int64)
print(obj.index)     #return: RangeIndex(start=0, stop=6, step=1)

#构建DataFrame
#最常用的一种是直接传入一个由等长列表或numpy数组组成的字典：
data={'names':['Bob','Jane','Jack','Ann'],
      'sex':['M','F','M','F'],
      'age':[21,30,26,28]}
frame=DataFrame(data)
"""没有指定索引，会自动加上索引，且全部列会被有序排列
frame=
   age names sex
0   21   Bob   M
1   30  Jane   F
2   26  Jack   M
3   28   Ann   F
"""
#若指定列序列，则会按照指定顺序排列
frame=DataFrame(data,columns=['names','sex','age'])
"""frame=
  names sex  age
0   Bob   M   21
1  Jane   F   30
2  Jack   M   26
3   Ann   F   28
"""
print(frame)
