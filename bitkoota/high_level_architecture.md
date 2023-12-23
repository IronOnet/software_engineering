# Bit Koota Architecture

- Clients
- API Gateway
- Order Management Service
- Order Sequencers (Buy/Sell Queues) 2 for each market
- Matching Engine (One for each Market)
- Price data publishing service
- Price data queue  
- Payment Service (Interact with payment gateways)
- Settlement Service (Settling trades and paying customers)
- Risk Management service.
