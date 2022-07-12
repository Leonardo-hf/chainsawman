package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVEdge;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.sql.Connection;
import java.sql.DatabaseMetaData;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.*;

@Repository
public class EdgeMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public void insertEdge(String tableName, CSVEdge edge) {
        String sql = "insert into " + tableName + "(source, target, attributes) values(?, ?, ?);";
        jdbcTemplate.update(sql, edge.getSource(), edge.getTarget(), edge.getAttributes());
    }

    /**
     * 根据用户输入的节点列表，查找节点对应的依赖的边
     *
     * @param type
     * @param nodeNameList
     * @return
     */
    public List<CSVEdge> search(String type, List<String> nodeNameList) {
        Set<Integer> nodeIDList = new HashSet<>();
        for (String nodeName : nodeNameList) {
            String sql = "select id from node" + type + " where name=" + nodeName;
            Map<String, Object> map = jdbcTemplate.queryForMap(sql);
            assert map != null;
            int id = (int) map.get("id");
            nodeIDList.add(id);
        }
        List<CSVEdge> edges = new ArrayList<>();
        for (Integer id : nodeIDList) {
            String sql1 = "select source,target,attributes from edge" + type + " where source=" + id;
            String sql2 = "select source,target,attributes from edge" + type + " where target=" + id;
            CSVEdge csvEdge = jdbcTemplate.queryForObject(sql1, (rs, rowNum) -> {
                CSVEdge csvEdge1 = new CSVEdge();
                assert rs != null;
                csvEdge1.setId(rs.getInt(1));
                csvEdge1.setSource(rs.getString(2));
                csvEdge1.setTarget(rs.getString(3));
                csvEdge1.setAttributes(rs.getString(4));
                return csvEdge1;
            });
            edges.add(csvEdge);
            csvEdge = jdbcTemplate.queryForObject(sql2, new RowMapper<CSVEdge>() {
                @Override
                public CSVEdge mapRow(ResultSet rs, int rowNum) throws SQLException {
                    CSVEdge csvEdge1 = new CSVEdge();
                    csvEdge1.setId(rs.getInt(1));
                    csvEdge1.setSource(rs.getString(2));
                    csvEdge1.setTarget(rs.getString(3));
                    csvEdge1.setAttributes(rs.getString(4));
                    return csvEdge1;
                }
            });
            edges.add(csvEdge);
        }
        return edges;
    }

    public List<CSVEdge> search(String type) {
        List<CSVEdge> nodeList = new ArrayList<>();
        String sql = "select id, source, target, attributes from edge" + type;
        List<Map<String, Object>> queryList = jdbcTemplate.queryForList(sql);
        if (queryList != null && !queryList.isEmpty()) {
            for (Map<String, Object> map : queryList) {
                CSVEdge edge = new CSVEdge();
                edge.setId((Integer) map.get("id"));
                edge.setSource((String) map.get("source"));
                edge.setTarget((String) map.get("target"));
                edge.setAttributes((String) map.get("attributes"));
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
                CREATE TABLE `edge%s`
                (
                `id`         int          NOT NULL AUTO_INCREMENT,
                `source`     varchar(255) NOT NULL,
                `target`     varchar(255) NOT NULL,
                `attributes` varchar(255) DEFAULT NULL,
                PRIMARY KEY (`id`)
                ) ENGINE = InnoDB
                DEFAULT CHARSET = utf8mb4
                COLLATE = utf8mb4_0900_ai_ci;
                """, graphName);
    }
}
