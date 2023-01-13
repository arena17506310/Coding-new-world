import os

currentPath = os.getcwd()

class Nothing:
        hp = 99999
        atk = 9999999
        att = "🌱"
        chr = open(f"{currentPath}\\monsters\\worm.txt","rt", encoding='UTF8').read()

class Worm:
        hp = 20
        atk = 10
        att = "🌱"
        chr = open(f"{currentPath}\\monsters\\worm.txt","rt", encoding='UTF8').read()