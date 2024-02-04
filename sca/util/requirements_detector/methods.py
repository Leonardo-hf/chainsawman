from typing import List, Tuple

import toml
from astroid import MANAGER
from astroid.builder import AstroidBuilder
from .poetry_semver import parse_constraint

from .handle_setup import SetupWalker

from .requirement import DetectedRequirement

_PIP_OPTIONS = (
    "-i",
    "--index-url",
    "--extra-index-url",
    "--no-index",
    "-f",
    "--find-links",
    "-r",
)


def from_requirements_txt(text) -> Tuple[str, str, List[DetectedRequirement]]:
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

    return '?', '?', requirements


def from_setup_py(text) -> Tuple[str, str, List[DetectedRequirement]]:
    try:
        ast = AstroidBuilder(MANAGER).string_build(text)
        walker = SetupWalker(ast)
        requirements = []

        for req in walker.get_requires():
            requirements.append(DetectedRequirement.parse(req))
        artifact, version = walker.get_package_info()
        return artifact, version, [requirement for requirement in requirements if requirement is not None]
    except Exception:
        # if the setup file is broken, we can't do much about that...
        return '?', '?', []


def from_setup_cfg(text) -> Tuple[str, str, List[DetectedRequirement]]:
    requirements = []
    start = False
    artifact = '?'
    version = '?'
    for line in text.split('\n'):
        if line.strip().startswith('name'):
            line = line.strip()
            artifact = line[line.find('=') + 1:].strip()
        if line.strip().startswith('version'):
            line = line.strip()
            version = line[line.find('=') + 1:].strip()
        if line.strip().startswith('install_requires'):
            start = True
            continue
        if start:
            if not line.startswith('\t'):
                start = False
                continue
            detected = DetectedRequirement.parse(line.strip())
            if detected is None:
                continue
            requirements.append(detected)

    return artifact, version, requirements


def from_pyproject_toml(text) -> Tuple[str, str, List[DetectedRequirement]]:
    requirements = []

    parsed = toml.loads(text)
    poetry_section = parsed.get("tool", {}).get("poetry", {})
    artifact = poetry_section.get("name", '?')
    version = poetry_section.get("version", '?')
    dependencies = poetry_section.get("dependencies", {})
    # dependencies.update(poetry_section.get("dev-dependencies", {}))

    for name, spec in dependencies.items():
        if name.lower() == "python":
            continue
        if isinstance(spec, dict):
            if "version" in spec:
                spec = spec["version"]
            elif 'git' in spec:
                t = spec.get('git')
                if 'rev' in spec:
                    tv = 'rev:' + spec.get('rev')
                elif 'tag' in spec:
                    tv = 'tag:' + spec.get('tag')
                else:
                    tv = 'branch:' + spec.get('branch', 'master')
                req = DetectedRequirement()
                req.name = t
                req.version_specs = [('', tv)]
                requirements.append(req)
                continue
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
    return artifact, version, requirements
