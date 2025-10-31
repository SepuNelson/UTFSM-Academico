txt = input("Ingrese texto: ")
tr = input("Ingrese significados: ")
f = 0
i = 0
l = 0
txt1 = ""
while i < len(tr):
    if tr[i] == "*":
        emoji = tr[f:i]
        f = i + 1
    elif tr[i] == "$":
        significado = tr[f:i]
        significado = significado.upper()
        f = i + 1
        while l < len(txt):
            if emoji in txt:
                txt = txt.replace(emoji,significado)
            else:
                break

            l += 1
    i += 1
print(txt)
