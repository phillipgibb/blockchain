import datetime as datetime

def create_genesis_block():
    # Manually construct a block with
    # index zero and arbitrary previous hash
    return Block(0, data.datetime.now(), "Genesis Block", "0")