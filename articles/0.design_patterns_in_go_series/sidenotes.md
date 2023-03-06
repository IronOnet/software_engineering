## Go Design Patterns Book: 

### The Factory Pattern 

- Delegationg the creation of new instances of structures to a different part of 
the program. 
- Working at the interface level instead of with concrete implementations. 
- Grouping families of objects to obtain a family object creator 

### Example: A factory of payment methods for a shop 

Providing different ways of paying at a ship. In the beginning, we will have 
two methods of paying casch and credit card.

A factory method has a very simple structure; we just need to identify how many 
implementations of our interface we are storing; and then provide a method 

```go 

type PaymentMethod interface{
    Pay (amount float32) string 
}

// The preceding lines define the interface of the payment method. They 
// define a way of making a payment at the shop. The factory method will 
// return instances of types that implement this interface. 

const (
    Cash = 1 
    DebitCard = 2
)

// We have to define the identified payment methods of the factory as 
// constants so that we can call and check the possible payment methods 
// from outside the package 
func GetPaymentMethod(m int) (PaymentMethod, error){
    switch m{
        case Cash:
        return new(CashPM), nil 
        case DebitCard:
        return new(DebitCardPM), nil 
        default: 
        return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
    }
}

type CashPM struct{} 
type DebitCardPM struct{} 

func (c *CashPM) Pay(amount float32) string{
    return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (d *DebitCardPM) Pay(amount float32) string{
    return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
```

### Go Abstract Factory 


```go 

package abstract factory 

type Vehicle interface{
    NumWheels() int 
    NumSeats() int 
}

type Car interface{
    GetMotorBikeType() int 
}

