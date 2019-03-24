#coding=utf-8
'''

  author: Zewen-Ye
  time: 2019-03-24
  function: this code is used to read an image and save RGB data to outfile

'''

import sys
import struct
from PIL import Image

if len(sys.argv)!=3:
  print("Usage: python read_image.py image_file outfile")
  exit()  

img = Image.open(sys.argv[1])
print("The input image mode is: ") # RGBA
print(img.mode)
print("The input image size is: ") # 2048,1024
w,h = img.size
print(w,h)

# out to file
outfile = open(sys.argv[2],'wb')
outfile.write(struct.pack('i',w))
outfile.write(struct.pack('i',h))

for i in range(h):# 1024
  for j in range(w):# 2048
    r,g,b,a = img.getpixel((j,i))
    outfile.write(struct.pack('i',r))    
    outfile.write(struct.pack('i',g))
    outfile.write(struct.pack('i',b))
    outfile.write(struct.pack('i',a))


img.close()
outfile.close()

