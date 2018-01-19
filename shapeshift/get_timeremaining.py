# get_timeremaining.py
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
When a transaction is created with a fixed amount requested there is a 10 minute window for the deposit. 
After the 10 minute window if the deposit has not been received the transaction expires and a new one must be created. 
This api call returns how many seconds are left before the transaction expires. 
Please note that if the address is a ripple address, it will include the "?dt=destTagNUM" appended on the end, and you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
"""

# Get time Remaining on Fixed Amount Transaction

import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_time_remaining_on_fixed_tx(address):
    
    url_path = "timeremaining/{}".format(address)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Time Remaining', get_time_remaining_on_fixed_tx('{it has to be a shapeshift deposit address}'))         
