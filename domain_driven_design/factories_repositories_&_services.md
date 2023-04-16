# Factories, Repositories & Services 

These are the last major building blocks of domain driven design. 

In this chapter, we will cover: 
    - The factory pattern 
    - The repository pattern 
    - Services


## The Factory pattern 

The factory pattern is typically used in object-oriented prog. 
It is defined as an object with the primary responsibility of creating 
other objects. 

```php 

class Factory
{
    public static function build($carType){
        if($carType == "tesla"){
            return new Tesla();
        }
        if ($carType == "bmv"){
            return new BMV();
        }
    }
}

$myCar = Factory::build("tesla");

```

```go 

package car 

import (
    "errors" 
    "log" 
)

type Car interface{
    BeepBeep() 
}

type BMV struct{
    heatedSubscriptionEnabled bool 
}

func (b BMV) BeepBeep(){
    panic("implement me")
}

type Tesla struct{
    autoPilotEnabled bool 
}

func (t Tesla) BeepBeep(){
    panic("imeplement me")
}

func BuildCar(carType string) (Car, error){
    swith carType{
        case "bmv": 
            return BMV{heatedSubscriptionEnabled: true}, nil 
        case "tesla": 
            return Tesla{autoPilotEnabled: true}, nil 
        default: 
            return nil, errors.Error("uknown car type")
    }
}

```

Factories are a great way to standardize the creation of complex structs and can be usefull 
as your application grows in complexity. Factories also provide encapsulation. 
Finally factories can help ensure business invariants are enforce at the time of 
object creatinn, which can dramatically simplify our domain model.

For example, if we were creating a booking system for a hair salon, and someone tried to 
create a booking outside of business hours. we might enforce this in our CreateBooking factory function

```go 

package booking 

import (
    "errors" 
    "time" 

    "github.com/google/uuid" 
)

type Booking struct{
    id          uuid.UUID 
    from        time.Time 
    to          time.Time 
    hairDresserID    uuid.UUID 
}

func CreateBooking(from, to time.Time, hairDresserID uuid.UUID) (*Booking, error){
    closingTime, _ := time.Parse(time.Kitchen, "17:00pm") 
    if from.After(closingTime){
        return nil, errors.New("no appointments after closing time")
    }
    return &Booking{
        hairDresserID: uuid.New(), 
        id: uuid.New(), 
        from: from, 
        to: to, 

    }, nil 
}

```

### Entity Factories

Entities have identities and they have a minimum set of requirements necessary to instantiate 
them. We should, therefore, ensure we create entities that satisfy this minimum set 
of requirements when we create them via a factory. If we want to set other properties 
wen can then provide other functions. 

When designing an entity factory function, we need to decide whether we want the 
factory function to be responsible for generating the identity for our struct 
or whether we want to pass one as a parameter. 

## Implementing the repository pattern in Golang 

Repositories are parts of our code that contain the logic necessary to access data sources. 
A data source can be a wide variety of things, such as a file on disk, a spreadsheet, 
or an AWS S3 bucket, but in most projects it is a database.

By using a repository layer, you can centralize common data access code and make your system 
more maintainable by decoupling from a specifc database technology. For example your 
company may have a desire to move from one cloud provider to another and the database options 
are slightly different (Mysql and Mongo). In this instance we know we only need to rearchitect 
a small portion of our system.

Some developers query the database using other channels (such as CQRS) 

One mistake we often make with the repository layers is to make one struct per database table. 
This should be avoided, instead aim to make one struct per aggregate. 

Let's continue with the booking system example. We want to save our hair booking appointment 
to a database. We might define our interface as follows: 

```go 

type BookingRepository interface{
    SaveBooking(ctx context.Context, booking Booking) error 
    DeleteBooking(ctx context.Context, booking Booking) error 
}

```

An implementation of a simple repository layer for a Postgres database might look like this: 

