# get_limit.py
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
Gets the current deposit limit set by Shapeshift. 
Amounts deposited over this limit will be sent to the return address if one was entered, otherwise the user will need to contact ShapeShift support to retrieve their coins. 
This is an estimate because a sudden market swing could move the limit.
"""



# Get limit between btc and ltc

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_deposit_limit(input_coin, output_coin):
    url_path = "limit/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Deposit Limit', get_deposit_limit('btc', 'ltc'))         
