# Fetch data:
with open('./Day24/input/24.txt', 'r') as file:
    lines = file.readlines()

MIN = 200000000000000
MAX = 400000000000000
LINES = 300

positions = []
vectors = []

for line in lines:
    parts = line.split('@')
    positions.append([int(num) for num in parts[0].split(',')])
    vectors.append([int(num) for num in parts[1].split(',')])

def get_scale_range(index):
    scale1, scale2 = (MIN - positions[index][0]) / vectors[index][0], (MAX - positions[index][0]) / vectors[index][0]
    if scale1 > scale2: scale1, scale2 = scale2, scale1
    if scale2 < 0: return None
    if scale1 < 0: scale1 = 0
    scale3, scale4 = (MIN - positions[index][1]) / vectors[index][1], (MAX - positions[index][1]) / vectors[index][1]
    if scale3 > scale4: scale3, scale4 = scale4, scale3
    if scale4 < 0: return None
    if scale3 < 0: scale3 = 0
    if scale3 > scale2 or scale1 > scale4: return None
    return (max(scale1, scale3), min(scale2, scale4))

ranges = []

for i in range(LINES): ranges.append(get_scale_range(i))

count = 0

# Test every ray against every other ray
for i in range(LINES):
    for j in range(1 + i, LINES):
        if not ranges[i] or not ranges[j]: continue
        point_i_1 = (positions[i][0] + vectors[i][0] * ranges[i][0], positions[i][1] + vectors[i][1] * ranges[i][0])
        point_i_2 = (positions[i][0] + vectors[i][0] * ranges[i][1], positions[i][1] + vectors[i][1] * ranges[i][1])
        point_j_1 = (positions[j][0] + vectors[j][0] * ranges[j][0], positions[j][1] + vectors[j][1] * ranges[j][0])
        point_j_2 = (positions[j][0] + vectors[j][0] * ranges[j][1], positions[j][1] + vectors[j][1] * ranges[j][1])
        # make sure the 1 point has the smaller x value
        if point_i_1[0] > point_i_2[0]: point_i_1, point_i_2 = point_i_2, point_i_1
        if point_j_1[0] > point_j_2[0]: point_j_1, point_j_2 = point_j_2, point_j_1
        if point_i_2[0] < point_j_1[0] or point_i_1[0] > point_j_2[0]: continue
        # get slopes of each line
        slope_i = (point_i_2[1] - point_i_1[1]) / (point_i_2[0] - point_i_1[0])
        slope_j = (point_j_2[1] - point_j_1[1]) / (point_j_2[0] - point_j_1[0])
        # choose the larger x value for the [0] points
        if point_i_1[0] > point_j_1[0]: point_j_1 = (point_i_1[0], point_j_1[1] + (point_i_1[0] - point_j_1[0]) * slope_j)
        elif point_j_1[0] > point_i_1[0]: point_i_1 = (point_j_1[0], point_i_1[1] + (point_j_1[0] - point_i_1[0]) * slope_i)
        # choose the smaller x value for the [1] points
        if point_i_2[0] < point_j_2[0]: point_j_2 = (point_i_2[0], point_j_2[1] + (point_i_2[0] - point_j_2[0]) * slope_j)
        elif point_j_2[0] < point_i_2[0]: point_i_2 = (point_j_2[0], point_i_2[1] + (point_j_2[0] - point_i_2[0]) * slope_i)
        # compare y values of points
        if point_i_1[1] > point_j_1[1] and point_i_2[1] > point_j_2[1]: continue
        elif point_j_1[1] > point_i_1[1] and point_j_2[1] > point_i_2[1]: continue
        # Within the range that the two rays share an x value, which ray has 
        # the larger y value must change (this is how we know they intersect)
        else: count += 1

print(count)