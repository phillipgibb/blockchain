# get_validateAddress.py
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
Allows user to verify that their receiving address is a valid address according to a given wallet daemon. 
If is valid returns true, this address is valid according to the coin daemon indicated by the currency symbol.
"""

# Get Validate an address, given a currency symbol and address.

import requests

BASE_URL = 'https://shapeshift.io/%s'

def validate_address(address, coin_symbol):
    
    url_path = "validateAddress/{}/{}".format(address, coin_symbol)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Validate an address, given a currency symbol and address', validate_address('{it has to be a shapeshift deposit address}','{coin_symbol}'))
    