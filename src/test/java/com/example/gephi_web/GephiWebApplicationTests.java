package com.example.gephi_web;

import com.example.gephi_web.service.GraphService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.transaction.annotation.Transactional;

@SpringBootTest
class GephiWebApplicationTests {

    @Autowired
    GraphService graphService;

    @Test
    void contextLoads() {
    }

    @Test
    public void testInsertData(){
        graphService.addNode("nodeJava","/Users/taozehua/Downloads/研究任务/joup/node.csv");
//        graphService.addEdge("edgeJava","/Users/taozehua/Downloads/研究任务/joup/edge.csv");
    }
}
