package com.example.gephi_web;

import com.example.gephi_web.dao.EdgeMapper;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.service.impl.GraphServiceImpl;
import com.example.gephi_web.util.Const;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.File;
import java.util.List;

@SpringBootTest
class GephiWebApplicationTests {

    @Autowired
    GraphService graphService;

    @Test
    void contextLoads() {
    }

    @Test
    public void testInsertData() {
        File fileNode = new File(Const.Resource + "node.csv");
        File fileEdge = new File(Const.Resource + "edge.csv");
//        graphService.addNode("test", fileNode);
        graphService.addEdge("test", fileEdge);
    }


    @Autowired
    NodeMapper nodeMapper;

    @Autowired
    EdgeMapper edgeMapper;

    @Test
    public void testBuildGraph() {
        GraphServiceImpl graphService = new GraphServiceImpl();
        List<CSVNode> nodes = nodeMapper.search("test");
        List<CSVEdge> edges = edgeMapper.search("test");
        graphService.buildGraph("test", nodes, edges);
    }

}
