package com.example.gephi_web.dao;

import com.example.gephi_web.pojo.CSVNode;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.io.*;
import java.sql.Connection;
import java.sql.DatabaseMetaData;
import java.sql.ResultSet;
import java.util.*;

@Repository
public class NodeMapper {
    @Resource
    JdbcTemplate jdbcTemplate;

    public void insertNode(String tableName, CSVNode node) {
        String sql = "insert into " + tableName + "(id,name,attributes) values(?, ?, ?);";
        jdbcTemplate.update(sql, node.getId(), node.getName(), node.getAttributes());
    }


    public CSVNode search(String type, String nodeName) {
        List<String> nodeNameList = new ArrayList<>();
        nodeNameList.add(nodeName);
        List<CSVNode> res = search(type, nodeNameList);
        if (res.size() == 0) {
            return null;
        }
        return res.get(0);
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
            String sql = "select id, name, attributes from node" + type + " where name= \"" + nodeName + "\"";
            getNodeIntoList(nodeList, sql);
        }
        return nodeList;
    }

    public List<CSVNode> search(String type) {
        List<CSVNode> nodeList = new ArrayList<>();
        String sql = "select id, name, attributes from node" + type;
        getNodeIntoList(nodeList, sql);
        return nodeList;
    }

    private void getNodeIntoList(List<CSVNode> nodeList, String sql) {
        List<Map<String, Object>> queryList = jdbcTemplate.queryForList(sql);
        if (queryList != null && !queryList.isEmpty()) {
            for (Map<String, Object> map : queryList) {
                CSVNode node = new CSVNode();
                node.setId((Integer) map.get("id"));
                node.setName((String) map.get("name"));
                node.setAttributes((String) map.get("attributes"));
                nodeList.add(node);
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


    /**
     * 目前是两种想法：
     * 第一种在我们算出来的clousre中的一个包的所有版本的依赖在我们算出来的包对应的某一个版本中才画出来
     * 第二种是在不在都画出来
     */
    /**
     * 将closure进行处理，找到计算得到的闭包中的包对应的所有版本的序号
     * 以及这些版本依赖的包的序号（依赖的包也要在closure中），找到对应的版本名，返回出去
     *
     * @param type        处理的图片的类型 java/python/……
     * @param closure     计算得到的闭包文件地址
     * @param requirments requirements文件地址
     * @return nodeNameList
     * @throws IOException
     */
    public List<String> dealClosure(String type, String closure, String requirments) throws IOException {

        List<String> nodeNameList = new ArrayList<>();
        // 读取closure file
        File closureCSV = new File(closure);
        BufferedReader closureFile = new BufferedReader(new FileReader(closureCSV));
        Set<String> packageSet = new HashSet<>(); // 记录经过pagerank算法计算后的包的名称
        String lineDta = "";
        // 将文档的下一行数据赋值给lineData，并判断是否为空，若不为空则输出
        while ((lineDta = closureFile.readLine()) != null) {
            packageSet.add(lineDta);
        }
        closureFile.close();

        // 读取原来的requirement文件
        File requirmentCSV = new File(requirments);
        BufferedReader requirmentFile = new BufferedReader(new FileReader(requirmentCSV));

        String line = "";
        while ((line = requirmentFile.readLine()) != null) {
            if (line.equals("package,edition,requirement,constraint")) continue;
            // 0 包  1 版本对应的序列号  2 依赖的包  3 依赖包的版本的序列号
            String[] data = line.split(",");
            if (packageSet.contains(data[0])) {
                // 算法一和二的区别就在这，要不要依赖的包也在我们算出来的closure中
                if (packageSet.contains(data[2])) {
                    String sql1 = "select name from node" + type + " where id= " + data[1];
                    String sql2 = "select name from node" + type + " where id= " + data[3];
                    Map<String, Object> map1 = jdbcTemplate.queryForMap(sql1);
                    nodeNameList.add((String) map1.get("name"));
                    Map<String, Object> map2 = jdbcTemplate.queryForMap(sql2);
                    nodeNameList.add((String) map2.get("name"));
                }
            }
        }
        requirmentFile.close();
        return nodeNameList;
    }


}
