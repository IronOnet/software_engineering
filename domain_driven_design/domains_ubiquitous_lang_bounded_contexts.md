## Understanding Domains, Ubiquitous language and Bounded Contexts

The domain is the central entity in DDD; its what we will model our entire language 
and system around. 

### Ubiquitous language

Ubiquitous language is the overlap of the language that domain experts and technical experts use.

This language should be used, when discussing requirements and system design and should even 
be used in the source code itself. Plus, it should eveove, therefore, you should spend time 
evaluating and updating it regularly

### Bounded context 

Bounded contexts are all about dividing large models into smaller, easier-to-understand chunks 
and beign explicit about how they relate to each other. 

Since several bounded contexts often must communicate, we often apply patterns to ensure our 
models can maintain integrity. The three main patterns are as follows: 

    - Open Host Service 
    - Published language 
    - Anti-corruption layer 

- An Open Host Service is a means of giving other systems (or sub-systems) access to ours. 
Typically an Open Host Service is an RPC. Some choices for RPC's might be to build a 
RESTFUL API, implement gRPC, or perhaps even an XML API. 

- Published Language: A ubiquitous language is a team's internal formally defined language; 
a published language is the opposite. if a team is going to expose some of it system's to other 
teams via an Open Host Service, we need to ensure the definition of what the team exposes to 
other teams in different bounded contexts is clear.

Two popular ways to present published languages are via OpenAPI or gRPC 

** OpenAPI 

We can use OpenAPI to define the schema. This is a popular approach as you can also generate client 
and server code to speed up development for your team and for consumer teams too. You can use 
a tool called Swagger for this.

```swagger 

swagger: "2.0" 
info: 
    description: "Public documentation for payment & subscription system" 
    version: "1.0.0" 
    title: "Payment & Subscription API" 
    contact: 
        email: "ourteams@subs.com" 
    host: "api.payments.com" 
    schemes: 
        - "https" 
    paths: 
        /users: 
            get: 
                summary: "Return details about users" 
                operationId: "getUsers" 
                produces: 
                    - "application/json" 
                responses: 
                    "200": 
                        description : "successful operation" 
                        schema: 
                            $ref: "#/definitions/User" 
                    "400": 
                        description: "bad request" 
                    "404": 
                        description: "users not found" 
                definitions:
                    User: 
                        type: "object" 
                        properties: 
                            id: 
                                type: "integer" 
                                format: "int64" 
                            username: 
                                type: "string" 
                            subscriptionStatus: 
                                type: "boolean" 
                            subscriptionType: 
                                type: "string" 
                            email: 
                                type: "string" 
                    ApiResponse: 
                        type:"object" 
                        properties: 
                            code: 
                                type:"integer" 
                                format: "int32" 
                            type: 
                                type: "string" 
                            message: 
                                type: "string
```

- Anti Corruption layer: Sometimes called an adapter, an anti-corruption layer can be used to 
translate models from different systems. It is a complementary pattern that works well withthe 
Open Host Service. For example, the marketing teams. 

Suppose that the markegint team published language that defines a campaign looks like the following 
```json 
    {
        "id": "4ccbd4ba9-7c04-4a3d-ac52-71f37ba75d7f", 
        "metadata":{
            "name": "some campaign", 
            "category": "growth", 
            "endDate": "2023-04-12"
        }
    }

``` 

However, our internal model for a campaign might look like this 

```go 

type Campaign struct{
    id  string 
    title string 
    goal string 
    endDate time.Time
}

``` 

As you can see, we caere about most of the same information, but we name it differently or have a 
slightly different format. We have two options here: 

    - We can swap our campaingn model to be exactly the same as the marketing model. This woud 
    go against the principle of DDD and mean we are strongly coupling the domain model to something 
    outside our control 
    - We can write an anti-corruption aleyer 

An anti corruption layer would look like this. 

```go 

import (
    "errors" 
    "time"
)

type Campaign struct{
    ID          string 
    Title       string
    Goal        string 
    EndDate     time.Time 
}

type MarketingCampaign struct{
    ID          string `json:"id"` 
    MetaData    struct{
        Name        string `json:"name"` 
        Category    string `json:"category"` 
        EndDate     string `json:"endDate"` 

    } `json:"metadata"`
}

func (m *MarketingCampaignModel) ToCampaign() (*Campaign, error){
    if m.Id == ""{
        return nil, errors.New("campagin ID cannot be empty")
    }
    formatDate, err := time.Parse("2006-0-1-02", m.Metadata.EndDate)

    if err != nil{
        return nil, errors.New("endDate was not in a parsable format")
    }

    return &Campaign{
        ID: m.Id, 
        Title: m.Metadata.Name, 
        Goal: m.Metadata.Category, 
        EndDate: formatDate,
    }, nil 
}

```