## HTML5 - Attributes

accessKy : Specifies a keyboard shortcut to access an element
align: right, left, center Horizontaly align tags

background (URL): Places a background image behind an element
bgColor (Numeric hexadecimal, RGB values): Places a background color behind an element

class (User defined): Classifies an element for use with Cascading Style Sheets

contenteditable (true, false): Specifies if the user can edit the elements content or not

contextmenu (Menu Id): Specifies the context menu for an element

draggable (true, false, auto): Specifies whether or not a user is allowed to drag an element

height (Numeric Value): Specifies the tables, imagges or table cells

hidden:

## Events

When users visit your website they perform various activities such as clicking
on text and images and links, hover over defined elements etc. These are examples
of what Javascript calls events.

We can write our event handler in Javascript and you can specify these event
handlers as value or event tag attribute.


## Webform

The <input> element

text: A free-form text field, nominally free of line breaks.
password: A free-form text field for sensitive information nominally free of line
breaks

checkbox: A set of zero or more values from a predefined list

radio: An enumerated value

submit:

## Server-Sent Events

Conventional web applications generate events which are dispatched to the web
server. For example a simple click on a link requests a new page from the server

The types of events which are flowing from the web browser to the web server may
be called client-sent events.

HTML5 and WHATWG Web Applications 1.0 introduced events which flow from
the web server to the web browser and they are called SSE. Using SSE you can
push DOM events continously from your web server to the visitor's browser.

The event streaming approach opens a persistent connection to the server, sending
data to the client when new information is available, eliminating the need
for continous polling.

### Web app for SSE

To use Server-Sent Events in a web application, you would need to add an
<eventsource> element to the document.

The src attribute of <eventsource> element should point to an URL which should
provide a persistent HTTP connection that sends data stream containing the
events.

The URL would point to a PHP, PERL or any Python script which would take care
of sending event data consistenttly. Following is

```html

<html>
  <head>

    <script type="text/Javascript">
      /** Define event handling logic here */
    </script>
  </head>

  <body>
      <div id="sse">
          <eventsource src="/cgi-bin/ticker.cgi"/>
      </div>

      <div id="ticker">
        <TIME>
      </div>
  </body>
</html>

```

### Server Side Script

A server side script should send Content-Type header specifying the type
of text/event-strea as follows:

```perl

print "Content-Type: text/event-stream\n\n";

```

After setting Content-Type, the server side script would send an Event tag
followed by event name. The following examplewould send Server-Time as event
name terminated by a new line character

```perl
print "Event: server-time\n";

while(true){
  print "Event: server-time\n";
  $time = localtime();
  print "Data: $time\n";
  sleep(5);
}
```

```html

<html>
  <head>
    <script type="text/javascript">
        document.getElementByTagName("eventsource")[0].addEventListener(eventHandler, false);

        function eventHandler(event){
          // Aalert time sent by the server
          document.querySelector("#ticker").innerHtml = event.data;
        }
    </script>
  </head>

  <body>
    <div id="sse">
      <eventsource src="/cgi-bin/ticker.cgi"/>
    </div>

    <div id="ticker" name="ticker">
        [TIME]
    </div>
  </body>
</html>
```

### Websockets

Websocket is a next-generation bidirectional communication technology for web
applications which operates over a single socket and is exposed via Javascript
interface in HTML 5 compliant browsers.

Once you get a WebSocket connectin with the web server, you can send data from
browser to server by calling a send() method and receive data from server to browser
by an onmessage event handler.

Following is the API which creates a new WebSocket object

```javascript

let websocket = new WebSocket(url, [protocol]);

```

Here the first argument, url specifies the URL to which to connect. The second
attribute, protocol is optional, and if present, specifies a sub-protocol that
the server must support for the connection to be successful.

#### Websocket Attributes

The readonly attribute readyState represents the state of the connection. it can
have the following values

