# Plutus Api
### Api for [plutus](https://github.com/torbenconto/plutus)

# Documentation
## Quote
| Route           | Type   | Returns                                                                                                 |
|-----------------|--------|---------------------------------------------------------------------------------------------------------|
| /quote/{ticker} | string | Json return of the Quote struct |

## Historical
| Route                                                  | Type | Returns                                                                                                                |
|--------------------------------------------------------| --- |------------------------------------------------------------------------------------------------------------------------|
| /historical/{ticker}?range={range}&interval={interval} | string | Json return of the Historical struct |
Range can be: 1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max
Interval can be: 1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo

## Dividend
| Route                                                | Type | Returns                                                                                                           |
|------------------------------------------------------| --- |-------------------------------------------------------------------------------------------------------------------|
| /dividend/{ticker} | string | Json return of the DividendInfo struct |


# Self Host
## Docker
```bash
docker pull ghcr.io/torbenconto/plutus-api:latest
```
```bash
docker run -p 8080:8080 plutus-api:latest
```

Plutus Api will be running on port 8080 and accessible at http://localhost:8080
