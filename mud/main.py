from colorama import init as colorama_init
from colorama import Fore
from colorama import Style
from monster import *
from importlib import reload

mons_list = {
    "worm": Worm,
    "": Nothing
}


skills = []

colorama_init()
name = "worm"

hp = mons_list[name].hp
atk = mons_list[name].atk
att = mons_list[name].att
chr = mons_list[name].chr

def check_letter(hp):
    l = str(input("할 동작의 문자를 입력하시오) "))
    if l == "A" or l == "a":
        hp -= YOU.atk
    elif l == "I" or l == "i":
        pass
    elif l == "S" | l == "s":
        pass
    elif  l == "R" | l == "r":
        breakpoint
    YOU.hp -= atk
    return hp
class YOU:
    hp = 30
    atk = 10



while hp > 0:
    
    print(f"""














Info) ❤️:{hp} ⚔️:{atk} {att}
┌─                ─┐
{Fore.MAGENTA}
{chr}
{Fore.RESET}
└─                ─┘""")
    print(f"""┌──────────────────┐
 ❤️:{YOU.hp} ⚔️:{YOU.atk}
│Attack[A]  Item[I]│
│Skill[S]    Run[R]│
└──────────────────┘
    """)

    check_letter()


