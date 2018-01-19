# shapeshift_overview.py
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
Shapeshift is a service will let you convert between cryptocurrencies. 
I dont know how they do it.
But it might be something similar to mixers, where there is a big pool and people are just putting their coins in and then the service randomly picks other people's coins and gives them different parts of it till they get the amount.
"""

# GET requests
"""
getcoins: Returns the list of available coins.
rate: Returns the exchange rate between two cryptocoins.
limit: Returns the maximum amount of coin you can deposit.
marketinfo: Returns the pair, rate, limit, minimum limit, and miner fee of a particular market.
recenttx: List of the most recent transactions.
txStat: Status of the most recent transaction at an address.
timeremaining: Time remaining to deposit on a fixed transaction.
txbyapikey: Allows vendors to get a list of all transactions that have ever been done using a specific API key.
txbyaddress: Allows vendors to get a list of all transactions that have ever been sent to one of their addresses.
validateAddress: Verifies that an address with work against a particular blockchain.
"""

# Get list of all the cryptocurrencies
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_coins():
    url_path = 'getcoins'
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Current Coins', get_coins())                  
"""

# Get rate between btc and ltc
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_rate(input_coin, output_coin):
    url_path = "rate/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Rate', get_rate('btc', 'ltc'))                  
"""

# Get limit between btc and ltc
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_deposit_limit(input_coin, output_coin):
    url_path = "limit/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Deposit Limit', get_deposit_limit('btc', 'ltc'))         
"""

# Get market info between btc and ltc
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_market_info(input_coin, output_coin):
    url_path = "marketinfo/{}_{}".format(input_coin, output_coin)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('BTC LTC Market Info', get_market_info('btc', 'ltc'))         
"""

# Get recent tx list 
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_recent_tx_list(max_transactions):
    
    url_path = "recenttx/{}".format(max_transactions)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Recent TX List', get_recent_tx_list('{add a number here or leave blank for 5}'))         
"""

# Get status on a shapeshift deposit address
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_recent_tx_list(address):
    
    url_path = "txStat/{}".format(address)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Status', get_recent_tx_list('{it has to be a shapeshift deposit address}'))         
"""

# Get time Remaining on Fixed Amount Transaction
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_time_remaining_on_fixed_tx(address):
    
    url_path = "timeremaining/{}".format(address)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get Time Remaining', get_time_remaining_on_fixed_tx('{it has to be a shapeshift deposit address}'))         
"""

# Get List of Transactions with a PRIVATE API KEY
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_tx_by_api_key(api_key):
    
    url_path = "txbyapikey/{}".format(api_key)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get List of Transation with Private API KEY', get_tx_by_api_key('{api key}'))         
"""

# Get List of Transactions with a Specific Output Address
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def get_tx_by_address(address, api_key):
    
    url_path = "txbyapikey/{}/{}".format(address, api_key)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Get List of Transactions with a Specific Output Address', get_tx_by_address('{it has to be a shapeshift deposit address}','{api key}'))         
"""

# Get Validate an address, given a currency symbol and address.
"""
import requests

BASE_URL = 'https://shapeshift.io/%s'

def validate_address(address, coin_symbol):
    
    url_path = "validateAddress/{}/{}".format(address, coin_symbol)
    url = BASE_URL % url_path
    response = requests.get(url)
    return response.json() 

if __name__ == '__main__':
    print('Validate an address, given a currency symbol and address', validate_address('{it has to be a shapeshift deposit address}','{coin_symbol}'))         
"""
