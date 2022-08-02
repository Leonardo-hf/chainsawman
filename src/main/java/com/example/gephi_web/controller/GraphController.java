package com.example.gephi_web.controller;

import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.vo.FilterVO;
import com.example.gephi_web.vo.GexfVO;
import com.example.gephi_web.vo.ResponseVO;
import com.example.gephi_web.vo.UploadVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.util.List;

@RestController
public class GraphController {

    @Autowired
    GraphService graphService;

    /**
     * 上传新的图表
     * @return
     */
    @PostMapping("/upload")
    public ResponseVO<GexfVO> upload(String graphName, MultipartFile nodeFile,MultipartFile edgeFile) {
        UploadVO uploadVO=new UploadVO(graphName,nodeFile,edgeFile);
        System.out.println(uploadVO);
        return graphService.upload(uploadVO);
    }

    /**
     * 查找图表的目标节点
     * @return
     */
    @PostMapping("/filter")
    public ResponseVO<GexfVO> filter(FilterVO filterVO) {
        System.out.println(filterVO);
        return graphService.searchNodes(filterVO);
    }

    /**
     * 查询现有图表
     * @return
     */
    @GetMapping("/look")
    public ResponseVO<List<GexfVO>> look() {
        return graphService.look();
    }

}
