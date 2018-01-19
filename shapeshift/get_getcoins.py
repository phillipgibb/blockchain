# get_getcoins.py
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
Gets list of all the current cryptocurrencies
"""

# Get list of all the cryptocurrencies

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_coins():
    url_path = 'getcoins'
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Current Coins', get_coins())                  
