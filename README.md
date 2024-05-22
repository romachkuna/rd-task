# Currency Conversion API

The project uses [third-party api](https://freecurrencyapi.com/) for currency conversion, register to acquire API KEY

## Get Started

1. run  `go mod tidy` to install the dependencies
2. configure `.env` file to match your postgres database. for
   example `dns = "host=localhost user=XXXX password=XXXX dbname=XXXX port=5432 sslmode=disable TimeZone=Etc/GMT+4"`
3. set your api key in `.env` file. `API_KEY="XXXX"`
4. run `go run . ` from root directory of the project

**For documentation of the api find out more [here](https://documenter.getpostman.com/view/22433649/2sA3QpBDCe)**