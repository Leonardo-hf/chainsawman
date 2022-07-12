package com.example.gephi_web.service;

import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface GraphService {

    void addNode(String tableName, String filePath);

    void addEdge(String tableName,String filePath);

    // 根据查询条件计算出相应的图并返回图片的url
    String searchNodes(String type,List<String> nodeNameList);
}
