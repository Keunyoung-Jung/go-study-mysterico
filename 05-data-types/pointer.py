import copy

x = [1,2]
y = x
z = copy.deepcopy(x)

print(id(x),":",x)
print(id(y),":",y)
print(id(z),":",z)

x.append(3)
print("---x.append(3) 선언---")

print(id(x),":",x)
print(id(y),":",y)
print(id(z),":",z)