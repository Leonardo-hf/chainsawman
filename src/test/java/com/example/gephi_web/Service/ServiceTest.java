package com.example.gephi_web.Service;

import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.vo.FilterVO;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.List;

@SpringBootTest
public class ServiceTest {
    @Autowired
    GraphService graphService;

    @Autowired
    NodeMapper nodeMapper;

    @Test
    public void testBuildGraphWithFliter() {
        String grahName = "test";
        List<String> nodeNameList = List.of("pytest", "requests", "six", "numpy", "flake8", "pandas", "matplotlib", "scipy", "sphinx", "nose", "toml", "lxml");
        FilterVO filterVO = new FilterVO();
        filterVO.setGraphName(grahName);
        filterVO.setNodeNameList(nodeNameList);
        filterVO.setDistance(2);
        System.out.println(graphService.searchNodes(filterVO));
    }
}
