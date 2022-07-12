package com.example.gephi_web.service.impl;

import com.example.gephi_web.dao.EdgeMapper;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import com.example.gephi_web.service.GraphService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

@Service
public class GraphServiceImpl implements GraphService {

    @Autowired
    EdgeMapper edgeMapper;

    @Autowired
    NodeMapper nodeMapper;


    /**
     * 从表中读取数据存入数据库
     * @param tableName
     * @param filePath
     */
    @Override
    public void addNode(String tableName, String filePath) {
        File file = new File(filePath);
        List<CSVNode> nodes=new ArrayList<>();
        try {
            BufferedReader nodeFile = new BufferedReader(new FileReader(file));
            String lineDta = "";
            while ((lineDta = nodeFile.readLine()) != null){
                if(!lineDta.startsWith("序号")){
                    String[] tmp=lineDta.split(",");
                    CSVNode csvNode=new CSVNode();
                    csvNode.setId(Integer.parseInt(tmp[0]));
                    csvNode.setName(tmp[1]);
                    nodes.add(csvNode);
                }
            }
            nodeFile.close();
            nodeMapper.insertNode(tableName,nodes);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void addEdge(String tableName, String filePath) {
        File file = new File(filePath);
        List<CSVEdge> edges=new ArrayList<>();
        try {
            BufferedReader edgeFile = new BufferedReader(new FileReader(file));
            String lineDta = "";
            while ((lineDta = edgeFile.readLine()) != null){
                if(!lineDta.startsWith("Source")){
                    String[] tmp=lineDta.split(",");
                    CSVEdge csvEdge=new CSVEdge();
                    csvEdge.setSource(tmp[0]);
                    csvEdge.setTarget(tmp[1]);
                    edges.add(csvEdge);
                }
            }
            edgeFile.close();
            edgeMapper.insertEdge(tableName,edges);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
