from pathlib import Path
from typing import Union

from astroid import MANAGER, AstroidBuildingException, AstroidSyntaxError
from astroid.builder import AstroidBuilder
from astroid.nodes import Assign, AssignName, Call, Const, Keyword, List, Name, Tuple

from .exceptions import CouldNotParseRequirements
from .requirement import DetectedRequirement


class SetupWalker:
    def __init__(self, ast):
        self._ast = ast
        self._setup_call = None
        self._top_level_assigns = {}
        self.walk()

    def walk(self, node=None):
        top = node is None
        node = node or self._ast

        # test to see if this is a call to setup()
        if isinstance(node, Call):
            for child_node in node.get_children():
                if isinstance(child_node, Name) and child_node.name == "setup":
                    # TODO: what if this isn't actually the distutils setup?
                    self._setup_call = node

        for child_node in node.get_children():
            if top and isinstance(child_node, Assign):
                for target in child_node.targets:
                    if isinstance(target, AssignName):
                        self._top_level_assigns[target.name] = child_node.value
            self.walk(child_node)

    def _get_list_value(self, list_node):
        values = []
        for child_node in list_node.get_children():
            values.append(self._get_value(child_node))
        return values

    def _get_value(self, node):
        if not isinstance(node, Const):
            # we can't handle anything fancy, only constant values
            raise CouldNotParseRequirements
        return node.value

    def get_package_info(self):
        # first, if we have a call to setup, then we can see what its "install_requires" argument is
        if not self._setup_call:
            raise CouldNotParseRequirements
        artifact = ''
        version = ''

        def get_str_attr(child_node) -> str:
            v = ''
            if isinstance(child_node.value, Const):
                v = self._get_value(child_node.value)
            if isinstance(child_node.value, Name):
                try:
                    v = self._top_level_assigns[child_node.value.name]
                    v = self._get_value(v)
                except:
                    pass
            return v

        for child_node in self._setup_call.get_children():
            if not isinstance(child_node, Keyword):
                # do we want to try to handle positional arguments?
                continue
            if child_node.arg == 'name':
                artifact = get_str_attr(child_node)
            if child_node.arg == 'version':
                version = get_str_attr(child_node)
        return artifact, version

    def get_requires(self):
        # first, if we have a call to setup, then we can see what its "install_requires" argument is
        if not self._setup_call:
            raise CouldNotParseRequirements

        found_requirements = []

        for child_node in self._setup_call.get_children():
            if not isinstance(child_node, Keyword):
                # do we want to try to handle positional arguments?
                continue

            if child_node.arg not in ("install_requires", "requires"):
                continue

            if isinstance(child_node.value, (List, Tuple)):
                # joy! this is a simple list or tuple of requirements
                # this is a Keyword -> List or Keyword -> Tuple
                found_requirements += self._get_list_value(child_node.value)
                continue

            if isinstance(child_node.value, Name):
                # otherwise, it's referencing a value defined elsewhere
                # this will be a Keyword -> Name
                try:
                    reqs = self._top_level_assigns[child_node.value.name]
                except KeyError:
                    raise CouldNotParseRequirements
                else:
                    if isinstance(reqs, (List, Tuple)):
                        found_requirements += self._get_list_value(reqs)
                        continue

            # otherwise it's something funky and we can't handle it
            raise CouldNotParseRequirements
        # if we've fallen off the bottom with nothing in our list of requirements,
        #  we simply didn't find anything useful
        return found_requirements


def from_setup_py(setup_file: Union[str, Path]):
    if isinstance(setup_file, str):
        setup_file = Path(setup_file)

    try:
        ast = AstroidBuilder(MANAGER).string_build(setup_file.open().read())
    except (SyntaxError, AstroidBuildingException, AstroidSyntaxError):
        # if the setup file is broken, we can't do much about that...
        raise CouldNotParseRequirements

    walker = SetupWalker(ast)

    requirements = []
    for req in walker.get_requires():
        requirements.append(DetectedRequirement.parse(req, setup_file))

    return [requirement for requirement in requirements if requirement is not None]
