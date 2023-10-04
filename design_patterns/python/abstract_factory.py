from __future__ import annotations
from abc import ABC, abstractmethod

class AbstractFactory(ABC):
    """
    The abstact factory interface declares a set of methods that retun
    different abstract products. These products are called a family and are
    related by a high-level theme or concept. Products of one family are usually
    able to collaborate among themselves. A family of products may have several
    variants, but the products of one variant are incompatible with products of another
    """

    @abstractmethod
    def create_product_a(self) -> AbstractProductA:
        pass

    @abstractmethod
    def create_product_b(self) -> AbstractProductB:
        pass

class ConcreteFactory(AbstractFactory):
    """
    Concrete factory produce a family of products that belong to a single
    variant. The factory guarantees that resulting products
    """

    def create_product_a(self) -> AbstractProductA:
        return ConcreteProductA()

    def create_product_b(self) -> AbstractProductB:
        return ConcreteProductB()


class ConcreteProductA(AbstractProductA):
    def useful_function(self) -> str:
        return "Teh result of the product A"

class ConcreteProductB(AbstractProductB):
    def useful_function(self) -> str:
        return "The result of the product B"
