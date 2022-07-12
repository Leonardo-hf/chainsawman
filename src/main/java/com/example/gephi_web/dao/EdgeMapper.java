package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVEdge;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.*;

@Repository
public class EdgeMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public void insertEdge(String tableName, List<CSVEdge> edges) {
        for (CSVEdge edge : edges) {
            String sql="insert into `"+tableName+"`(`source`, `target`, `attributes`) values(?, ?, ?);";
            jdbcTemplate.update(sql, edge.getSource(), edge.getTarget(), edge.getAttributes());
        }
    }

    /**
     * 根据用户输入的节点列表，查找节点对应的依赖的边
     * @param type
     * @param nodeNameList
     * @return
     */
    public List<CSVEdge> search(String type, List<String> nodeNameList){
        Set<Integer> nodeIDList=new HashSet<>();
        for(String nodeName:nodeNameList){
            String sql="select id from node"+type+" where name="+nodeName;
            Map<String,Object> map=jdbcTemplate.queryForMap(sql);
            int id= (int) map.get("id");
            nodeIDList.add(id);
        }
        List<CSVEdge> edges=new ArrayList<>();
        for(Integer id:nodeIDList){
            String sql1="select source,target,attributes from edge"+type+" where source="+id;
            String sql2="select source,target,attributes from edge"+type+" where target="+id;
            CSVEdge csvEdge=jdbcTemplate.queryForObject(sql1, new RowMapper<CSVEdge>() {
                @Override
                public CSVEdge mapRow(ResultSet rs, int rowNum) throws SQLException {
                    CSVEdge csvEdge1=new CSVEdge();
                    csvEdge1.setId(rs.getInt(1));
                    csvEdge1.setSource(rs.getString(2));
                    csvEdge1.setTarget(rs.getString(3));
                    csvEdge1.setAttributes(rs.getString(4));
                    return csvEdge1;
                }
            });
            edges.add(csvEdge);
            csvEdge=jdbcTemplate.queryForObject(sql2, new RowMapper<CSVEdge>() {
                @Override
                public CSVEdge mapRow(ResultSet rs, int rowNum) throws SQLException {
                    CSVEdge csvEdge1=new CSVEdge();
                    csvEdge1.setId(rs.getInt(1));
                    csvEdge1.setSource(rs.getString(2));
                    csvEdge1.setTarget(rs.getString(3));
                    csvEdge1.setAttributes(rs.getString(4));
                    return csvEdge1;
                }
            });
            edges.add(csvEdge);
        }
        return  edges;
    }
}
