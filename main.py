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

a = int(input("곱할 횟수 | "))
b = int(input("곱할 숫자 | "))
print(mup1(a, b))
