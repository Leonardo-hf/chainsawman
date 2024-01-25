from .index import ArchiveDepsHandler
from .go import GoDepsHandler
from .java import JavaDepsHandler
from .python import PyDepsHandler
from .rust import RustDepsHandler

__all__ = ['ArchiveDepsHandler', 'GoDepsHandler', 'JavaDepsHandler', 'PyDepsHandler',
           'RustDepsHandler']
