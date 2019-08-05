a = input().split()
for i in range(len(a)):
    a[i] = int(a[i])
max1 = max2 = max3 = min1 = min2 = 1
for i in range(len(a)):
    if(max1 < a[i]):
        max3 = max2
        max2 = max1
        max1 = a[i]
    if(max2 < a[i] and max1 > a[i]):
        max3 = max2
        max2 = a[i]
    if(max3 < a[i] and max2 > a[i]):
        max3 = a[i]
    
    if(min1 > a[i]):
        min2 = min1
        min1 = a[i]            
    if(min1 < a[i] and min2 > a[i]):
        min2 = a[i]

print(max1,max2,max3,min1,min2)        
mul1 = max1 * max2 * max3
mul2 = max1 * min1 * min2
if(mul1 > mul2):
    print(mul1)
else:
    print(mul2)
