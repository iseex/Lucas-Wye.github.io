# -*- coding: UTF-8 -*-
'''
  update: 2019-10-04
'''
import os, re, codecs

def merge_html(filePath, filelist):
  for file in filelist:
    if 'html' in file:
      print(file)
      # '''
      with open(filePath+"/"+file,'r') as f:
        tmp = f.read()
      txtAll = txtFront + tmp + txtBack

      # 输出
      with open('../' + file, 'w') as f:# 后打印
        f.write(txtAll)
      # '''


# 构建header
with open('./fixed/header.html', 'r') as f:
  txtFront = f.read()

# 构建footer
with open('./fixed/footer.html', 'r') as f:
  txtBack = f.read()

filePathlist = [
  '.', 
  './note',
  # './post'
]

for filePath in filePathlist:
  filelist = os.listdir(filePath)
  merge_html(filePath,filelist)


# 结束
print("Processing End.")
