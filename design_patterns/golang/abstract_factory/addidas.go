package main

type Addidas struct{

}

func (a *Addidas) makeShoe() Ishoe{
  return &AddidasShoe{
    Shoe: Shoe{
      logo: "addidas",
      size: 14
    },
  }
}

func (a *Addidas) makeShirt() IShirt{
  return &AddidasShirt{
    Shirt: Shirt{
      logo: "addidas",
      size: 14,
    }
  }
}
