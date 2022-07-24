package com.example.gephi_web.Service;

import com.alibaba.fastjson.JSON;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVNode;
import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.vo.FilterVO;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

@SpringBootTest
public class ServiceTest {
    @Autowired
    GraphService graphService;

    @Autowired
    NodeMapper nodeMapper;

    @Test
    public void testBuildGraphWithFliter() throws IOException {
        List<String> nodeNameList = new ArrayList<>();
        try {
            File file=new File("/Users/taozehua/Downloads/研究任务/可视化系统/gephi_toolkits_service/static/node.csv");
            BufferedReader nodeFile = new BufferedReader(new FileReader(file));
            String lineDta;
            while ((lineDta = nodeFile.readLine()) != null) {
                if (!lineDta.isEmpty() && !lineDta.startsWith("id")) {
                    String[] columns = lineDta.split(",");
                    if (columns.length < 2) {
                        continue;
                    }
                    nodeNameList.add(columns[1]);

                }
            }
            nodeFile.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
        FilterVO filterVO=new FilterVO();
        String grahName = "test";
        filterVO.setGraphName(grahName);
        filterVO.setNodeNameList(nodeNameList);
        System.out.println(graphService.searchNodes(filterVO));
    }
}
