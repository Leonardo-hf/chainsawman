package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVEdge;
import org.checkerframework.checker.units.qual.A;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class EdgeMapper {
    @Autowired
    JdbcTemplate jdbcTemplate;

    public void insertEdge(String tableName, List<CSVEdge> edges) {
        for (CSVEdge edge : edges) {
            String sql="insert into `"+tableName+"`(`source`, `target`, `attributes`) values(?, ?, ?);";
            jdbcTemplate.update(sql, edge.getSource(), edge.getTarget(), edge.getAttributes());
        }
    }
}
