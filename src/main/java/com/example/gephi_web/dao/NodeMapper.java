package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVNode;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.sql.Connection;
import java.sql.DatabaseMetaData;
import java.sql.ResultSet;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Repository
public class NodeMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public void insertNode(String tableName, CSVNode node) {
        String sql = "insert into " + tableName + "(id,name,attributes) values(?, ?, ?);";
        jdbcTemplate.update(sql, node.getId(), node.getName(), node.getAttributes());
    }

    /**
     * 根据用户输入的节点名称列表，查找对应的完整的CSVNode
     *
     * @param type
     * @param nodeNameList
     * @return
     */
    public List<CSVNode> search(String type, List<String> nodeNameList) {
        List<CSVNode> nodeList = new ArrayList<>();
        for (String nodeName : nodeNameList) {
            String sql = "select id, name, attributes from node" + type + " where name= \"" + nodeName+"\"";
            CSVNode csvNode = jdbcTemplate.queryForObject(sql, (rs, rowNum) -> {
                CSVNode node = new CSVNode();
                assert rs != null;
                node.setId(rs.getInt(1));
                node.setName(rs.getString(2));
                node.setAttributes(rs.getString(3));
                return node;
            });
            nodeList.add(csvNode);
        }
        return nodeList;
    }

    public List<CSVNode> search(String type) {
        List<CSVNode> nodeList = new ArrayList<>();
        String sql = "select id, name, attributes from node" + type;
        List<Map<String, Object>> queryList = jdbcTemplate.queryForList(sql);
        if (queryList != null && !queryList.isEmpty()) {
            for (Map<String, Object> map : queryList) {
                CSVNode node = new CSVNode();
                node.setId((Integer) map.get("id"));
                node.setName((String) map.get("name"));
                node.setAttributes((String) map.get("attributes"));
            }
        }
        return nodeList;
    }

    public void createTable(String graphName) {
        try (Connection conn = jdbcTemplate.getDataSource().getConnection()) {
            DatabaseMetaData dbMetaData = conn.getMetaData();
            String[] types = {"TABLE"};
            ResultSet tabs = dbMetaData.getTables(null, null, graphName, types);
            if (tabs.next()) {
                return;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        jdbcTemplate.update("""
                CREATE TABLE `node%s`
                (
                    `id`         int NOT NULL,
                    `name`       varchar(255)                                                  DEFAULT NULL,
                    `attributes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                    PRIMARY KEY (`id`)
                ) ENGINE = InnoDB
                  DEFAULT CHARSET = utf8mb4
                  COLLATE = utf8mb4_0900_ai_ci;
                """, graphName);
    }

}