```go 

type PostgresRepository struct{
    connPool *pgx.Conn 
}

func NewPostgresRepository(ctx context.Context, dbConnString string) (*PostgresRepository, error){
    conn, err := pgx.COnnect(ctx, dbConnString) 
    if err != nil{
        return nil, fmt.Errorf("failed to connect to db: %w", err)
    }

    defer conn.Close(ctx)

    return &PostgresRepository{connPool: conn}, nil 
}

func (p PostgresRepository) SaveBooking(ctx context.Context, booking Booking) error{
    _, err := p.connPool.Exec(
        ctx, 
        "INSERT into bookings (id, from, to , hair_dresser_id) VALUES ($1, $2, $3, $4)", booking.id.String(),
        booking.from.String(), 
        booking.to.String(), 
        booking.hairDresserID.String(),
    )

    if err != nil{
        return fmt.Errorf("failed to SaveBooking: %w", err)
    }

    return nil 
}

func (p PostgresRepository) DeleteBooking(ctx context.Context, booking Booking) error{
    _, err := p.connPool.Exec(
        ctx, 
        "DELETE from bookings WHERE id = $1", booking.id,
    )

    if err != nil{
        return fmt.Errorf("failed to DeleteBooking: %w", err)
    }
    return nil 
}

```

## Understanding Services 

In DDD, we use a few different types of services to help us organize our code. These are 
application services, domain services, and infrastructure services. 

### Domain Services 

Domain services are stateless operation within a domain that complete a certain activity. 
Sometimes, we will come across processes we cannot find a good way to model in an entity or value object; in these cases; it's a good idea to use a domain service. 

It is particulary tricky to outline rules to use domain services; however, some things that 
should look out for are the following: 

    - The code you are about to write performas a significant piece of business logic within the domain  
    - You are transfroming one domain object into another 
    - You are taking the properties of two or more domain objects to calcualte value. 

Services should always be expressed using ubiquitous language from within the bounded context, 
just like everything else we do in DDD. 

Look at the following code  

```go 

package shopping 

type Product struct{
    ID      int 
    InStock bool 
    InSomeonesCart bool
}

func (p *Product) CanBeBought() bool{
    return p.InStock && !p.InSomeonesCart 
}

type ShoppingCart struct{
    ID      int 
    Products []Product 
    IsFull bool 
    MaxCartSize int 
}

func (s *ShoppingCart) AddToCart(p Product)bool{
    if s.IsFull{
        return false
    }

    if p.CanBeBought(){
        s.Products = append(s.Products, p)
        return true 
    }
    if s.MaxCartSize == len(s.Products){
        s.IsFull = true 
    }
    return true 
}

``` 
The code looks reasonable at first, but it is problematic. While impleemnting ShoppingCart, 
we referenced another entity and added business logic, which does not really belong to 
ShoppingCart. To avoid this issue, we move the logic to a domain service, as follows. 

```go 

package checkout 

type CheckoutService struct{
    shoppingCart *ShoppingCart 
}

func NewCheckoutService(shoppingCart *ShoppingCart) *CheckoutService{
    return &CheckoutService{ shoppingCart: shoppingCart }
}

func (c CheckoutService) AddProductToBasket(p *Product) error{
    if c.shoppingCart.IsFull{
        return errors.New("cannot add to cart, its full")
    }
    if p.CannotBeBought(){
        c.ShoppingCart.Products = append(c.shoppingCart.Products, *p)
        return nil 
    }

    if c.shoppingCart.MaxCartSize == len(c.shoppingCart.Products){
        c.shoppingCart.IsFull = true
    }
    return nil 
}


``` 

We now have a central place to house domain logic that spans two entities. This will become 
even more useful as we add more logic to CheckoutService that may use more entities (perhaps) 
a discount entity or a shipping entity). 

Domain services are perfect for when we need to compose domain logic in a stateless manner. 
However, if this doesn't fit our use case, we likely need an application service. 

### Application services 

Application service are used to compose other services and repositories. They are responsible 
for managing transactional guarranties in place among various models. They should not contain 
domain logic (this belongs in the domain service, as discussed in the previous section)