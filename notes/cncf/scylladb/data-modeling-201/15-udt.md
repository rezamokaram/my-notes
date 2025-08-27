# UDT

User-defined types allow the user to define more complex data structures and attach multiple data fields each with a name and type, to a single column. This adds flexibility to the data model. This is covered in detail here. The fields used to define the UDT can be of any valid type, including collections and other UDTs.

Let’s look at a simple example (make sure you are in the same active cqlsh prompt you created before):

```cql
CREATE TYPE phone (
    country_code int,
    number text,
);
```

```cql
CREATE TYPE address (
   street text,
   city text,
   zip text,
   phones map<text, frozen<phone>>
);
```

```cql
CREATE TABLE pets_v4 (
  name text PRIMARY KEY,
  addresses map<text, frozen<address>>
);
```

We define a type “phone” which has a country code and a number. We then define a type “address,” which has the fields street, city, zip, and a collection (map in this case) of phones. Our table “pets_v4” has the username as the Primary key and a collection (again map) of addresses for that pet.  An insert would then look like:

```cql
INSERT INTO pets_v4 (name, addresses) 
               VALUES ('Rocky', {
                  'home' : {
                   street: '1600 Pennsylvania Ave NW',
                   city: 'Washington',
                   zip: '20500',
                   phones: { 'cell' : { country_code: 1, number: '202 456-1111' },
                             'landline' : { country_code: 1, number: '202 456-1234' } }
                           },
                   'work' : {
                             street: '1600 Pennsylvania Ave NW',
                             city: 'Washington',
                             zip: '20500',
                             phones: { 'fax' : { country_code: 1, number: '202 5444' } }
                            }
               });
```

`A UDT must be frozen – meaning you cannot update individual components of the UDT value. The whole value must be overwritten.`
