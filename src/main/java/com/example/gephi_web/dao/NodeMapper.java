package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import org.jfree.data.io.CSV;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.*;

@Repository
public class NodeMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public void insertNode(String tableName, CSVNode node) {
            String sql="insert into "+tableName+"(id,name,attributes) values(?, ?, ?);";
            jdbcTemplate.update(sql, node.getId(),node.getName(),node.getAttributes());
    }

    /**
     * 根据用户输入的节点名称列表，查找对应的完整的CSVNode
     * @param type
     * @param nodeNameList
     * @return
     */
    public List<CSVNode> search(String type, List<String> nodeNameList) {
        List<CSVNode> nodeList = new ArrayList<>();
        for (String nodeName : nodeNameList) {
            String sql = "select id,name,attributes from node" + type + " where name=" + nodeName;
           CSVNode csvNode=jdbcTemplate.queryForObject(sql, new RowMapper<CSVNode>() {
                @Override
                public CSVNode mapRow(ResultSet rs, int rowNum) throws SQLException {
                    CSVNode csvNode1=new CSVNode();
                    csvNode1.setId(rs.getInt(1));
                    csvNode1.setName(rs.getString(2));
                    csvNode1.setAttributes(rs.getString(3));
                    return csvNode1;
                }
            });
            nodeList.add(csvNode);
        }
        return nodeList;
    }

}
