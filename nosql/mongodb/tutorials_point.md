# MongoDB

MongoDB is a cross-platform, document oriented database that provides
high performance, high availability and easy scalability. MongoDB works
on the concept of collection and document

## Database

A database is a physical container of collections. Each database gets its own
set of files on the file system. A single MongoDB server typically has
multiple databases.

## Collection

A collection is a group of MongoDB documents. It's the equivalent of a table
in an RDBMS  

## Document

A document is a set of key-value pairs. Documents have dynamic schema.
Dynamic schema means that document in the same collection do not need to
have the same set of fields or structure. and common fields in a collection's
document may hold different types of data.

the "_id" field in a document is a 12 bytes hexadecimal number which assures
the uniqueness of every document. You can provide _id while inserting the
document. If you don't provide them MongoDB provides a unique id for every
document.

## Why Use MongoDB

- Document oriented storage: Data is stored in the form of JSON style documents
- Index on any attribute
- Replication and high availability
- Auto-Sharding
- Rich queries
- Fast in-place updates
- Professional support by MongoDB



## Where to use MongoDB?

  - Big Data
  - Content management and delivery
  - Mobile and social infrastructure
  - User data management
  - Data Hub


# MongoDB Data Modeling

Data in MongoDB has a flexible schema.documents in the same colection. They do
not need to have the same set of fields or structure. Common fields in a colleciton
documents may hold different types of data.

## Data Modeling Design  

MongoDB provides two types of data models. Embedded data model and
Normalized data model.

### Embedded data modeling

In this model, you can have (embedded) all the related data in a single document,
it is also known as de-normalized data model.

For example, assume we are getting the details of employees in three different
documents namely, Personal_details, Contact, and Address, you can embedd all the
three documents in a single one as shown below.

```js

{
  _id:,
  Emp_ID: "1002AE356",
  Personal_details:{
    First_Name: "Radhika",
    Last_Name: "Sharma",
    Date_Of_Birth: "1995-12-12"
  },
  Contact:{
    e-mail: "radhika_sharma.123@gmail.com",
    phone: "9848222338",
  },
  Address:{
    city: "Hyderabhad",
    Area: "Madapur",
    State: "Telangana"
  }
}

```

### Normaized Data Model

In this model, you can refer to the sub documents in the original document,
using references. For example, you can re-write the above document in the
normalized model as:

```js

// Employee
{
  _id: <ObjectId101>,
  Emp_ID: "10025AE336"
}

// Personal_details
{
  _id: <ObjectId102>,
  empDocId: "ObjectId101",
  FirstName: "Radhika",
  Last_Name: "Sharma",
  Date_Of_Birth: "1995-12-12"
}

// Contact

{
  _id: <ObjectID103>,
  empDocId: "ObjectId101",
  e-mail: "radhika_sharma.123@gmail.COM",
  phone: "984802238",
}

// Address
{
  _id: <ObjectId104>,
  empDocId: "ObjectId101",
  city: "Hyderabhad",
  Area: "Madapur",
  State: "Telangana",
}

```

### Considerations while designing a schema in MongoDB

- Design your schema according to user requirements
- Combine objects into one document if you will use them together.
Otherwise separate them (but make sure there should not be need of joins).
- Duplicate the data (but limited) because disk space is cheap as compare to
compute time

- Do joins at write time but not at read time
- Optimize your schema for most frequent use cases
- Do complex aggregation in the schema.


## Mongo Create Database

```js

>>use mydb // creates a new database calle mydb

```

## The createCollection method

MongoDB db.createCollection(name, options) is used to create collection

In the command, name is the name of the colleciton to be created, Options is
a document and is used to specify configuration of collection.

=> Options parameter is optional, so you need to specivy the name of the collection
Following is the list of options you can use

- capped (Boolean) if true, enables a capped collection. Capped collection
is a fixed size collection that automatically overwrites its oldes entries when it
reaches its maximum size. if you specify true, you need to specify size parameter
also

- autoIndexed (Boolean): (Optional) if treu, automatically create index in _id fields
Default value is false

- size (number) (Optional) Specifies a maximu size in bytes for a capped collection

- max (number) (Optional) Specifies the maximum number of documents allowed in the
capped collection.

# Data Types

