import json
items = json.loads(open("countries.json").read())
d = {}
for item in items:
    for k, v in item["languages"].iteritems():
        if k in d:
            d[k] += 1
        else:
            d[k] = 1

for k, v in d.iteritems():
    print k,v
