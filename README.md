[![Build Status](https://travis-ci.org/blockassets/cgminer_client.svg?branch=master)](https://travis-ci.org/blockassets/cgminer_client)

# cgminer golang client

Just enough of the client to get some work done. PR's welcome.

Thanks to [HyperBitShop.io](https://hyperbitshop.io) for sponsoring this project.

### Usage:

``
	client := cgminer_client.New("10.0.0.1", 4028, 5)
``

### Build

Install [dep](https://github.com/golang/dep) and the dependencies...

`make dep`
`make`