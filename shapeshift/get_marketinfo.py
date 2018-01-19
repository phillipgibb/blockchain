# get_marketinfo.py
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
This gets the market info (pair, rate, limit, minimum limit, miner fee)
"""




# Get market info between btc and ltc

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_market_info(input_coin, output_coin):
    url_path = "marketinfo/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Market Info', get_market_info('btc', 'ltc'))         