A value of 0 indicates that the conneciton has not yet been established.
A value of 1 indicates that the connection is established and communication is possible.
A value of 2 indicates that the connection is going through the closing handshake
A value of 3 indicates that the connection has been closes or could not be opened.

### Websocket Events

Following are the events associated with WebSocket object. Assuming we created
Socket object as mentioned above

open ===> Socket.onOpen : this event occurs when socket connection is established
message ====> Socket.onMessage: this event occurs when client receives data from server
error ===> Socket.onError: this event occurs when there is any error in communication
clise ===> Socket.onClose This event occurs when the connection is closed.

### Client-side websocket code

```html

<html>
  <head>
    <script type="text/javascript">
      function webSocketTest(){
        if ("Websocket" in window){
          alert("Websocket is supported by your browser!");

          // Let us open a web socket
          var ws = new WebSocket("ws://localhost:9998/echo");

          ws.onopen = function(){
            // Web Socket is connected, send data using send()
            ws.send("Message to send");
            alert("Message is sent...");
          };

          ws.onmessage = function(event){
            var received_msg = evt.data;
            alert("Message is received...");
          };

          ws.onclose = function(){
            // Websocket is closed.
            alert("Connection is closed...");
          };
        } else{
          // The Browser doesn't support Websocket
          alert("Websocket Not Supporrted by your Browser!");
        }
      }

    </script>
  </head>

  <body>
    <div id="sse">
      <a href="javascript:WebSocketTest()">Run Websocket</a>
    </div>
  </body>
</html>


```

HTML5 element <canvas> gives you an easy and powerful way to draw graphics using
javascript. It can be used to draw graphs, make photo compositions or do simple
(and no so simple) animations.

```html

<canvas id="mycanvas" width="100" height="100">My Canvas</canvas>

```

## Audio & Video

Here's the simples form of embedding a video file in your webpage

```html

<video src="foo.mp4" width="300" height="200" controls>
  Your browser does not support the <video> element
</video>

```
```html

<html>
  <body>
    <video width="300" height="200" controls autoplay>
      <source src="/html5/foo.ogg" type="video/ogg"/>
      <source src="/html5/foo.mp4" type="video/mp4"/>
      Your browser does not support the <video> element.
    </video>
  </body>
</html>

```


### Geolocation

Refer to the Geolocation API

### Microdata

Microdata is a standardized way to provide additonal semantics in your webpages
Microdata lets you defined your own customized eleements and start embedding
custom properties in your web pages. At a high level, microdata consists of a group
of name-value pairs

The groups are called items, and each name-value pair is a property. Item and
properties are represented by regular elements.

Example:
  - To create an item, the itemscope attribute is used.
  - To add a property to an item, the itemprop attribute is used on one of the item's
  descendants.

Here are two items, each of which has the property "name"

```html

<html>
  <body>
      <div itemscope>
        <p>My Name is <span itemprop="name">Nuha</span>.</p>
      </div>

      <div itemscope>
        <p>My Name is <span itemprop="name">Nuha</span></p>
      </div>

  </body>
</html>

```

## Drag and Drop  

To achieve drag and drop functionality with traditional HTML4, developers would
either have to use complex JS programming or other Javascript frameworks.

HTML5 came up with a native DND support to the browser making it much easier to
code up.

### Drag and Drop Events

There are number of events which are fired during various stages of the drag and
drop operatons. These events are listed below:

- dragStart: fires when the user starts dragging of the object
- dragenter: fired when the mouse is first moved over the target element while
a drag is occuring. A listener for this event should indiciate wheter a drop is allowed
over this location. if there are no listeners, or the listeners perform no operatons
that occurs during a listener it will be the same as the dragenter event.

- dragleave: this event is fired when the mouse leaves an element while a drag is occuring
Listeners should remove any highlighting or insertin markers used for drop feedback.

- drag: Fires every time the mouse is moved while the object is being dragged.

