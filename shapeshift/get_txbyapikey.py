# get_txbyapikey.py
""" 
Name: Jaemin Kim
Email: jaemnkm@gmail.com
Company: sfdev.tech
Date: 8/5/2017
Topic: Shapeshift api
Reference: https://info.shapeshift.io/api
"""

# Comment
"""
Allows vendors to get a list of all transactions that have ever been done using a specific API key. 
Transactions are created with an affilliate PUBLIC KEY, but they are looked up using the linked PRIVATE KEY, to protect the privacy of our affiliates' account details.
"""

# Get List of Transactions with a PRIVATE API KEY

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_tx_by_api_key(api_key):
    
    url_path = "txbyapikey/{}".format(api_key)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get List of Transation with Private API KEY', get_tx_by_api_key('{api key}'))         
