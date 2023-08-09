import re

from lxml import etree

from util.http import spider

central_url = 'https://repo1.maven.org/maven2'
base_url = central_url

TAG_GROUP = "groupId"
TAG_ARTIFACT = "artifactId"
TAG_VERSION = "version"
TAG_PARENT = "parent"
TAG_DEPENDENCIES_MANAGE = "dependencyManagement"
TAG_DEPENDENCIES = "dependencies"
TAG_DEPENDENCY = "dependency"
TAG_OPTIONAL = "optional"
TAG_SCOPE = "scope"
TAG_PROPERTIES = "properties"

UNDEFINED = "$undefined"


def is_undefined(v):
    return v == UNDEFINED


def is_none(v):
    return v is None


def is_valid(v):
    return not is_undefined(v) and not is_none(v)


class POM:

    def __init__(self):
        self._root = UNDEFINED
        self._namespace = UNDEFINED
        self._artifact = UNDEFINED
        self._group = UNDEFINED
        self._version = UNDEFINED
        self._dependencies = UNDEFINED
        self._dependencies_management = UNDEFINED
        self._properties = UNDEFINED
        self._parent = UNDEFINED
        self._plain = UNDEFINED
        self._url = UNDEFINED

    def iter(self, node, tag):
        return node.iter(self._namespace + tag)

    def search(self, node, tag):
        v = node.find(self._namespace + tag)
        if is_none(v):
            return None
        return v

    def value(self, node, tag):
        v = node.find(self._namespace + tag)
        if is_none(v):
            return None
        return v.text

    def value_or_default(self, node, tag, default):
        v = self.value(node, tag)
        if is_none(v):
            return default
        return v

    def get_root(self):
        if not is_undefined(self._root):
            return self._root
        if not is_undefined(self._plain):
            self._root = etree.fromstring(self._plain.encode('UTF-8'), parser=etree.XMLParser(remove_comments=True))
            self._namespace = ''
            match = re.search('{.*}', self._root.tag)
            if match:
                span = match.span()
                self._namespace = self._root.tag[span[0]: span[1]]
            return self._root
        if is_valid(self._url):
            try:
                self._plain = spider(self._url).text
            except:
                raise Exception('[POM] invalid coordinate/url, url={}'.format(self._url))
            return self.get_root()
        if is_valid(self._group) and is_valid(self._artifact) and is_valid(self._version):
            self._url = '{}/{}/{}/{}/{}-{}.pom'.format(
                base_url, self._group.replace('.', '/'), self._artifact, self._version,
                self._artifact, self._version)
            return self.get_root()
        raise Exception('[POM] require plain/url/coordinate to parse pom at least')

    def get_group_id(self):
        if not is_undefined(self._group):
            return self._group
        root = self.get_root()
        v = self.value(root, TAG_GROUP)
        if is_none(v):
            parent = self.get_parent()
            if is_none(parent):
                raise Exception('[POM] cannot find {}'.format(TAG_PARENT))
            v = parent.get_group_id()
        self._group = v
        return v

    def get_artifact(self):
        if not is_undefined(self._artifact):
            return self._artifact
        root = self.get_root()
        v = self.value(root, TAG_ARTIFACT)
        if is_none(v):
            raise Exception('[POM] cannot find {}'.format(TAG_ARTIFACT))
        self._artifact = v
        return v

    def get_version(self):
        if not is_undefined(self._version):
            return self._version
        root = self.get_root()
        v = self.value(root, TAG_VERSION)
        if is_none(v) or v == '${parent.version}':
            parent = self.get_parent()
            if is_none(parent):
                raise Exception('[POM] cannot find {}'.format(TAG_PARENT))
            v = parent.get_version()
        self._version = v
        return v

    def get_dependencies(self):
        if not is_undefined(self._dependencies):
            return self._dependencies
        root = self.get_root()
        deps = self.search(root, TAG_DEPENDENCIES)
        if is_none(deps):
            self._dependencies = None
            return None
        self._dependencies = self._extract_dependencies(deps)
        return self._dependencies

    def get_dependencies_management(self):
        if not is_undefined(self._dependencies_management):
            return self._dependencies_management
        root = self.get_root()
        deps_manage = self.search(root, TAG_DEPENDENCIES_MANAGE)
        if is_none(deps_manage):
            self._dependencies_management = None
            return None
        deps = self.search(deps_manage, TAG_DEPENDENCIES)
        self._dependencies_management = self._extract_dependencies(deps, is_manage=True)
        return self._dependencies_management

    def _extract_dependencies(self, deps, is_manage=False):
        dependencies = []
        for dep in self.iter(deps, TAG_DEPENDENCY):
            # get group
            group = self.value(dep, TAG_GROUP)
            if group == '${project.groupId}':
                group = self.get_group_id()
            if is_none(group):
                raise Exception('[POM] dependency has no {}'.format(TAG_GROUP))
            # get artifact
            artifact = self.value(dep, TAG_ARTIFACT)
            if is_none(artifact):
                raise Exception('[POM] dependency has no {}'.format(TAG_ARTIFACT))
            # get version
            version = self.value(dep, TAG_VERSION)
            if version == '${project.version}':
                version = self.get_version()
            if not is_none(version):
                match = re.search('\\${.*}', version)
                if match:
                    span = match.span()
                    key = version[span[0] + 2: span[1] - 1]
                    pom = self
                    while not is_none(pom):
                        properties = self.get_properties()
                        if key in properties:
                            version = properties[key]
                            break
                        pom = pom.get_parent()
            if is_none(version) and not is_manage:
                parent = self.get_parent()
                while not is_none(parent) and is_none(version):
                    pdeps = parent.get_dependencies_management()
                    parent = parent.get_parent()
                    if is_none(pdeps):
                        continue
                    for d in pdeps:
                        if group == d.group and artifact == d.artifact:
                            version = d.version
                            break
            if is_none(version):
                raise Exception('[POM] dependency has no {}'.format(TAG_VERSION))
            # get scope & optional
            scope = self.value_or_default(dep, TAG_SCOPE, default='provided')
            optional = self.value_or_default(dep, TAG_OPTIONAL, default='false')
            dependencies.append(
                Dependency(group=group, artifact=artifact, version=version, scope=scope, optional=optional))
        return dependencies

    def get_properties(self):
        if not is_undefined(self._properties):
            return self._properties
        root = self.get_root()
        v = self.search(root, TAG_PROPERTIES)
        self._properties = v
        if not is_none(v):
            self._properties = {}
            for p in v:
                self._properties[p.tag[len(self._namespace):]] = p.text
        return self._properties

    def get_parent(self):
        if not is_undefined(self._parent):
            return self._parent
        root = self.get_root()
        parent = self.search(root, TAG_PARENT)
        if is_none(parent):
            self._parent = None
            return self._parent
        group = self.value(parent, TAG_GROUP)
        artifact = self.value(parent, TAG_ARTIFACT)
        version = self.value(parent, TAG_VERSION)
        self._parent = self.from_coordinate(artifact=artifact, group=group, version=version)
        return self._parent

    @classmethod
    def from_coordinate(cls, artifact, group, version):
        pom = cls()
        pom._artifact = artifact
        pom._group = group
        pom._version = version
        return pom

    @classmethod
    def from_string(cls, plain):
        pom = cls()
        pom._plain = plain
        return pom

    @classmethod
    def from_url(cls, url):
        pom = cls()
        pom._url = url
        return pom


class Package:
    def __init__(self, group, artifact, version):
        self.artifact = artifact
        self.group = group
        self.version = version

    def __dict__(self):
        return {
            'artifact': self.artifact,
            'group': self.group,
            'version': self.version
        }


class Dependency(Package):
    def __init__(self, group, artifact, version, scope, optional):
        super().__init__(group, artifact, version)
        self.scope = scope
        self.optional = optional

    def __dict__(self):
        d = super().__dict__()
        d.update({
            'scope': self.scope,
            'optional': self.optional
        })
        return d
