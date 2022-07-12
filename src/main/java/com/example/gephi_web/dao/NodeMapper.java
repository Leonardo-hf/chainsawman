package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVNode;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class NodeMapper {
    @Autowired
    JdbcTemplate jdbcTemplate;

    public void insertNode(String tableName,List<CSVNode> nodes) {
        for (CSVNode node : nodes) {
            String sql="insert into `"+tableName+"`(`id`,``name`, `attributes`) values(?, ?, ?);";
            jdbcTemplate.update(sql, node.getId(),node.getName(),node.getAttributes());
        }
    }
}
