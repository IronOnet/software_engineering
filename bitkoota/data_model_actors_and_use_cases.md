# Modeling  


## Bounded Contexts

    - Trading Context
    - Administration Context
    - Payment Settlement Context
    - Auditing Context


### Trading Context
    - Trades
    - Orders
    - Members
    - KyC_docs
    - Currencies
    - Markets
    - Wallets
    - Account
    - Fills
    - Deposit

### Administration Context

    - Every Models


### Payment Settlement Context

    - Wallets
    - Trades
    - Fills
    - Deposits

### Auditing Context

## Actors

    - Traders
    - Admins
    - Accountants
    - Auditors
    - Compliance Officers

## Data Model


    - Members:
        - member_id
        - first_name
        - last_name
        - email
        - password
        - phone_number
        - is_identified

    - Kyc_docs:
        - member_id: fk
        - doc_type (ID_CARD, PASSPORT)
        - doc_url_path

    - Products
      -product_id
      -product_code

    - Currencies:
        - currency_id
        - currency_code
        - currency_homepage

    - Ticket
        - ticket_id
        - user_id
        - ticket_title
        - ticket_detail

    - Beneficiaries
        - beneficiary_id
        - transaction_id


    - Transaction
         - transaction_id
         - transaction_type

    - Blockchains

    - Wallets
        - wallet_id
        - currency_id
        - user_id
        - balance

    - Markets
        - market_id
        - exchange_rate

    - Accounts  
        - account_id
        - user_id

    - Orders
        - order_id
        - market_id
        - user_id
        - order_status
        - order_type (Market, Limit)
        - order_side
        - quantity  

    - Fill (or Execution)
        - fill_id
        - order_id
        - market_id


    - Trade
        - trade_id
        - order_id


    - Deposit
        - deposit_id
        - currency_id
        - quantity


    - OrderBook
        - orderbook_id
        - orders  
        - market_id


    - Engine
        - engine_id
        - orderbooks

    - PaymentAddress
        - address_id
        - member_id
        - currency_id


    - TradingFee
        - fee_id
        - market_id
        - fee_rate

    - Withdraw

    - Watchlist:
        - user_id
        - market_id

    - Portfolio
