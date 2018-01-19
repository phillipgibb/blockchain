# get_txStat.py
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
Allows vendors to get a list of all transactions that have ever been sent to one of their addresses.
The affilliate's PRIVATE KEY must be provided, and will only return transactions that were sent to output address AND were created using / linked to the affiliate's PUBLIC KEY. 
Please note that if the address is a ripple address and it includes the "?dt=destTagNUM" appended on the end, you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
"""

# Get List of Transactions with a Specific Output Address

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_tx_by_address(address, api_key):
    
    url_path = "txbyapikey/{}/{}".format(address, api_key)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get List of Transactions with a Specific Output Address', get_tx_by_address('{it has to be a shapeshift deposit address}','{api key}'))         