- drop: The drop event is fired on the element where the drop has occurred at the
end of the drag operation. A listener would be responsibble for retrieving the data
being dragged and listening it at the drop location.

- dragend: Fires when the listener releases the mouse button while dragging an object.

#### Drag and Drop Process

Following are the steps to be carried out to implement Drag and Drop operation

Step 1 - Making an Object Draggable

Here are steps to be taken:
  - If you want to drag an eleemnt, you need to set the draggable attribute to true
  for that element
  - Set an event listener for dragstart that stores the data being dragged

```html

<html>
  <head>
    <style type="text/css">
        #boxA, #boxB{
          float: left; padding: 10px; margin:10px; -moz-user-select:none;
        }

        #boxA { background-color: #6633FF; width:75px; height: 75px; }
        #boxB { background-color: #FF6699; widfth: 150px; height: 150px; }
    </style>

    <script type="text/javascript">
      function dragStart(ev){
        ev.dataTransfer.effectAllowed = "move";
        ev.dataTransfer.setData("Text", ev.target.getAttribute("id"))
        ev.dataTransfer.setDragImage(ev.target, 0, 0);

        return true;
      }

      function dragEnter(ev){
        event.preventDefault();
        return true;
      }

      function dragOver(ev){
        return false;
      }

      function dragDrop(ev){
        var src = ev.dataTransfer.getData("Text");
        ev.target.appendChild(document.getElementById(src))
        ev.stoPropagation();
        return false;
      }

    </script>
  </head>

  <body>
    <center>
      <h2>Drag and drop HTML5 demo</h2>
      <div>Try to drag the purple box around.</div>

      <div id="boxA" draggable="true"
        ondragstart = "return dragStart(ev)">
        <p>Drag Me</p>
      </div>

      <div id="boxB">Dustbin</div>
    </center>

  </body>
</html>

```

## Web Workers

Javascript was designed to run in a single-threaded environment, meaning
multiple scripts cannot run at the same time. Consider a situation where you
need to handle UI events, query and process large amount of API data, and
manipulate the DOM.

Javascript will hang your browser in situation where CPU utilization is high.

```html

<html>
  <head>
    <title>Big for loop</title>
    <script>
        function bigLoop(){
          for(let i = 0; i < 1000; i++){

            let j = i;
          }
          alert("Completed " + j + "iterations");
        }

        function sayHello(){
          alert("Hello sir...");
        }
      </script>
    </head>

    <body>
      <input type="button" onClick="bigLoop();" value="Big Loop"/>
      <input type="button" onClick="sayHello();" value="Say Hello"/>
    <body>
</html>

```

Web workers can handle situations where computationaly expensive tasks that
may hang the browser need to be performed without interrupting the user interface
and typically run on seprate threads.

Web workers allow for long-running scripts that are not interrupted by scripts
that respond to clicks or other user interactions, and allows long tasks to
be executed without yielding to keep the page responsive.

## How Web Workers Work?

Web workers are initialized with the URL of a Javascript file, which contains
the code the worker will execute. This code sets event listeners and communicates
with the script that spawned it from the main page. Following is the simple syntax

```javascript

let worker = new Worker('bigLoop.js');

```

If the specified javascript file exists, the browser will span a new worker thread
which is downloaded asynchronously. If the path to your worker returns a 404 error.
The worker will fail silently.

If your application has multiple surrounding Javascript files, you can import them
with importScripts() method which takes file name(s) as arguments separated by comma

```javascript

importScripts("helper.js", "anotherHelper.js");

```

Once the web worker is spawned, communication between web orker and its parent
page is done using the postMessage() method. Depending on your browser/version,
postMessage() can accept either a string or JSON object as its single argument.

Message passed by Web Worker is accessed using onmessage event in the main
page.

