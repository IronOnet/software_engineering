# Entities, Value Objects, and Aggregates 

Entities and value objects is where we write most of the business logic for 
a domain driven application. 

Aggregates are usedful when we need to cluster domain objects together and treat them 
as a single item. 

## Entities 

In domain-driven design, entities are defined by their identity. Their attributes do not 
define them. and it is expected that although their attributes may change over time, 
their identity will not. While the entity may change so much that it is indistinguishable from 
where it started, it retains the same identity, and we treat it as the same object.

Lets look at an exaample. On ebay.com you can sign up as user. If you choose to sell something 
you become a seller, you can also choose to bid on items.

```UML 

Auctions: 
    id: int 
    starting_price: double 
    seller_id: int 
    created_at: timestamp 
    auction_start: timestamp 
    auction_end: timestamp 

Bids: 
    id: int 
    bid_amount: double 
    user_id: int 
    created_at: timestamp 

Users: 
    id: int 
    first_name: varchar 
    last_name: varchar 
    email: varchar 
    address: varchar 
    created_at: timestamp 
``` 

Some actions that could take place in our system are as follows: 

    - A user updates their address 
    - A user updates their email address 
    - An auction end time is updated 

These actions do not change the identity of our entity. We are still referencing the 
same ID, but some attributes may have changed. 

An implementation of the auction entity might look as follows 

```go 

package chap3 

import (
    "time" 
    "github.com/Rhymond/go-money"
)

// Auction is an entity to represent our auction construct 
type Auction struct{
    ID      int 
    // We use a specific money lib as floats are not good 
    // ways to represent money. 
    startingPrice   money.Money 
    sellerID        int 
    auctionStart            time.Time 
    auctionEnd              time.Time
    createdAt               time.Time 

}

type User struct{
    ID          int 
    FirstName           string 
    LastName            string 
    Email               string 
    Address             string 
    CreatedAt           time.Time 
}

type Bid struct{
    ID          int 
    Amount      float64 
    UserID      int 
    CreatedAt   time.Time 
}

``` 

using UUIDs to future proof your entities, especially for systems of substantial scale 

```go 

type SomeEntity struct{
    id uuid.UUID 
}

func NewSomeEntity() *SomeEntity{
    id := uuid.New() 
    return &SomeEntity{ id: id} 
}
``` 
#### A warning when defining entities 

Due to the focus of entities being on their identity, it is very easy to fall into the trap of letting 
the database design dictate what your domain model will look like. This can lead to what is known as 
an anemic model.

Anemic models have little or no domain behavior as part of their design. This means that you are not 
getting the full benefit of DDD.


```go 

package anemic 

import (
    "time" 
    "github.com/Rhymond/go-money"
)

type AnemicAuction struct{
    id      int 
    startingPrice   money.Money 
    sellerID        int 
    createdAt       time.Time
    auctionStart    time.Time 
    auctionEnd      time.Time
}

func (a *AnemicAuction) GetID() int{
    return a.id 
}

func (a *AnemicAuction) GetSellerID() int{
    return a.sellerID
}

func (a *AnemicAuction) StartingPrice() money.Money{
    return a.startingPrice 
}

func (a *AnemicAuction) SetStartingPrice(startingPrice money.Money){
    a.sartingPrice = startingPrice 
}

func (a *AnemicAuction) SetSellerId(id int){
    a.sellerId = id 
}

``` 

A refactored anemic entity would look like this

```go 

package non_anemic 

import (
    "errors" 
    "time" 

    "github.com/Rhymond/go-money"
)

type AuctionRefactored struct{
    id      int 
    startingPrice       money.Money 
    sellerID            int 
    createdAt           time.Time 
    auctionStart        time.Time 
    auctionEnd          time.Time 
}

func (a *AuctionRefactored) GetAuctionElapsedDuration() time.Duration{
    return a.auctionStart.Sub(a.auctionEnd)
}

func (a *AuctionRefactored) GetAuctionEndTimeInUTC() time.Time{
    return a.auctionEnd 
}

func (a *AuctionRefactored) SetAuctionEnd(auctionEnd time.Time) error {
    if err := a.validateTimeZone(auctionEnd); err != nil{
        return err 
    }
    a.auctionEnd = auctionEnd 
    return nil 
}

func (a *AuctionRefactored) validateTimeZone(t time.Time) error{
    tz, _ := t.Zone() 
    if tz != time.UTC.String(){
        return errors.New("time zone must be UTC")
    }
    return nil 
}

``` 
Even in our simple example, we can see the benefit of our entity having some business logic. 


#### A note on object-relational mapping 

ORMs are a popular approach to database persistence, but they are not a DDD concept. 
By using an ORM, you are delegating control of query creatin and planning. 

If you want to use an ORM, ensure it does not control how you write your entities in your DDD 
context, otherwise, you may end up with an anemic model. We also want to keep the coupling 
between our entity and ORM to a minimum. Therefore, it is recomended to use an adaptor layer 
to decouple your ORM and DDD entity layer

### Working with value objects 

Value, objects are in some ways the opposite of entities. With value objects, we want to assert that 
two objects are the same given their values. Value objects do not have identities and are often used 
in conjunction with entities and aggregates to enable us to build a rich model of our domain. 
We typically use them to measure, quantify, or describe something about our domain. 


