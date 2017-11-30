import math
prev=0
for i in range(1000000):
    num = math.log10(2**i)//1
    if prev != num:
        print(i,',',end="")
        prev=num
