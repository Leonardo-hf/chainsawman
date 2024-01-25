import json
from typing import List, Tuple, Optional, Dict

import toml

from common import HttpStatus, RustLang
from util import spider, Singleton
from vo import Dep, ModuleDeps
from .index import DepsHandler


@Singleton
class RustDepsHandler(RustLang, DepsHandler):

    @staticmethod
    def rust_v_spec(v) -> List[Tuple[str, str]]:
        tvs = v.split(',')
        res: List[Tuple[str, str]] = []
        for tv_specs in tvs:
            tv_specs = tv_specs.strip()
            if tv_specs.startswith('^') or tv_specs.startswith('~'):
                tv = tv_specs[1:]
                limit = tv_specs[0]
            elif ' ' in tv_specs:
                tv_specs = tv_specs.split(' ')
                tv = tv_specs[1]
                limit = tv_specs[0]
            else:
                tv = tv_specs
                limit = ''
            res.append((tv, limit))
        return res

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        if module.endswith(self.MODULE_CARGO):
            parsed = toml.loads(data.decode())
            info = parsed.get('package', {})
            artifact = info.get('name', '?')
            version = info.get('version', '?')

            def parse_deps(deps: Dict[str, Dict], scope: str = '') -> List[Dep]:
                res: List[Dep] = []
                for t, detail in deps.items():
                    tv = 'latest'
                    t_optional = False
                    if isinstance(detail, str):
                        tv = detail
                    if isinstance(detail, dict):
                        if 'path' in detail:
                            continue
                        t = detail.get('package', t)
                        t_optional = detail.get('optional', t_optional)
                        if 'git' in detail:
                            t = detail.get('git')
                            if 'rev' in detail:
                                tv = 'rev@' + detail.get('rev')
                            elif 'tag' in detail:
                                tv = 'tag@' + detail.get('tag')
                            else:
                                tv = 'latest@' + detail.get('branch', 'master')
                        else:
                            tv = detail.get('version', tv)
                    tvs = self.rust_v_spec(tv)
                    for tv_specs in tvs:
                        res.append(
                            Dep(artifact=t, version=tv_specs[0], limit=tv_specs[1], optional=t_optional, scope=scope))
                return res

            deps = parse_deps(parsed.get('dependencies', {}))
            dev_deps = parse_deps(parsed.get('dev-dependencies', {}), 'dev')
            deps.extend(dev_deps)
            return ModuleDeps(lang=self.lang(), path=module, artifact=artifact, version=version,
                              dependencies=deps), HttpStatus.OK
        return None, HttpStatus.NOT_SUPPORT

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        package = package[package.find(' ') + 1:]
        s = package.rfind(':')
        artifact = package[:s].strip()
        version = package[s + 1:]
        if len(artifact) == 0:
            return None, HttpStatus.NOT_FOUND
        elif 1 <= len(artifact) <= 2:
            url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/{}/{}'.format(len(artifact), artifact)
        elif len(artifact) == 3:
            url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/3/{}/{}'.format(artifact[0], artifact)
        else:
            url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/{}/{}/{}'.format(artifact[0:2], artifact[2:4],
                                                                                         artifact)
        module_deps: List[Dep] = []
        content = spider(url).text
        content = content.strip().split('\n')
        target = None
        try:
            if version == 'latest':
                target = json.loads(content[-1].strip())
            else:
                for j in content:
                    j = json.loads(j.strip())
                    if j['vers'] == version:
                        target = j
                        break
        except Exception:
            pass
        if target is None:
            return None, HttpStatus.NOT_FOUND

        def trans_scope(s: str) -> str:
            if s == 'normal':
                return ''
            else:
                return s

        def get_name(d) -> str:
            if 'package' in d:
                return d['package']
            return d['name']

        deps = list(map(lambda d: (get_name(d), self.rust_v_spec(d['req']), trans_scope(d['kind'])), target['deps']))
        for name, v_spec, scope in deps:
            for v, limit in v_spec:
                module_deps.append(Dep(artifact=name, version=v, limit=limit, scope=scope))
        return ModuleDeps(lang=self.lang(), artifact=artifact, version=version, dependencies=module_deps), HttpStatus.OK
