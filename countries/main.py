import json
items = json.loads(open("countries.json").read())

country2timezone = json.loads(open("countries_with_timeZones.json").read())
m = {}
for item in country2timezone:
    m[item["IsoAlpha3"]] = item

d = {}
for item in items:
    for k, v in item["name"]["native"].iteritems():
        if k not in d:
            d[k] = [item["cca3"],]
        else:
            d[k].append(item["cca3"])
        print k
        try:
            print m[item["cca3"]]["TimeZones"]
        except:
            pass
for k, v in d.iteritems():
    #print k, v
    pass


