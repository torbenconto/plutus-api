# Plutus Api
### Api for [plutus](https://github.com/torbenconto/plutus)

# Documentation
## Quote
| Route           | Type   | Returns                                                                                                 |
|-----------------|--------|---------------------------------------------------------------------------------------------------------|
| /quote/{ticker} | string | Json return of the [Quote](https://github.com/torbenconto/plutus/blob/master/quote/quote.go#L21) struct |

## Historical
| Route                                   | Type | Returns                                                                                                                |
|-----------------------------------------| --- |------------------------------------------------------------------------------------------------------------------------|
| /historical/{ticker}/{range}/{interval} | string | Json return of the [Historical](https://github.com/torbenconto/plutus/blob/master/historical/historical.go#L15) struct |
Range can be: 1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max
Interval can be: 1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo

## News
| Route         | Type | Returns                                                                                                 |
|---------------| --- |---------------------------------------------------------------------------------------------------------|
| /news/{query} | string | Json return of the [News](https://github.com/torbenconto/plutus/blob/master/news/news.go#L17) struct |
Query should be either a ticker, company name, or other search term.