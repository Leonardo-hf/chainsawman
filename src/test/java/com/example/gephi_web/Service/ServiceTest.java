package com.example.gephi_web.Service;

import com.example.gephi_web.service.GraphService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@SpringBootTest
public class ServiceTest {
    @Autowired
    GraphService graphService;

    @Test
    public void testBuildGraphWithFliter() throws IOException {
        String grahName = "Java";
        List<String> nodeNameList = new ArrayList<>();
        File file = new File("/Users/taozehua/Downloads/研究任务/joup/node.csv");
        BufferedReader edgeFile = new BufferedReader(new FileReader(file));
        String lineDta;
        for (int i = 0; i < 5; i++) {
            lineDta = edgeFile.readLine();
            if (lineDta != null && lineDta.contains("/")) {
                String pkg = lineDta.split(",")[1];
                nodeNameList.add(pkg);
            }
        }
        System.out.println(graphService.searchNodes(grahName, nodeNameList));
    }
}
