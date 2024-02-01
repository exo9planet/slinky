# Websocket Providers

## Overview

Websocket providers utilize websocket APIs / clients to retrieve data from external sources. The data is then transformed into a common format and aggregated across multiple providers. To implement a new provider, please read over the base provider documentation in [`providers/base/README.md`](../base/README.md).

Websockets are preferred over REST APIs for real-time data as they only require a single connection to the server, whereas HTTP APIs require a new connection for each request. This makes websockets more efficient for real-time data. Additionally, web sockets typically have lower latency than HTTP APIs, which is important for real-time data.

## Supported Providers

The current set of supported providers are:

* [BitFinex](./bitfinex/README.md) - BitFinex is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. BitFinex is a **primary data source** for the oracle.
* [Bitstamp](./bitstamp/README.md) - Bitstamp is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Bitstamp is a **primary data source** for the oracle.
* [ByBit](./bybit/README.md) - ByBit is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. ByBit is a **primary data source** for the oracle.
* [Coinbase](./coinbase/README.md) - Coinbase is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Coinbase is a **primary data source** for the oracle.
* [Crypto.com](./cryptodotcom/README.md) - Crypto.com is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Crypto.com is a **primary data source** for the oracle.   
* [Gate](./gate/README.md) - Gate.io is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Gate.io is a **primary data source** for the oracle.
* [Kraken](./kraken/README.md) - Kraken is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Kraken is a **primary data source** for the oracle.
* [Huobi](./huobi/README.md) - Huobi is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. Huobi is a **primary data source** for the oracle.
* [KuCoin](./kucoin/README.md) - KuCoin is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. KuCoin is a **primary data source** for the oracle.
* [MEXC](./mexc/README.md) - MEXC is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. MEXC is a **primary data source** for the oracle.
* [OKX](./okx/README.md) - OKX is a cryptocurrency exchange that provides a free API for fetching cryptocurrency data. OKX is a **primary data source** for the oracle.