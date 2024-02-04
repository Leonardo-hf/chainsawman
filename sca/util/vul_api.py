import json
from collections import defaultdict
from functools import reduce
from typing import List, Tuple, Mapping

import requests
from packageurl import PackageURL

from vo import OSV


class VulAPI:
    def __init__(self, base_url: str):
        self._base_url = base_url

    def query_vul_by_purl(self, purl: PackageURL) -> List[OSV]:
        return self.query_vul_by_purls([purl])[0]

    def query_vul_by_purls(self, purls: List[PackageURL]) -> List[List[OSV]]:
        def merge_from_purl(m, e: Tuple[int, PackageURL]) -> Mapping:
            purl = e[1]
            v = ''
            if purl.version:
                v = purl.version
            m.update({
                e[0]: {
                    'purl': PackageURL(type=purl.type, namespace=purl.namespace, name=purl.name).to_string(),
                    'version': v
                }
            })
            return m

        def parse_res(ret: List[Mapping]) -> List[OSV]: return list(
            map(lambda v: OSV(id=v['id'], aliases=v.get('aliases', ''), summary=v['summary'], details=v['details'],
                              cwe=v.get('database_specific', defaultdict(str))['cwe_ids'],
                              severity=v.get('database_specific', defaultdict(str))['severity'],
                              ref=v.get('references', [defaultdict(str)])[0]['url'])
                , map(lambda v: v['database_specific']['osv'][0], ret)))

        api_name = 'searchByAffected'
        url = '%s%s' % (self._base_url, ','.join(list(map(lambda _: api_name, range(len(purls))))))
        json_str = json.dumps(reduce(lambda a, b: merge_from_purl(a, b), enumerate(purls), {}))
        r = requests.get(url, params={'batch': 1, 'input': json_str})
        return list(map(lambda x: parse_res(x['result']['data']), json.loads(r.text)))
