# get_recenttx.py
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
Get a list of the most recent transactions.
"""





# Get recent tx list 

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_recent_tx_list(max_transactions):
    
    url_path = "recenttx/{}".format(max_transactions)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Recent TX List', get_recent_tx_list('{add a number here or leave blank for 5}'))         
