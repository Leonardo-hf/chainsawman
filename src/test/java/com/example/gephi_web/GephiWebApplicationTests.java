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

import java.io.*;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;

@SpringBootTest
class GephiWebApplicationTests {

    @Autowired
    GraphService graphService;

    @Test
    public void testInsertData() {
        File fileNode = new File(Const.Resource + "node.csv");
        File fileEdge = new File(Const.Resource + "edge.csv");
        graphService.addNode("test", fileNode);
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

    @Test
    public void dealDRequirements() throws IOException {
        File d_requirements = new File("static/d_requirements.csv");
        BufferedReader requirmentFile = new BufferedReader(new FileReader(d_requirements));
        String lineDta;
        HashSet<String> packageSet = new HashSet<>();
        while ((lineDta = requirmentFile.readLine()) != null) {
//            System.out.println(lineDta);
            if (!lineDta.isEmpty() && !lineDta.startsWith("package,requirement")) {
                String[] columns = lineDta.split(",");
                if (columns.length < 2) {
                    packageSet.add(columns[0]);
                } else {
                    packageSet.add(columns[0]);
                    packageSet.add(columns[1]);
                }
            }
        }
        requirmentFile.close();
        Map<String, Integer> packageMap = new HashMap<>();
        int nodeId = 1;
        File node = new File("static/node.csv");
        BufferedWriter writeNode = new BufferedWriter(new FileWriter(node));
        writeNode.write("id,name,attributes");
        for (String pkg : packageSet) {
            packageMap.put(pkg, nodeId);
            writeNode.newLine();
            writeNode.write(nodeId + "," + pkg + ",");
            nodeId++;
        }
        writeNode.flush();
        writeNode.close();
        File egde = new File("static/edge.csv");
        BufferedWriter writeEdge = new BufferedWriter(new FileWriter(egde));
        BufferedReader requirmentFile2 = new BufferedReader(new FileReader(d_requirements));
        String lineDta2;
        int id = 1;
        writeEdge.write("id,source,target,attributes");
        while ((lineDta2 = requirmentFile2.readLine()) != null) {
            if (!lineDta2.isEmpty() && !lineDta2.startsWith("package,requirement")) {
                String[] columns = lineDta2.split(",");
                if (columns.length < 2) continue;
                writeEdge.newLine();
                writeEdge.write(id + "," + packageMap.get(columns[0]) + "," + packageMap.get(columns[1]) + ",");
                id++;
            }
        }
        writeEdge.flush();
        writeEdge.close();
        requirmentFile2.close();
    }

}
