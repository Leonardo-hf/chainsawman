package com.example.gephi_web.service;

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

    // 根据查询条件计算出相应的图并返回图片的url
    ResponseVO<GexfVO> searchNodes(String type, List<String> nodeNameList);

    ResponseVO<GexfVO> upload(UploadVO uploadVO);

    ResponseVO<List<GexfVO>> look();
}
