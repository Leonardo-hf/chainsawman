from .index import AllDepsHandler
from .go import GoDepsHandler
from .java import JavaDepsHandler
from .python import PyDepsHandler
from .rust import RustDepsHandler

__all__ = ['AllDepsHandler', 'GoDepsHandler', 'JavaDepsHandler', 'PyDepsHandler',
           'RustDepsHandler']
