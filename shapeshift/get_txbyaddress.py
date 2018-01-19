# get_txbyaddress.py
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
This returns the status of the most recent deposit transaction to the address.
"""

# Get status on a shapeshift deposit address

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_recent_tx_list(address):
    
    url_path = "txStat/{}".format(address)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Status', get_recent_tx_list('{it has to be a shapeshift deposit address}'))         
