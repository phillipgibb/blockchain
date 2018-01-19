# get_rate.py
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
Gets the current rate offered by Shapeshift. 
This is an estimate because the rate can occasionally change rapidly depending on the markets. 
The rate is also a 'use-able' rate not a direct market rate. 
Meaning multiplying your input coin amount times the rate should give you a close approximation of what will be sent out. 
This rate does not include the transaction (miner) fee taken off every transaction.
"""

# Get rate between btc and ltc

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_rate(input_coin, output_coin):
    url_path = "rate/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Rate', get_rate('btc', 'ltc'))                  
