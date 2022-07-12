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
        try {
            BufferedReader nodeFile = new BufferedReader(new FileReader(file));
            String lineDta = "";
            while ((lineDta = nodeFile.readLine()) != null){
                if(lineDta.equals("")) continue;
                if(!lineDta.startsWith("序号")){
                    int index=lineDta.indexOf(",");
                    CSVNode csvNode=new CSVNode();
                    csvNode.setId(Integer.parseInt(lineDta.substring(0,index)));
                    if(lineDta.length()>index+1)
                        csvNode.setName(lineDta.substring(index+1));
                    // todo attributes
                    csvNode.setAttributes("");
                    nodeMapper.insertNode(tableName,csvNode);
                }
            }
            nodeFile.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void addEdge(String tableName, String filePath) {
        File file = new File(filePath);
        try {
            BufferedReader edgeFile = new BufferedReader(new FileReader(file));
            String lineDta = "";
            while ((lineDta = edgeFile.readLine()) != null){
                if(lineDta.equals("")) continue;
                if(!lineDta.startsWith("Source")){
                    int index=lineDta.indexOf(",");
                    CSVEdge csvEdge=new CSVEdge();
                    csvEdge.setSource(lineDta.substring(0,index));
                    csvEdge.setTarget(lineDta.substring(index+1));
                    // todo attributes
                    csvEdge.setAttributes("");
                    edgeMapper.insertEdge(tableName,csvEdge);
                }
            }
            edgeFile.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 根据查询条件计算出相应的图并返回图片的url
     * @param type 查询的节点的种类 java/python/cpp/github……
     * @param nodeNameList
     * @return
     */
    @Override
    public String searchNodes(String type, List<String> nodeNameList) {
        List<CSVNode> nodes=nodeMapper.search(type,nodeNameList);
        List<CSVEdge> edges=edgeMapper.search(type,nodeNameList);
        //todo 生成图片并返回url
        return "";
    }
}
