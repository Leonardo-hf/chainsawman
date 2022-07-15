package com.example.gephi_web.Service;

import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.service.GraphService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.IOException;
import java.util.List;

@SpringBootTest
public class ServiceTest {
    @Autowired
    GraphService graphService;

    @Autowired
    NodeMapper nodeMapper;

    @Test
    public void testBuildGraphWithFliter() throws IOException {
        String grahName = "Java";
        List<String> nodeNameList = nodeMapper.dealClosure(grahName,"/Users/taozehua/Downloads/研究任务/图谱构建/closure.csv","/Users/taozehua/Downloads/研究任务/图谱构建/requirements.csv");

        System.out.println(graphService.searchNodes(grahName, nodeNameList));
    }
}
