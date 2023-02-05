# Description 

- Bitkoota is a fintech platform that allows it's users to trade, invest in various 
types of crypto currency instruments 



## Software Requirements : Crypto exchange scenario


### Functional requirements

- Users can register themselves with their email 
- Users can login or register themselved with google or facebook 
- Users can login or register themselves with two factor authentication
- Users can trade up to a 100 instruments on the platform 
- Users can use the platform as a custodian of their funds 
- The platform matches buyers and sellers of crypto currency coins
- Users can get to see real time prices of various instruments traded on the platform 
- Users can swap their crypto currency with other listed assest or cash 
- Users can buy crypto-currencies with their bank card, credit card, and mobile money account 
- Users can make fund deposits in their account in various currency denominations 
- Users can make limit orders and market orders on the exchange
- Users can cancel orders  
- Users can receive matched trades in real-time 
- Users can view the realtime limit order book 
-  
- Users can have a platform native wallet for their funds (optional) 
- Users must send their credential for transactions bigger than 10.000 USD a month 
- Users can trade crypto stocks
- Users can have wallets on the exchange
- Users can access the platform on a mobile app. (both Android & IOS)
- The web version of the platform should be responsive to allow different types of users to access the platform as a risk mitigation strategy for potential bans
- The platforms check whether users have sufficient funds before trading takes place
- Thex exchange nees to support thousands of users trading at the same time 
- The exchange should be able to support billions of orders per day. 


### Non functional requirements 

- The exchange needs to support thousands of users trading at the same time 

- The exchange should be able to support billions of orders per day. 

- Availability: At least 99.9% availability is crucial, 

- Fault-tolerance: Fault tolerance and fast recovery mechanisms are needed to limit the 
impact of production incidents

- Latency: The round-trip latency should be at the millisecond level, with a particular focus 
on the 99th percentile latency. The roundtrip latency is measured from the moment the order enters the exchange to the point where the market order returns as a filled execution

- Security: The exchange should have an account management system. For legal and compliance
the exchange performs a KYC (know your customer) check to verify a user's identity before 
a new account is opened. For public resources, we should prevent DDOS attacks. 

