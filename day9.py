with open("input/day9.txt", "r") as f:
    disk_map = [int(x) for x in f.readlines()[0]]

# memory1 = []
# is_file = True
# id_num = 0
# for num in disk_map:
#     if is_file:
#         for _ in range(num):
#             memory1.append(id_num)
#         id_num += 1
#     else:
#         for _ in range(num):
#             memory1.append(None)
#     is_file = not is_file
# memory2 = memory1[:]


# def checksum(mem):
#     result = 0
#     for i, id_num in enumerate(mem):
#         if id_num is not None:
#             result += i * id_num
#     return result


# # part 1
# def defrag1(mem):
#     p_free = 0
#     p_file = len(mem) - 1
#     while True:
#         # push p_free forwards
#         while p_free < p_file and mem[p_free] is not None:
#             p_free += 1
#         # pull p_file back
#         while p_free < p_file and mem[p_file] is None:
#             p_file -= 1
#         if p_free >= p_file:
#             return
#         mem[p_free], mem[p_file] = mem[p_file], mem[p_free]


# defrag1(memory1)
# print("".join([str(x) if x is not None else "." for x in memory1]))
# print(checksum(memory1))


# # part 2
# def defrag2(mem):
#     p_file = len(mem) - 1
#     id_num = max(x for x in mem if x is not None)
#     while id_num >= 0:
#         # pull p_file back
#         while 0 < p_file and mem[p_file] != id_num:
#             p_file -= 1
#         if p_file == 0:
#             return
#         # find size of file
#         p_temp = p_file
#         while mem[p_temp] == id_num:
#             p_temp -= 1
#         file_size = p_file - p_temp
#         # find first free block to fit file
#         for p_free in range(0, p_file - file_size + 1):
#             if all(mem[p_free + k] is None for k in range(file_size)):
#                 # move file
#                 for k in range(file_size):
#                     mem[p_free + k], mem[p_file - k] = mem[p_file - k], mem[p_free + k]
#                 break
#         p_file -= file_size
#         id_num -= 1


# defrag2(memory2)
# print(checksum(memory2))
# print(memory2)

# from aoc import read_input

# disk_map = read_input(split_lines=False)


def checksum(space):
    total = 0
    for i, n in enumerate(space):
        if n == ".":
            continue
        total += int(n) * i
    return total


space = []
free = False
ID = 0
free_space = {}
for c in disk_map:
    n = int(c)
    if free:
        free_space[len(space)] = n
        for _ in range(n):
            space.append(".")
    else:
        for _ in range(n):
            space.append(str(ID))
        ID += 1
    free = not free

space1 = list(space)
for i, c in enumerate(space1):
    if c == ".":
        l = i
        break

# print(space1)
r = len(space1) - 1
while l < r:
    if space1[l] != ".":
        l += 1
        continue
    if space1[r] == ".":
        r -= 1
        continue
    space1[l] = space1[r]
    space1[r] = "."
    l += 1
    r -= 1

print("".join([str(x) if x is not None else "." for x in space1]))
print(checksum(space1))

space2 = list(space)
r = len(space2) - 1
seen = set()
while r > 0:
    if space2[r] == ".":
        r -= 1
        continue
    c = space2[r]
    if c in seen:
        r -= 1
        continue
    seen.add(c)

    file_size = 0
    while space2[r] == c and r >= 0:
        file_size += 1
        r -= 1

    candidates = [(k, v) for (k, v) in free_space.items() if k <= r and v >= file_size]
    if len(candidates) == 0:
        continue

    i, length = min(candidates, key=lambda x: x[0])
    for j in range(i, i + file_size):
        space2[j] = c
    for j in range(r + 1, r + 1 + file_size):
        space2[j] = "."
    del free_space[i]

    if file_size < length:
        free_space[i + file_size] = length - file_size


print(checksum(space2))
