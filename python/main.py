import requests
import os
import json

url = os.environ['TOKENURL']
myobj = {'grant_type': 'client_credentials','client_id':os.environ['CLIENTID'],'client_secret':os.environ['CLIENTSECRET']}
r = requests.post(url, data = myobj)
result = r.json()
print('Token expires in:',result['expires_in'] ,'seconds\n')
print('Token: ',result['access_token'])