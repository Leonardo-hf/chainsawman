package com.example.gephi_web;

import com.example.gephi_web.service.GraphService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.File;

@SpringBootTest
class GephiWebApplicationTests {

    @Autowired
    GraphService graphService;

    @Test
    void contextLoads() {
    }

    @Test
    public void testInsertData(){
        File fileNode=new File("/Users/taozehua/Downloads/研究任务/joup/node.csv");
        File fileEdge=new File("/Users/taozehua/Downloads/研究任务/joup/edge.csv");
        graphService.addNode("nodeJava",fileNode);
        graphService.addEdge("edgeJava",fileEdge);
    }

}