```go 

package point 

type Point struct{
    x       int 
    y       int 
}

func NewPoint(x, y int) *Point{
    return &Point{
        x: x, 
        y: y, 
    }
}

``` 

We write the following test to check if two points witht the same 
coordinates are equal. 

```go 
package point_test 

import (
    "testing" 
    "ddd-golang/point"
)

func Test_Point(t *testing.T){
    a := point.Newpoint(1, 1) 
    b := point.NewPoint(1, 1) 
    if a != b{
        t.Fatal("a and be were not equal")
    }
}

```

To a human, these two points are equal, but in Go this will fail 
because when we use the & symbol, we create a pointer to a memory adddress 
where points A and B are stored.

if we change the point definition to the following 

```go 

type Point struct{
    y int 
    x int 
}

func NewPoint(x, y int) Point{
    return Point{
        x: x, 
        y: y, 
    }
}
``` 

If we run a test with this implementation, the test passes, this is because 
the two points are now beign compared on their values when we do an equality 
check. They are value objects, we can treat them equaly if their values are 
equal.  

Value objects should be replaceable.

#### How shold one decide whether to use an entity or value object? 

We should aim to use value objects as much as possible when modeling our domain. 
This is because they are the safest constructis we can use when implemented 
correctly. We do not have to worry about consumers incorreclty modifying our 
instance in a way we did not intend. 

If you care only about the values of an object, then it should preferably be a value object. 
Some other questions to ask yourself to ensure a value object is the right choice for you are: 

    - Is it possible for me to treat this object as immutable? 
    - Does it measure, quantify, or describe a domain concept? 
    - Can it be compared to other objects of the same type by its values?


### The aggregate Pattern 

In domain-driven design, the aggregate pattern refers to a group of domain objects that can 
be treated as one for some behaviors. Some example of aggregate patterns are : 

    - An order: Typically an order consists of individual items, but it is helpful to treat them 
    as a single thing (an order) for some purposes within our system. 

    - A team: A team consists of many employees. In our system, we should likely have a domain object for employees, but grouping them and applying behaviors to them as a team would be helpful in situations such as organizing departments. 

    - A Wallet: Typically a wallet (even a virtual one) contains many cards and potential currencies for many countries and maybe cryptocurrencies. We may want to track the value of the wallet over time and to do that, we may treat the wallet as an aggregate.

Aggregates are not to be confused with data structures used for collections of data, such as arrays, maps and slices. While an aggregate may use these collections, an aggregate is a DDD concept and therefore, will usually contain multiple collections, fields, functions, and methods.

The job of an aggregate pattern is to act as a transaction boundary for the domain objects within. 
Loading serving, editing and deleting should happen to all objects within the aggregate or not at all.
Let's illustrate with an example. 

    - If an order is cancelled, we should return all items within that order to stock. We may also want to trigger a refund. 
    - If a new employee joins a team, we may need to update the line manager structure. 

    - If a user adds a new card to their wallet, we need to ensure its balance is reflected in the total wallet balance.

Let's look at how we might implement the wallet aggregate we mentioned here: 

```go 

type WalletItem interface{
    GetBalance() (money.Money, error)
}

type Wallet struct{
    id uuid.UUID 
    ownerID uuid.UUID 
    walletItems []WalletItem 
}

func (w Wallet) GetWalletBalance() (*money.Money, error){
    var bal money.Money 
    for _, v := range w.walletItems{
        itemBal, err := v.GetBalance() 
        if err != nil{
            return nil, errors.New("failed to get balance")
        }
        bal, err = bal.Add(&itemBal) 
        if err != nil{
            return nil, errors.New("failed to increment balance")
        }
    }
    return bal, nil 
}

```


#### Discovering aggregates 

Before trying to cluster our domain models into aggregates, we need to find our bounded 
context's invariants. An invariant is simply a rule in our domain that must always be true. 
For example, we may say that in our system, for an order to be created, we must have 
the item in stock. this is a business invariant. 

#### Designing aggregates 

Generally, we should aim for small aggregates. Keeping aggregates small willhelp make 
our system more scalable, improve performance, and give transactions more chances of success.

```go 

type item struct{
    name string 
}

type Order struct{
    items       []item 
    taxAmount   money.Money 
    discount    money.Money 
    paymentCard uuid.UUID 
    customerID  uuid.UUID 
    markegintOptIn bool 
}

// This order struct seems reasonable and in line with what we see in many order 
// flows online today. 
// However, including marketing opt-in in this aggregate is a bad design for a couple 
// of reasons: 
// - firstly: from a bounded context perspective, marketing opt-in has nothing to do with 
// the order object 
// secondly, if a user were to opt out of marketing between starting an order and completing it 
// we would not want the order to not complete. Therefore, removing it from our aggrate mkaes sense 

type Order struct{
    items   []item 
    taxAmount money.Money 
    discount    money.Money 
    paymentCardID   uuid.UUID 
    customerID  uuid.UUID 
}

 ``` 

 #### Aggregates beyond a single bounded context 

 Especially at the business scale, there will be situations, where our bounded context changes 
 and other sub-systems would like to be notified. Beyoned our bounded context, we should expect (and aim for) eventual consitency. This means we expect the other systems to receiv and process our event in a reasonable amount of time, but we do not expect it to be atomically up-to-date as we would expect our bounded contexts to be.  This leads to more decoupled systems with stronger resilience and scalability possibilities. 