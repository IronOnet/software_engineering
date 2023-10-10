package main

type AK47 struct{
  Gun
}

func newAK47() IGun{
  return &AK47{
    Gun: Gun{
      name: "AK47 Gun",
      power: 4,
    }
  }
}
