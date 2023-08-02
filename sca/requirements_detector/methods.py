from typing import List

import toml
from astroid import MANAGER
from astroid.builder import AstroidBuilder
from semver import parse_constraint

from requirements_detector.handle_setup import SetupWalker

from requirements_detector.requirement import DetectedRequirement

_PIP_OPTIONS = (
    "-i",
    "--index-url",
    "--extra-index-url",
    "--no-index",
    "-f",
    "--find-links",
    "-r",
)


def from_requirements_txt(text) -> List[DetectedRequirement]:
    # see http://www.pip-installer.org/en/latest/logic.html
    requirements = []

    for req in text.split('\n'):
        req = req.strip()
        if req == "":
            # empty line
            continue
        if req.startswith("#"):
            # this is a comment
            continue
        if req.split()[0] in _PIP_OPTIONS:
            # this is a pip option
            continue
        if req.startswith('[') and not req.startswith('[:'):
            # this is not necessary requirements
            break
        detected = DetectedRequirement.parse(req)
        if detected is None:
            continue
        requirements.append(detected)

    return requirements


def from_setup_py(text):
    try:
        ast = AstroidBuilder(MANAGER).string_build(text)
        walker = SetupWalker(ast)
        requirements = []
        for req in walker.get_requires():
            requirements.append(DetectedRequirement.parse(req))
        return [requirement for requirement in requirements if requirement is not None]
    except:
        # if the setup file is broken, we can't do much about that...
        return []


def from_setup_cfg(text) -> List[DetectedRequirement]:
    requirements = []
    start = False
    for req in text.split('\n'):
        if req.strip().startswith('install_requires'):
            start = True
            continue
        if start:
            if not req.startswith('\t'):
                break
            detected = DetectedRequirement.parse(req.strip())
            if detected is None:
                continue
            requirements.append(detected)
    return requirements


def from_pyproject_toml(text) -> List[DetectedRequirement]:
    requirements = []

    parsed = toml.loads(text)
    poetry_section = parsed.get("tool", {}).get("poetry", {})
    dependencies = poetry_section.get("dependencies", {})
    # dependencies.update(poetry_section.get("dev-dependencies", {}))

    for name, spec in dependencies.items():
        if name.lower() == "python":
            continue
        if isinstance(spec, dict):
            spec = spec["version"]
            if "version" in spec:
                spec = spec["version"]
            else:
                req = DetectedRequirement.parse(f"{name}")
                if req is not None:
                    requirements.append(req)
                    continue
        parsed_spec = str(parse_constraint(spec))
        if "," not in parsed_spec and "<" not in parsed_spec and ">" not in parsed_spec and "=" not in parsed_spec:
            parsed_spec = f"=={parsed_spec}"

        req = DetectedRequirement.parse(f"{name}{parsed_spec}")
        if req is not None:
            requirements.append(req)

    project_section = parsed.get('project', {}).get("dependencies", [])

    for req in project_section:
        req = DetectedRequirement.parse(req)
        if req is not None:
            requirements.append(req)
    return requirements
