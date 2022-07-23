package com.example.gephi_web.dao;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Repository
public class GraphMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public List<String> looksAll() {
        String sql = "select name from graph";
        List<Map<String, Object>> queryList = jdbcTemplate.queryForList(sql);
        List<String> graphs = new ArrayList<>();
        if (queryList != null && !queryList.isEmpty()) {
            for (Map<String, Object> map : queryList) {
                graphs.add((String) map.get("name"));
            }
        }
        return graphs;
    }

    public boolean contains(String graphName) {
        String sql = String.format("select id from graph where name = '%s'", graphName);
        List<Map<String, Object>> queryList = jdbcTemplate.queryForList(sql);
        return !(queryList == null || queryList.isEmpty());
    }

    public void insert(String name) {
        String sql = String.format("insert into graph(name) values('%s')", name);
        jdbcTemplate.update(sql);
    }
}
