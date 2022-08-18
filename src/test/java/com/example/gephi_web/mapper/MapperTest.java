package com.example.gephi_web.mapper;

import com.example.gephi_web.dao.EdgeMapper;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.ArrayList;
import java.util.List;

@SpringBootTest
public class MapperTest {

    @Autowired
    NodeMapper nodeMapper;

    @Autowired
    EdgeMapper edgeMapper;

    @Test
    public void testSearchNodeByNodeNameList() {
        String type = "Java";
        List<String> nodeNameList = new ArrayList<>();
        nodeNameList.add("abbot/abbot/1.4.0");
        nodeNameList.add("junit/junit/4.8.2");
        nodeNameList.add("androidx/compose/runtime/runtime/1.0.0-alpha09");
        List<CSVNode> nodeList = nodeMapper.search(type, nodeNameList);
        for (CSVNode node : nodeList) {
            System.out.println(node);
        }
        assert nodeList.get(0).getId() == 3;
        assert nodeList.get(1).getId() == 4;
        assert nodeList.get(2).getId() == 14;
    }

    @Test
    public void testSearchEdgeByNodeNameList() {
    }
}