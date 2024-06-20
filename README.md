# Plutus Api
### Api for [plutus](https://github.com/torbenconto/plutus)

# Documentation
## Quote
| Route           | Type   | Returns                                                                                                 |
|-----------------|--------|---------------------------------------------------------------------------------------------------------|
| /quote/{ticker} | string | Json return of the [Quote](https://github.com/torbenconto/plutus/blob/master/stock/quote.go#L20) struct |

## Historical
| Route                                                  | Type | Returns                                                                                                                |
|--------------------------------------------------------| --- |------------------------------------------------------------------------------------------------------------------------|
| /historical/{ticker}?range={range}&interval={interval} | string | Json return of the [Historical](https://github.com/torbenconto/plutus/blob/master/historical/historical.go#L15) struct |
Range can be: 1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max
Interval can be: 1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo

## Dividend
| Route                                                | Type | Returns                                                                                                           |
|------------------------------------------------------| --- |-------------------------------------------------------------------------------------------------------------------|
| /dividend/{ticker} | string | Json return of the [DividendInfo](https://github.com/torbenconto/plutus/blob/master/stock/dividend.go#L27) struct |


# Self Host
## Docker
```bash
docker pull ghcr.io/torbenconto/plutus-api:latest
```
```bash
docker run -p 8081:8081 plutus-api:latest
```

Plutus Api will be running on port 8081 and accessible at http://localhost:8081