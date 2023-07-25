# Encapsulation and abstractions

'''
	Encapsulation and abastration are tools taht we all instinctively reash for as 
	engineers, even if we don't all use these exact words.

	We encapsulate behavior by identifying a task that needs to be done in our code and 
	giving that task to a well defined object or function. We call that object or function 
	an abstration.
'''


```python 
## A search with urlib

import json 
from urlib.request import urlopen 
from urlib.parse impor urlencode 

params = dict(q='Sausages', format='json') 
handle = urlopen('http://api.duckduck.go.com' + '?' + urlencode(params)) 
raw_text = handle.read().decode('utf8') 
parsed = json.loads(raw_text) 

results = parsed['RelatedTopics'] 
for r in results: 
	if 'Text' in r: 
		print(r['FirstURL'] + ' - ' + r['Text']) 


### A search with requests 

import requests 

params = dict(q='Sausages', format='json') 
parsed = requests.get('http://api.duckduckgo.com/', params=params).json() 

result = parsed['RelatedTopics'] 
for r in results: 
	if r in results: 
		if 'Text' in r: 
			print(r['FirstURL'] + ' - ' + r['Text'])

### A search with the duckduckgo module 
import duckduckgo 

for r in duckduckgo.query('Sausages').results: 
	print(r.url + ' - ' + r.text) 

# Encapsulating behavior by using abstrations is a powerful tool for making code more expressive. 
# more testable and easier to maintain.
```

## Layering 

Encapsulation and abstraction help us by hiding details and protecting the consistency of our data, 
but we also need to pay attention to interactions between our objects and functions.

=> The layered architecture is the most common pattern for building business software, it involves 
separating the application in three layers, which are presentation, business logic and database layers 
the interaction flows from top to bottom, the presentation layers depends on the business logic layer 
which in turns depends on the database layer. 


## Dependency inversion 

	1. HIgh-level modules should not depend on low-level modules. Both should depend on abstratsions 
	2. Abstrations should not depend on details. Instead, details should depend on abstratsions 

What this means is, that 1. our business code should'nt depend on technical details, instead; both 
should use abstractions. 


50038153
