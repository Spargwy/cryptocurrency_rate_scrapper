# cryptocurrency scrapper

Service for collecting actual cryptocurrency rate from https://min-api.cryptocompare.com/data/pricemultifull.
Response structure is not equal cryptocompare.com response and has the following json structure:
```
{
  "RAW": {
    "BTC": {
      "USD": {
        "CHANGE24HOUR": 353.1999999999971,
        "CHANGEPCT24HOUR": 0.6465657692118947,
        "OPEN24HOUR": 54627.08,
        "VOLUME24HOUR": 15632.196644309997,
        "VOLUME24HOURTO": 849521265.3776581,
        "LOW24HOUR": 53346.3,
        "HIGH24HOUR": 55193.65,
        "PRICE": 54980.28,
        "LASTUPDATE": 1638133426,
        "SUPPLY": 18873906,
        "MKTCAP": 1037692636573.68
      }
    }
  },
  "DISPLAY": {
    "BTC": {
      "USD": {
        "CHANGE24HOUR": "$ 353.20",
        "CHANGEPCT24HOUR": "0.65",
        "OPEN24HOUR": "$ 54,627.1",
        "VOLUME24HOUR": "Ƀ 15,632.2",
        "VOLUME24HOURTO": "$ 849,521,265.4",
        "HIGH24HOUR": "$ 55,193.7",
        "PRICE": "$ 54,980.3",
        "FROMSYMBOL": "Ƀ",
        "TOSYMBOL": "$",
        "LASTUPDATE": "Just now",
        "SUPPLY": "Ƀ 18,873,906.0",
        "MKTCAP": "$ 1,037.69 B"
      }
    }
  }
}
```
Also, data periodically, by crontab scheduler collects from resource above and store it to MySQL database. Therefore, if resource is not available, data will be extract from database.

A lot of parameters such as
- DB params/Docker db params
- Crontab
- App port
- Available fsyms/tsyms parameters to request

is configurable from .env file.

Application is dockerized.

# Endpoints
- /price 

Primary endpoint that will collect data from https://min-api.cryptocompare.com/data/pricemultifull with specified parameters.
For example ```/price?fsyms=BTC&tsyms=USD``` will return response above.

# Pre-requests
- golang 1.17 version https://go.dev/doc/install
- mysql database https://dev.mysql.com/

- docker and docker-compose(only if running with docker)


# Run
- Execute ```git clone https://github.com/Spargwy/cryptocurrency_rate_scrapper```
and ```cd cryptocurrency_rate_scrapper```


- add .env file that will contains data for configurable parameters. .env.example contains all necessery fields that you need to copy to your .env file and replace with your own

## Standart way
- Execute ```go build ```
- Execute ```./cryptocurrency_rate_scrapper```

## Via Docker
- Execute ```make docker-run``` command
- Note: As default, via makefile, app is running as daemon

## Arguments
- If you want to run application as daemon process, you can add `-d` argument when calling binary `./cryptocurrency_rate_scrapper -d`