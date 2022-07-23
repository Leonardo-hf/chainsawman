package com.example.gephi_web.service;

import com.example.gephi_web.vo.FilterVO;
import com.example.gephi_web.vo.GexfVO;
import com.example.gephi_web.vo.ResponseVO;
import com.example.gephi_web.vo.UploadVO;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.List;

@Service
public interface GraphService {

    void addNode(String tableName, File file);

    void addEdge(String tableName, File file);

    // 根据查询条件计算出相应的图
    ResponseVO<GexfVO> searchNodes(FilterVO filterVO);

    ResponseVO<GexfVO> upload(UploadVO uploadVO);

    ResponseVO<List<GexfVO>> look();
}
