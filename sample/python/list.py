
# 列表非常适合用于存储在程序运行期间可能变化的数据集。
bicycles = ['trek', 'cannondale', 'redline', 'specialized']
print(bicycles)

# 删除列表数据，根据索引
del bicycles[2]
print(bicycles)

# 弹出
pop_bicycles = bicycles.pop()
print(bicycles)
print(pop_bicycles)

# 按值来删除数据
bicycles.remove("cannondale")
print(bicycles)

# 开头和结尾添加数据
bicycles.insert(1, "wenbin")
bicycles.append("cannondale")
print(bicycles)

#排序
bicycles.sort()
print(bicycles)

# 遍历列表
for bicycle in bicycles:
    print(bicycle)

# 列表解析
squares = [value**2 for value in range(1,11)]
print(squares)

# 双for列表
d_squares = [x * y for x in range(1,11) for y in range(12,14)]
print(d_squares)

# 切片
players=['charles', 'martina', 'michael', 'florence', 'eli']
print(players[0:3])
print(players[3:])

# 循环切片
for player in players[:3]:
    print(player.title())

# 复制列表
my_foods = ['pizza', 'falafel', 'carrot cake']
friend_foods = my_foods[:]

my_foods.append("cannoli")
print("My favorite foods are:")
print(my_foods)

friend_foods.append("ice cream")
print("\nMy friend's favorite foods are:")
print(friend_foods)