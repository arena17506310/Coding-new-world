def sum(a, b):
    return a + b

def mup(a, b):
    for i in range(a):
        b = sum(b, b)
    return b
def mup1(a, b):
    for i in range(a):
      b = mup(b , b)
    return b

print(mup1(3,2))