Mongodb supports many datatypes. Some of them are:

  - String: This is the most commonly used datatype to store the data. string
  in mongoDB must be UTF-8 valid

  - Integer: this type is used to store a numerical value. Integer can be 32 bit
  or 64 bit depending upon your server

  - Boolean: This type is used to store a boolean (true/false) value

  - Double: This type is used to store floating point values.

  - Min/max keys: This type is used to compare a value against the lowest and higest
  BSON elements

  - Arrays: This type is used to store arrays or list or multiple values into one
  key

  - Timestamp: ctimestam. this can be handy for recording when a document has been
  modified or added.

  - Object: this datatype is used for embedded docuemnts

  - Null: This type is used to store a Null value

  - Symbol: this datatype is used identically to a string; it is generally reserved
  for languages that use a specific symbol type

  - Data - This data type is used to store the current date or time in UNIX time
  format, you can specify your own date tie by creating object of Date and passing
  day month, year into it.

  - ObjectID - This datatype is used to store the document's ID

  - Binary data: this dattype is used to stre Javascript Code into the document

  - Regular expression - This datatype is used to store regular expressions.                                                                 


## Dropping a collection

- the dropCollection method drops (deletes) a collection


## Insert Document


```js

db.user.insert({
  _id: ObjectId("......."),
  title: "MongoDB Overview",
  description: "MongoDB is a nosql database",
  by: "Tutorials point",
  url: "http://www.tutorialspoint.com",
  tags: ["mongodb", "database", "NoSQL"],
  likes: 100
})
```

You can also pass an array of documents into the insert() method as
shown below

```js

db.createCollection("post");
db.post.insert([
  {
    title: "MongoDB Overview",
    description: "MongoDB is a NoSQL database",
    by: "Tutorials Point",
    url: "http://www.tutorialspoint.com",
    tags: ["mongodb", "database", "NoSQL"],
    likes: 100
  },
  {
    title: "NoSQL Databases",
    description: "NoSQL databases dont have tables",
    by: "tutorials point",
    url: "http://www.tutorialspoint.com",
    tags: ["mongodb", "database", "NoSQL"],
    likes: 20,
    comments:[
      {
        user: "user1",
        message: "My First Comment",
        dateCreated: new Date(2013, 11, 10, 2, 35),
        like: 0
      }
    ]
  }
])

```

If you need to insert only one document into a collection you can use
the insertOne method

```js

db.createCollection("empDetails");

db.empDetails.insertOne(
  {
    First_Name: "Radhika",
    Last_Name: "Sharma",
    Date_Of_Birth: "1995-09-26",
    e_mail: "radhika_sharma123@gmail.com",
    phone: "98803027"
  }
)

```

You can insert multiple document using the insertMany() method. To this
method you need to pass an array of docments

```js

db.empDetails.insertMany([
  {
    First_Name: "Radhika",
    Last_Name: "Sharma",
    Date_Of_Birth: "1995-09-26",
    e_mail: "radhika_sharma123@gmail.com",
    phone: "9000012345",
  },
  {
    First_Name: "Vincent",
    Last_Name: "Vaghner",
    Date_Of_Birth: "1996-09-26",
    email: "vincevagner@gmail.com",
    phone: "80088998"
  }
])

```

## Query Document

### The find method

To query data from a MongoDB collection, you need to use MongoDB's find()
method

```js

db.collectionName.find()

```

The Pretty method display the results in a formatted way, you can use
pretty() method

```js

db.collectionName.find().pretty()

```

Apart from the find() method, there is findOne() method, that returns only one
document

```js

db.COLLECTION_NAME.findOne()

// Example
// the following example retrieves the docuemnt with the title
// MongoDB Overview

db.mycol.findOne({title: "MongoDB Overview"})


```

## MongoDB Save() method

The save() method replaces the existing document with the new document passed
in the save() method.

Syntax

>db.collection_name.save({_id: ObjectId(), NEW_DATA})

Example

```js

db.mycol.save(
  {
    "_id": ObjectId("50....860ea"),
    "title": "Tutorials Point New Topic",
    "by": "Tutorials Point"
  }
)

```

The findOneAndUpdate() method updates the values in  the existing document.
The basic syntax of findOneAndUpdate method is a sfollows

>db.COLLECTION_NAME.findOneAndUpdate(SELECTION_CRITERIA, UPDATED_DATA)

Example

Assume we have created a collection named empDetails and inserted three
docuemnts in it as shown below

```js

db.empDetails.insertMay([
  {
    First_Name: "Radhika",
    Last_Name: "Vishnu",
    Age: 26,
    e_mail: "radhikavish@gmail.com",
    phone: "9000054321"
  },
  {
    First_Name: "Subramanian",
    Last_Name: "Vishnu",
    Age: 27,
    e_mail: "submanvish@gmail.com",
    phone: "9000054321"
  }
])

```

Following example updates the age and email values of the document with
Name Radhika

```js

db.empDetils.findOneAndUpdate({
  {First_Name: 'Radhika'},
  {$set: {Age: '30', e_mail: 'radhika_newmeail@gmail.com'}}
})

```

## MongoDB updateOne() method

