# Builder pattern
# Builder is a creational design pattern, which allows
# constructing complex objects step by step.
# Unlike other creational pattern, builder doesn't require products
# to have a common interface that makes it possible to produce differnt
# products.

# Usage examples: The builder pattern is a well known pattern in python
# It's especially ueful when you need to create an object with lots of
# possible confiuration options.

# Identification: The builder pattern can be recognized in a class, which has
# a single creation method and several methods to configure the resulting object.
# Builde methods often support chaining (for example, someBuilder.setValue(1).setValue(2))

from __future__ import annotations
from abc import ABC, abstractmethod

class Builder(abc):
    """The builder interface specifies methods for creating the different
    parts of the product objects"""


    @property
    @abstractmethod
    def product(self) -> None:
        pass

    @abstractmethod
    def produce_part_a(self) -> None:
        pass

    @abstractmethod
    def produce_part_b(self) -> None:
        pass

    @abstractmethod
    def produce_part_c(self) -> None:
        pass

class ConcreteBuilder1(Builder):
    """
    The concrete builder classes follow the Builder interface interface and
    provide specific implementations of the building steps. Your program may
    have several variations of Builders, implemented differently.
    """

    def __int__(self) -> None:
        self.reset()

    def reset(self) -> None:
        self._product = Product1()

    @property
    def product(self) -> Product1:
        product = self._product
        self.reset()
        return product

    def produce_part_a(self) -> None:
        self._product.add("PartA1")

    def produce_part_b(self) -> None:
        return self._product.add("PartB1")

    def produce_part_c(self) -> None:
        self._product.add("PartC1")


class Product1():

    """
    It makes sense to use the builder pattern only when your
    products are quite complex. and require extensive configuration

    Unlike in other creational patterns, different concrete builders can
    produce unrelated products. In other words, results of various builders
    may not always follow the same interface."""

    def __init__(self) -> None:
        self.parts = []

    def add(self, part: Any) -> None:
        self.parts.append(part)

    def list_parts(self) -> None:
        print(f"Product parts: {', '.join(self.parts)}", end="")


class Director:
    def __init__(self) -> None:
        self._builder = None

    @property
    def builder(self) -> Builder:
        return self._builder

    @builder.setter
    def builder(self, builder: Builder) -> None:
        self._builder = builder


    def build_minimal_viable_product(self) -> None:
        self.builder.produce_part_a()

    def build_full_featured_product(self) -> None:
        self.builder.produce_part_a()
        self.builder.produce_part_b()
        self.builder.produce_part_c()


if __name__ == '__main__':
    director = Director()
    builder = ConcreteBuilder1()
    director.builder = builder

    print("Standard basic product:")
    director.build_minimal_viable_product()
    builder.product.list_parts()

    print("\n")

    print("Standard full featured product: ")
    director.build_full_featured_product()
    builder.product.list_parts()

    print("\n")

    # Remember, the builder pattern can be used without a
    # Director class
    print("Custom products: ")
    builder.produce_part_a()
    builder.produce_part_b()
    builder.product.list_parts() 
