# GraphQL

=> GraphQL is a query language for your API, and a server-side runtime
for executing queires using a type system you define for your data. GraphQL
isn't tied to any specific database or storage engine and is instead backed
by your existing code and data.

=> A GraphQL service is created by defining types and fields on those types,
then providing functions for each field on each type.

```graphql

type Query{
  me: User
}

type User{
  id: ID
  name: String
}

// Along with functions for each field on each type:

function Query_me(user){
  return request.auth.user
}

function User_name(user){
  return user.getName()
}

```

After a GraphQL service is running (typically at a URL on a web service), it
can receive GraphQL queries to validate and execute. The service first checks
a query to ensure it only refers to the types and fields defined, and then
runs the provided functions to produce a result.

For example, the query:

{
  me{
    name
  }
}

could produce the following json result :

```json

{
  "me":{
    "name": "Luke Skywalker"
  }
}
```

## Queries and Mutations

==> Fields

At its simples, GraphQL is about asking for specific fiels on objects.

{
  hero{
    name          
  }
}

```json

{
  "data":{
    "hero":{
      "name": "R2-D2"
    }
  }
}
```

You can see immediatley that the query has exactly the shame shape as the
result. This is essential to GraphQL. because you always get back what you
expect, and the server knows exactly what fields the client is asking for.

## Arguments

You can pass arguments to fields

{
  human(id: "1000"){
    name
    height
  }
}

```json

{
  "data": {
    "human":{
      "name": "Luke Skywalker",
      "height": 172
    }
  }
}

```


## Aliases

Aliases let you rename the result of a field to anything you want

{
  empireHero: hero(episode: EMPIRE){
    name
  }
  jediHero: hero(episode: JEDI){
    name
  }
}

```json

{
  "data":{
    "empireHero":{
      "name": "Luke Skywalker"
    },
    "jediHero":{
      "name": "R2-D2"
    }
  }
}
```

## Fragments

GraphQL includes reusable units called fragments. Fragments let
you construct sets of fields, and then include them in queries where
you need to.

```gql

{
  leftComparison: hero(episode: EMPIRE){
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI){
    ...comparisonFields
  }
}

fragmentComparisonFields on Character{
  name
  appearsIn
  friends{
    name
  }
}

```

## Operation name

The operation type is either query, mutation, or subscription and describes what
type of operation you're intending to do. The operation type is required unless
you're using the query shorthand syntax, in which case you can't supply a name
or a variable definitions for your operation.