This method updates a single document which matches the given filter

The basic syntax of updateOne() method is as follosw

>db.COLLECTION_NAME.updateOne(<filter>, <update>)

```js

db.empDetails.updateOne(
  {First_Name: 'Rhadika'},
  {$set: {Age: '30', e_mail: 'radhika_newmeail@gmail.com'}}
)

```

## updateMany() method

The updateMany() method updates all the documents that matches the given
filter.

>db.COLLECTION_NAME.update(<filter>, <update>)

```js

db.empDetails.updateMany(
  {Age: {$gt: 25}},
  {$set: {Age: '00'}}
)


```

# MongoDB - Delete document

## The remove() Method

MongoDB remove() method is used to remove a document from the collection.
remove() accepts two parameters. One is the deletion criteria and the second is
justOne flag

>db.COLLECTION_NAME.remove(DELETION_CRITERIA)

```js

db.mycol.remove({'title': 'MongoDB Overview'})

```

### Remove only one

if there are multiple records and you want to delete only the first record
then set jusOne parameter in the remove method()

```js

db.collection_name.remvoe(DELETION_CRITERIA, 1)

```

## Remove All Documents

db.mycol.remove({})


## MongoDB - Limiting Records

### The Limit() method

To limit the records in MongoDB, you need to use limit() method. The method
accepts one number type argument, which is the number of documents that
you want to be displayed.

=> Syntax

> db.collection_name.find().limit(number)

```js

db.myCol.find({}, {"title": 1, _id: 0}).limit(2)
```

### Skip() method

A part from the limit() method, there is one more skip() which also accepts
number type arguments and is used to skip the number of docuemnts.

> db.collectiona_name.find().limit(number).skip(number)

```js

db.mycol.find({}, {"title": 1, _id:0}).limit(1).skp(1)
```

### The sort() method

To sort documents in MongoDB, you need to use sort() method. The method
accepts a document containing a list of fileds along with their sorting
order. To specify sorting order 1 and -1 are used. 1
is used for ascending order while -1 used for descending order

>db.collection_name.find().sort({Key: 1})  or db.collection_name.find().sort({key: -1})

```js

db.myCol.find({}, {"title": 1, _id: 0}).sort({"title": -1})

```


## MongoDB  - Indexing

Indexes support the eficient resolution of queries. Without indexes, MongoDB
must scan every document of a collection to select those documents that match
the query statement.

### The createIndex() Method

To create an index, you need to use createIndex() method of MongoDB

### Syntax

The basic syntax of createIndex() method is as follows(.

>db.collectio_name.createIndex({KEY: 1})

Here key is the name of the field on which you want to create index and 1 is
for ascending order. to create index in descending order you need to use -1.

>db.myCol.createIndex({"title": 1})


```js

db.myCol.createIndex({"title": 1})

```

In createIndex() method you can pass multiple fields, to create index on multiple
fields.

```js

db.mycol.createIndex({"title": 1, "description": -1})

```

This method also accepts list of options (which are optional). Following is the
list :

  - background (Boolean): Builds the index in the bacground so that building an
  index does not block other database activities. Specify true to build in
  the background. The default value is false.

  - unique (Boolean): Creates a unique index so that the collection will not accept
  insertion of documents where the index key or keys match an existing value in
  the index. Specify true to create a unique index. The default value is false.

  - name (string): The name of the index. if unspecified. MongoDB generates an
  index name by concatenating the names of the indexed fields and the sort order

  -


## MongoDB - Aggregation

Aggregation operations process data records and return computed results.
Aggregation operations group values from multiple documents together, and
can perform a variety of operations on the grouped data to return a single
results

### The aggregate() method

For the aggregation in MongoDB, you should use aggregate() method

>db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)

```js

db.mycol.aggregate([{$group: {_id: "$by_user", num_tutorial : {$num: 1 }}}])
```

the SQL equivalent of the above command is

```sql

SELECT COUNT(*) from mycol
GROUP BY by_user;

```

The following is a list of available aggregation expressions.


- $sum: Sums up the defined value from all documents

```js

db.mycol.aggregate([$group: {_id: "$by_user", num_tutorial: {$sum: "$likes"}}])
```

- $avg: Calculates the average of all given values from all documents in the
collection

```js

db.mycol.aggregate([{$group: {_id: "$by_user", num_tutorial: {$avg:"$likes"}}}])
```

- $min: Gets the minimul of the corresponding values from all documents
in the collection

- $max: gets the maximum of the corresponding values from all documents
in the collection

- $push: Insert the value to an array in the resulting document

-


### Pipeline operations

Mongodb supports the concatenation of various operations

$db.find().limit().skip().sort()


## MongoDB - Replication
