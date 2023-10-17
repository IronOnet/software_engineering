## Grokking DDD 

==> Entities are objects that are mutable, like database table rows often identified with 
a unique ID 

==> Value object on the other hand are immutable, the don't often require an ID field, they are 
objects without an identifieer and persistent values after creation. Value objects are often found inside domains and used to describe certain aspects in that domain. 

==> Aggregates are combined Entities and Value objects, An aggregate is a set of entities and 
value objects combined. 

==> The reason for an aggregate is taht the business logic will be applied on the Customer aggregate 
instead of each Entity holding the logic. An aggregate does not allow direct access to underlying 
entities. It is also common that multiple entities are needed to correctly represent data in real life.

==> An important rule in DDD aggregates is that they should only have one entity act as a root entity. 
What this means is that the reference of the root entitiey is also used to reference the aggregate. 
For our customer aggregate, this means that Person ID is the unique identifier. 



===> Factories  - Encapsulate complex logic 
The factory design pattern is a design pattern used to encapsulate complex logic in functions that 
creates the wanted instance, without the caller knowing anything about the implementation details 

```go
    func NewClient(cfg Config) (*Client, error){
        var addrs []string 

        if len(cfg.Addresses) == 0 && cfg.CloudID == ""{
            addrs = addrsFromEnvironment()
        } else{
            if len(cfg.Addresses) > 0 && cfg.CloudID != ""{
                return nil, errors.New("cannot create client: both addresses and cloudID are set") 
            }

            addrs = append(addrs, cloudAddr)
        }

        if len(cfg.Adresses) > 0{
            addrs = append(addrs, cfg.Addresses...)
        }

        ... 

        client := &Client{Transport: tp} 
        client.API = esapi.New(client) 

        if cfg.DiscoverNodeOnStart{
            go client.DiscoverNodes()
        }

        return client, nil 
    }

```


==> DDD suggest using factories for creating complex aggregates, repositories, and services. We wi

// customer.go 

```go

package aggregate 

import (
    "errors" 

    "github.com/google/uuid" 
    "github.com/irononet/ddd/entity" 
    "github.com/irononet/ddd-go/valueobject"
)

var (
    ErrInvalidPerson = errors.New("a customer has to have a valid person")
)


// Customer is an aggregate that combines all entities needed to represent a customer
type Customer struct{

    // Person is the root entity of a customer 
    // which means the person.ID is the main idenfier for this aggregate

    person *entity.Person 

    products []*entity.Item 
}


// NewCustomer is a factory to create a new Customer 
// aggregate, It will validate that the name is not 
// empty
func NewCustomer(name string) (Customer, error){


    // Validate that the Name is not empty
    if name == ""{
        return Customer{}, ErrInvalidPerson
    }

    person := &entity.Person{
        Name: name, 
        ID: uuid.New(), 
    }

    return Customer{
        person: person, 
        products: make([]*entity.Item, 0), 
        transactions: make([]*valueobject.Transaction, 0),
    }, nil 
}
```

The customer factory now helps with validating input, creating a new ID, and making sure all 
values are properly initialized. 