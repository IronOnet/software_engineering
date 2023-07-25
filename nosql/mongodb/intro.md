## MongoDB

A document oriented database replaces the concept of a "row" with a more
flexible model, the "document", . By allowing embedded documents and arrays
the document oriented approach makes it possible to represent complex hierarchical
relationships with a single record.

MongoDB was designed to scale out. The document-oriented data model makes it
easier to split data accross multiple servers. MongoDB automatically takes care
of balancing data and load accross a cluster, redistributing documents automatically
and routing reads and writes to the correct machines.

### Features:

#### Indexing

MongoDB supports generic secondary indexes and provides unique, compoound,
geospatial, and full-text indexing capabilities as well. Secondary indexes on
hierarchical structures such as nested documents and arrays are also supported
and enable developers to take full advantage of the ability to model in ways
that best suit their applications.


#### Aggregation

MongoDB provides an aggregation framework based on the concept of data
processing pipelines. Aggregation pipelines allow you to build complex analytics
engines by processing data through a series of relatively simple stages on
the server side, taking full advantage of database optimizations.

#### Special collection and index types

MongoDB supports time-to-live (TTL) collections for data that should expire
at a certain time, such as sessions and fixed-size (capped) collections. for holding
recent data, such as logs.


### File storage

MongoDB supports an easy-to-use protocol for storing large files and file metadata.


Some features common to relational databases are not present in mongoDB, notably
complex joins. MongoDB supports joins in a very limited way through the use  
of the $lookup aggregation operator introduced in the 3.6 release. In the
3.6 release, more complex joins are possible using multiple join conditions as
well as unrelated subqueries.
