package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVEdge;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.ParameterizedPreparedStatementSetter;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.sql.*;
import java.util.*;

@Repository
public class EdgeMapper {

    @Resource
    JdbcTemplate jdbcTemplate;

    public void batchInsert(String tableName, List<CSVEdge> edges) {
        String sql = "insert into edge" + tableName + "(source, target, attributes) values(?, ?, ?);";
        jdbcTemplate.batchUpdate(sql, edges, 4096, new ParameterizedPreparedStatementSetter<CSVEdge>() {
            public void setValues(PreparedStatement ps, CSVEdge edge)
                    throws SQLException {
                ps.setInt(1, edge.getSource());
                ps.setInt(2, edge.getTarget());
                ps.setString(3, edge.getAttributes());
            }
        });
    }

    public void insert(String tableName, CSVEdge edge) {
        String sql = "insert into edge" + tableName + "(source, target, attributes) values(?, ?, ?);";
        jdbcTemplate.update(sql, edge.getSource(), edge.getTarget(), edge.getAttributes());
    }

    /**
     * 根据用户输入的节点列表，查找节点对应的依赖的边
     *
     * @param type
     * @param nodeIdList
     * @return
     */
    public List<CSVEdge> search(String type, List<Integer> nodeIdList) {
        Set<CSVEdge> edges = new HashSet<>();
        for (Integer id : nodeIdList) {
            String sql1 = "select source,target,attributes from edge" + type + " where source= " + id + "";
            String sql2 = "select source,target,attributes from edge" + type + " where target= " + id + "";
            getEdgeIntoSet(edges, sql1);
            getEdgeIntoSet(edges, sql2);
        }
        return new ArrayList<>(edges);
    }

    public List<CSVEdge> search(String type) {
        Set<CSVEdge> edges = new HashSet<>();
        String sql = "select id, source, target, attributes from edge" + type;
        getEdgeIntoSet(edges, sql);
        return new ArrayList<>(edges);
    }

    private void getEdgeIntoSet(Set<CSVEdge> edges, String sql) {
        List<Map<String, Object>> queryList2 = jdbcTemplate.queryForList(sql);
        if (queryList2 != null && !queryList2.isEmpty()) {
            for (Map<String, Object> map : queryList2) {
                CSVEdge edge = new CSVEdge();
                edge.setId((Integer) map.get("id"));
                edge.setSource((Integer) map.get("source"));
                edge.setTarget((Integer) map.get("target"));
                edge.setAttributes((String) map.get("attributes"));
                edges.add(edge);
            }
        }
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
        jdbcTemplate.update(String.format("""
                DROP TABLE IF EXISTS `edge%s`;
                create table `edge%s`
                (
                    id         int auto_increment
                        primary key,
                    source     int          not null,
                    target     int          not null,
                    attributes varchar(255) null
                ) ENGINE = InnoDB
                  DEFAULT CHARSET = utf8mb4
                  COLLATE = utf8mb4_0900_ai_ci;

                create index source__index
                    on edge%s (source);

                create index target__index
                    on edge%s (target);
                                """, graphName, graphName, graphName, graphName));
    }
}
