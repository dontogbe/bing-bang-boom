# bing-bang-boom
Assign Speakers to domains

## Python client script
```azure
import argparse
import os
import sys
import json
from requests import Request, Session

parser = argparse.ArgumentParser(
    prog='bigbangboom client', description="Get Speaker domain", epilog='A new era of audio! :)')
parser.add_argument('id', type=int, help='the id of the speaker')
args = parser.parse_args()

speaker_id = args.id

s = Session()
json_data = json.dumps({"id": speaker_id})
req = Request('POST', 'http://localhost:8080/map', data=json_data)
prepped = req.prepare()
resp = s.send(prepped)
domain = resp.json()['domain']
if domain is None:
    print('Speaker role is Meh')
elif len(domain) > 1:
    print('Speaker role is Boom')
elif len(domain) == 1 and domain[0] == 'A':
    print('Speaker role is Bing')
elif len(domain) == 1 and domain[0] == 'B':
    print('Speaker role is Bang')
```

