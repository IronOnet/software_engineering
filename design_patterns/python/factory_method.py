# The Factory method is a creational design pattern which
# solves the problem of creating product objects without specifying
# their concrete classes.

# The factory method defines a method, which should be used for creating
# objects instead of using a direct constructor call (new operator).
# Subclasses can override this method to change the class of objects that
# will be created

# THE Factory method patter is widely used in Python code. It's very
# useful when you need to provide a high level of flexibility for your
# code.

from __future__ import annotations
from abc import ABC, abstractmethod

class Creator(ABC):
    """
    The creator class declares the factory method that is supposed to
    return an object of a Product class. The creators subclasses usually
    provide the implementation of this method"""

    @abstractmethod
    def factory_method(self):
        pass

    def some_operation(self) -> str:
        product = self.factory_method()

        result = f"Creator: the same creator's code has just worked with {product.operator}"

        return result


class ConcreteCreator(Creator):
    """
    Note that the signature of the method still uses the abstract product
    type, even though the concrete product is actually returned from the
    method. This way the creator can stay independent of concrete product classes."""

    def factory_method(self) -> Product:
        return ConcreteProduct1()

class ConcreteCreator2(Creator):
    def factory_method(self) -> Product:
        return ConcreteProduct2()


class ConcreteProduct(Product):
    def operation(self) -> str:
        return "{Result of the Concrete Product1}"

class ConcreteProduct2(Product):
    def operation(self) -> str:
        return "{Result of the concreteProduct2}"

def client_code(creator: Creator) -> None:
    """
    The client code works with an instance of concrete creator, albeit
    through its base interface. As long as the client keeps working with
    the creator via the base interface, you can pass it any creator's subclass"""

    print(f"Client: I'm not aware of the creator's class, but it still works.\n"
    f"{creator.some_operation()}", end="")


if __name__ == '__main__':
    print("App: Launched with the ConcreteCreator1.")
    client_code(ConcreteCreator1())
    print("\n")

    print("App: Launched with the ConcreteCreator2.")
    client_code(ConcreteCreator2())