```html

<html>
  <head>
    <title>Big For Loop</title>
    <script>
      let worker = new Worker("bigLoop.js");

      worker.onmessage = function(event){
        alert("Completed " + event.data + "iterations");
      };

      function sayHello(){
        alert("Hello Sir...");
      }

    </script>
  </head>

  <body>
    <input type="button" onClick="sayHello();" vlaue="Say Hello"/>
  </body>
</html>

```

## IndexDB

IndexDB is a new HTML5 concept to store data inside user's browser. indexdb
has more power than local storage and useful for applications that requires
to store large amount of data. These application can run more efficiently and
load faster.

### Why use IndexDB?

- It stores key-value pairs
- It's not a relational db
- It's API is mostly asynchronous
- It's not a structured query language.
- It's support access to data from the same domain;


## Web messaging

Web messaging is the ability to send realtime messages from the server to the
client browser. It overrides the cross domain communicaiton problem in different
domains, protocols or ports.  

For example, you want to send the data from your page to ad container which is placed
at iframe or vice-versa, in this scenario, Browser throws a security exception.
With web messaging we can pass the data across as a message event.

### Message Event

Message events fires Cross-Document messaging, channel messagin, server-sent events
and web sockets. i

#### Examples

==> Sending messag from iframe to button

```javascript

let iframe = document.querySelector('iframe');
let button = document.querySelector('button');

let clickHandler = function(){
  iframe.contentWindow.postMessage('The message to send.', 'https://www.tutorialspoint.com');
}

button.addEventListener('click', clickHandler, false);

```

===> Receiving a cross-document message in the receiving document.

```js

let messageEventHandler = function(event){
  // Check that the origin is one we want
  if(event.origin === 'https://www.tutorialspoint.com'){
    alert(event.data);
  }
}

window.addEventListener('message', messageEventHandler, false);

```

### CORS

Cross-Origin Resource Sharing (CORS) is a mechanism that allows the restricted from
another domain to access the web browser

### Web RTC

Web RTC is a technology specs that support browser-to-browser applications for
voice calling, video chat, and P2P file sharing.

Web RTC implements three APIs as shown below:

- MediaStream - get access to the user's camera and microphone
- RTCPeerConnection: get acces to audio or video calling facility
- RTCDataChannel : get access to peer-to-peer communication

#### Media Stream

The media stream represent synchronoized streams of media.

```javascript

function getStream(stream){
  window.AudioContext = window.AudioContext || window.webKitAudioContext
  var audioContext = new AudioContext();

  // create an AudioNode from the stream
  var mediaStreamSource = audioContext.createMediaStreamSource(stream);

  // Connect it to destination to hear yourself
  // or any other node for processing
  mediaStreamSource.connect(audioContext.destination);
}

navigator.getUserMedia({ audio: true}, getStream);

```

#### Screen Capture


#### Session Control, Network & Media Information

Web RCT requires peer-to-peer communication between browsers. This mechanism
requires signaling, network information, session control and media information.
developers can pick from such as SIP or XMPP or any two way communications.

##### Samp^le code of createSignalChannel()

```javascript

var singnalingChannel = createSignalChannel();
var pc;
var configuration = ...;

// run start(true) to initiate a call
function start(isCaller){
  pc = new RTCPeerConnection(configuration);

  // send any ice candidates to the other peer
  pc.onicecandidate = function(evt){
    singnalingChannel.send(JSON.strigify({"candidate": evt.candidate }))
  };

  // Once remote streams arrives, show it in the remote video element
  pc.onAddStream = function(evt){
    remoteView.src = URL.createObjectURL(evt.stream);
  }

  navigator.getUserMedia({ "audio": true, "video": true }, function(stream){
    selfView.src = URL.createObjectURL(stream)
    pc.addStream(stream);

    if(isCaller){
      pc.createOffer(getDescription);
    } else{
      pc.createAnswer(pc.remoteDescription, gotDescription);

      function gotDescription(desc){
        pc.setLocalDescription(desc);
        singnalingChannel.send(JSON.stringify({ "sdp": desc }))
      }
    }
  })
}
