import json
items = json.loads(open("countries.json").read())
d = {}
'''
for item in items:
    for k, v in item["languages"].iteritems():
        if k in d:
            d[k].add(item["name"]["official"])
        else:
            d[k] = set([item["name"]["official"],])

for k, v in d.iteritems():
    print k,len(v)
'''

for item in items:
    try:
        for k, v in item["name"]["native"].iteritems():
            print k,v
    except:
        print json.dumps(item)
