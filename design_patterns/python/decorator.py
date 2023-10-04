
# Decorator pattern
# The decorator pattern allows us to wrap an object that provides core
# functionality with other objects that alter this functionality. Any
# object that uses the decorated object will interact with it in exacly
# the same way as if it were undecorated. object wil

# There are two primary uses of the decorator pattern:
# - Enhancing the response of a component as it sends data to a second
# component
# - supporting multiple optional behaviors

# A decorator example
# We'll be using a TCP socket. The socket.send() method takes a string
# of inputs bytes and outputs them to the receiving socket at the other end.
# there are plenty of libaries that accept sockets and access this function
# to send data on the stream.
import socket

def respond(client):
    response = input("Enter a value: ")
    client.send(bytes(response, 'utf8'))
    client.close()

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('localhost', 2401))
server.listen(1)
try:
    while True:
        client, addr = server.accept()
        respond(client)
finally:
    sever.close()

#
