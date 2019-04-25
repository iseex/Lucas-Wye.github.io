#coding=utf-8
'''
  author: lustre
  time: 2019-03-24
  function: this code is used to read the RGBA data from a binary file and then plot the image
'''

import sys
import struct
from PIL import Image

if len(sys.argv)!=2:
  print("Usage: python read_from_bin.py in_binary_file")

infile=open(sys.argv[1],"rb")

w = int(struct.unpack("i",infile.read(4))[0])
h = int(struct.unpack("i",infile.read(4))[0])

new_img = Image.new('RGBA',(w,h))
for i in range(h):
  for j in range(w):
    new_img.putpixel((j,i),(int(struct.unpack("i",infile.read(4))[0]),int(struct.unpack("i",infile.read(4))[0]),int(struct.unpack("i",infile.read(4))[0]),int(struct.unpack("i",infile.read(4))[0])))

new_img.show()
