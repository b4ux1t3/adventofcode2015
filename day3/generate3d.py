import random, sys

if len(sys.argv)!= 2:
    print "Usage: python generate.py <how many instructions you want>"
    sys.exit()


choices = ("<", ">", "^", "v", "u", "d")
output = ""
for x in range(int(sys.argv[1])):
    output += random.choice(choices)

f = open("randout3d", "w")

f.write(output)

f.close

print "Created an instruction set that is " + sys.argv[1] + " characters long"
