package com.example.gephi_web.service;

import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface GraphService {

    void addNode(String tableName, String filePath);

    void addEdge(String tableName,String filePath);
}
