# -*- coding: UTF-8 -*-
'''
  update: 2019-10-04
'''
import os, re, codecs


# 构建header
with open('./fixed/header.html', 'r') as f:
  txtFront = f.read()

# 构建footer
with open('./fixed/footer.html', 'r') as f:
  txtBack = f.read()

  
filePath = '.'
filelist = os.listdir(filePath)
for file in filelist:
  if 'html' in file:
    print(file)
    # '''
    with open('./'+file,'r') as f:
      tmp = f.read()
    txtAll = txtFront + tmp + txtBack

    # 输出
    with open('../' + file, 'w') as f:# 后打印
      f.write(txtAll)
    # '''


# 结束
print("Processing End.")
