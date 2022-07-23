package com.example.gephi_web.service.impl;

import com.alibaba.fastjson.JSON;
import com.example.gephi_web.dao.EdgeMapper;
import com.example.gephi_web.dao.GraphMapper;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.util.Const;
import com.example.gephi_web.util.FileUtil;
import com.example.gephi_web.util.GexfUtil;
import com.example.gephi_web.util.GraphUtil;
import com.example.gephi_web.vo.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;

@Service
public class GraphServiceImpl implements GraphService {

    @Autowired
    EdgeMapper edgeMapper;

    @Autowired
    NodeMapper nodeMapper;

    @Autowired
    GraphMapper graphMapper;

    /**
     * 从表中读取数据存入数据库
     *
     * @param tableName
     * @param file
     */
    @Override
    public void addNode(String tableName, File file) {
        try {
            BufferedReader nodeFile = new BufferedReader(new FileReader(file));
            String lineDta;
            List<String> attrNames = new ArrayList<>();
            List<CSVNode> nodes = new ArrayList<>();
            boolean header = true;
            while ((lineDta = nodeFile.readLine()) != null) {
                if (header) {
                    String[] columns = lineDta.split(",");
                    if (columns.length >= 2 && columns[0].equals("id")) {
                        attrNames.addAll(Arrays.asList(columns).subList(2, columns.length));
                        // 创建表格
                        nodeMapper.createTable(tableName);
                    } else {
                        // TODO: 异常处理
                    }
                    header = false;
                    continue;
                }
                if (!lineDta.isEmpty() && !lineDta.isBlank()) {
                    String[] columns = lineDta.split(",");
                    CSVNode csvNode = new CSVNode();
                    csvNode.setId(Integer.parseInt(columns[0]));
                    if (columns.length < 2) {
                        continue;
                    }
                    csvNode.setName(columns[1]);
                    Map<String, Object> attrMap = new HashMap<>();
                    for (int i = 2, p = 0; i < columns.length; i++, p++) {
                        attrMap.put(attrNames.get(p), columns[i]);
                    }
                    csvNode.setAttributes(JSON.toJSONString(attrMap));
                    nodes.add(csvNode);
                }
            }
            nodeMapper.batchInsert(tableName, nodes);
            nodeFile.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void addEdge(String tableName, File file) {
        try {
            BufferedReader edgeFile = new BufferedReader(new FileReader(file));
            String lineDta;
            List<CSVEdge> edges = new ArrayList<>();
            List<String> attrNames = new ArrayList<>();
            boolean header = true;
            while ((lineDta = edgeFile.readLine()) != null) {
                if (header) {
                    String[] columns = lineDta.split(",");
                    if (columns.length >= 2 && columns[1].equals("source") && columns[2].equals("target")) {
                        attrNames.addAll(Arrays.asList(columns).subList(2, columns.length));
                        // 创建表格
                        edgeMapper.createTable(tableName);
                    } else {
                        //TODO: 异常处理
                    }
                    header = false;
                    continue;
                }
                if (!lineDta.isEmpty() && !lineDta.isBlank()) {
                    String[] columns = lineDta.split(",");
                    CSVEdge csvEdge = new CSVEdge();
                    csvEdge.setSource(Integer.parseInt(columns[1]));
                    csvEdge.setTarget(Integer.parseInt(columns[2]));
                    Map<String, Object> attrMap = new HashMap<>();
                    for (int i = 3, p = 0; i < columns.length; i++, p++) {
                        attrMap.put(attrNames.get(p), columns[i]);
                    }
                    csvEdge.setAttributes(JSON.toJSONString(attrMap));
                    edges.add(csvEdge);
                }
            }
            edgeMapper.batchInsert(tableName, edges);
            edgeFile.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 根据查询条件计算出相应的图并返回图片
     */
    @Override
    public ResponseVO<GexfVO> searchNodes(FilterVO filterVO) {
        String graphName = filterVO.getGraphName();
        List<CSVNode> nodes = nodeMapper.search(graphName, filterVO.getNodeNameList());
        Set<Integer> nodeIdSet = nodes.stream().map(CSVNode::getId).collect(Collectors.toSet()),
                visitIdSet = new HashSet<>();
        Set<CSVEdge> allEdge = new HashSet<>();
        int distance = filterVO.getDistance();
        while (distance > 0) {
            distance--;
            Set<Integer> toFindSet = new HashSet<>(nodeIdSet);
            toFindSet.removeAll(visitIdSet);
            List<CSVEdge> edges = edgeMapper.search(graphName, new ArrayList<>(toFindSet));
            allEdge.addAll(edges);
            visitIdSet.addAll(nodeIdSet);
            edges.forEach(edge -> {
                nodeIdSet.add(edge.getSource());
                nodeIdSet.add(edge.getTarget());
            });
        }
        List<CSVNode> allNode = nodeMapper.searchById(graphName, new ArrayList<>(nodeIdSet));
        return buildTempGraph(allNode, new ArrayList<>(allEdge));
    }

    // TODO: TEST
    @Override
    public ResponseVO<GexfVO> upload(UploadVO uploadVO) {
        String graphName = uploadVO.getGraphName();
        if (graphMapper.contains(graphName)) {
            return new ResponseVO<>(HttpStatus.FILE_ALREADY_EXISTS);
        }
        File node = FileUtil.save(uploadVO.getNodeFile());
        addNode(graphName, node);
        File edge = FileUtil.save(uploadVO.getEdgeFile());
        addEdge(graphName, edge);
        List<CSVNode> nodes = nodeMapper.search(graphName);
        List<CSVEdge> edges = edgeMapper.search(graphName);
        graphMapper.insert(graphName);
        return buildGraph(graphName, nodes, edges);
    }

//    public void addEdgeWithName(String tableName, File edge) {
//        try {
//            BufferedReader edgeFile = new BufferedReader(new FileReader(edge));
//            String line;
//            List<String> attrNames = new ArrayList<>();
//            boolean header = true;
//            while ((line = edgeFile.readLine()) != null) {
//                if (header) {
//                    String[] columns = line.split(",");
//                    if (columns.length >= 2 && columns[0].equals("Source") && columns[1].equals("Target")) {
//                        attrNames.addAll(Arrays.asList(columns).subList(2, columns.length));
//                        // 创建表格
//                        edgeMapper.createTable(tableName);
//                    } else {
//                        System.err.println("table format error");
//                        return;
//                    }
//                    header = false;
//                    continue;
//                }
//                if (!line.isEmpty() && !line.isBlank()) {
//                    String[] columns = line.split(",");
//                    CSVEdge csvEdge = new CSVEdge();
//                    CSVNode node = nodeMapper.search(tableName, columns[0]);
//                    if (node == null) {
//                        continue;
//                    }
//                    csvEdge.setSource(node.getId());
//                    node = nodeMapper.search(tableName, columns[1]);
//                    if (node == null) {
//                        continue;
//                    }
//                    csvEdge.setTarget(node.getId());
//                    Map<String, Object> attrMap = new HashMap<>();
//                    for (int i = 2; i < columns.length; i++) {
//                        attrMap.put(attrNames.get(i), columns[i]);
//                    }
//                    csvEdge.setAttributes(JSON.toJSONString(attrMap));
//                    edgeMapper.insert(tableName, csvEdge);
//                }
//            }
//            edgeFile.close();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }
//    }

    public ResponseVO<GexfVO> buildTempGraph(List<CSVNode> nodes, List<CSVEdge> edges) {
        return buildGraph(UUID.randomUUID().toString(), nodes, edges);
    }

    /**
     * 根据`节点`和`边`及其他属性构建图
     *
     * @param graphName 查询的节点的种类 java/python/cpp/github……
     * @param nodes     节点
     * @param edges     边
     * @return 标准返回
     * @todoparam strategy 策略相关
     */
    // TODO: TEST
    public ResponseVO<GexfVO> buildGraph(String graphName, List<CSVNode> nodes, List<CSVEdge> edges) {
        List<GexfUtil.GexfNode> gexfNodes = nodes.stream().map(n -> GexfUtil.GexfNode.getInstance(n.getId(), n.getName(), JSON.parseObject(n.getAttributes()))).collect(Collectors.toList());
        List<GexfUtil.GexfEdge> gexfEdges = edges.stream().map(e -> GexfUtil.GexfEdge.getInstance(e.getSource(), e.getTarget(), JSON.parseObject(e.getAttributes()))).collect(Collectors.toList());
        String rawPath = Const.Resource + FileUtil.getMD5Name(graphName, Const.RAW) + Const.GexfEXT;
        GexfUtil.createGexf(gexfNodes, gexfEdges, rawPath);
        String modPath = Const.Resource + FileUtil.getMD5Name(graphName, Const.MODIFY) + Const.JsonEXT;
        GraphUtil.getGraph(rawPath, modPath);
        return new ResponseVO<>(HttpStatus.COMMON_OK, new GexfVO(graphName, modPath));
    }

    @Override
    public ResponseVO<List<GexfVO>> look() {
        List<String> graphs = graphMapper.looksAll();
        List<GexfVO> gexfs = new ArrayList<>();
        for (String graph : graphs) {
            gexfs.add(new GexfVO(graph, Const.Resource + FileUtil.getMD5Name(graph, Const.MODIFY) + Const.JsonEXT));
        }
        return new ResponseVO<>(HttpStatus.COMMON_OK, gexfs);
    }
}
