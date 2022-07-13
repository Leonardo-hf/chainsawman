package com.example.gephi_web.service.impl;

import com.example.gephi_web.dao.EdgeMapper;
import com.example.gephi_web.dao.GraphMapper;
import com.example.gephi_web.dao.NodeMapper;
import com.example.gephi_web.pojo.CSVEdge;
import com.example.gephi_web.pojo.CSVNode;
import com.example.gephi_web.service.GraphService;
import com.example.gephi_web.service.impl.strategy.DefaultStrategy;
import com.example.gephi_web.service.impl.strategy.GraphStrategy;
import com.example.gephi_web.util.FileUtil;
import com.example.gephi_web.util.GexfUtil;
import com.example.gephi_web.util.JsonUtil;
import com.example.gephi_web.vo.GexfVO;
import com.example.gephi_web.vo.HttpStatus;
import com.example.gephi_web.vo.ResponseVO;
import com.example.gephi_web.vo.UploadVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.*;
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
            boolean header = true;
            while ((lineDta = nodeFile.readLine()) != null) {
                if (header) {
                    String[] columns = lineDta.split(",");
                    if (columns.length >= 2 && columns[0].equals("序号")) {
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
                    // TODO：还有没名字的吗？
                    if(columns.length>=2)
                        csvNode.setName(columns[1]);
                    Map<String, Object> attrMap = new HashMap<>();
                    for (int i = 2; i < columns.length; i++) {
                        attrMap.put(attrNames.get(i), columns[i]);
                    }
                    csvNode.setAttributes(JsonUtil.marshall(attrMap));
                    nodeMapper.insertNode(tableName, csvNode);
                }
            }
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
            List<String> attrNames = new ArrayList<>();
            boolean header = true;
            while ((lineDta = edgeFile.readLine()) != null) {
                if (header) {
                    String[] columns = lineDta.split(",");
                    if (columns.length >= 2 && columns[0].equals("Source") && columns[1].equals("Target")) {
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
                    csvEdge.setSource(columns[0]);
                    csvEdge.setTarget(columns[1]);
                    Map<String, Object> attrMap = new HashMap<>();
                    for (int i = 2; i < columns.length; i++) {
                        attrMap.put(attrNames.get(i), columns[i]);
                    }
                    csvEdge.setAttributes(JsonUtil.marshall(attrMap));
                    edgeMapper.insertEdge(tableName, csvEdge);
                }
            }
            edgeFile.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 根据查询条件计算出相应的图并返回图片
     *
     * @param graphName    查询的节点的种类 java/python/cpp/github……
     * @param nodeNameList
     * @return
     */
    @Override
    public ResponseVO<GexfVO> searchNodes(String graphName, List<String> nodeNameList) {
        List<CSVNode> nodes = nodeMapper.search(graphName, nodeNameList);
        List<CSVEdge> edges = edgeMapper.search(graphName, nodeNameList);
        return buildGraph(graphName, nodes, edges);
    }

    // TODO: TEST
    @Override
    public ResponseVO<GexfVO> upload(UploadVO uploadVO) {
        String graphName = uploadVO.getGraphName();
        File node = FileUtil.save(uploadVO.getNodeFile());
        addNode(graphName, node);
        File edge = FileUtil.save(uploadVO.getEdgeFile());
        addEdge(graphName, edge);
        List<CSVNode> nodes = nodeMapper.search(graphName);
        List<CSVEdge> edges = edgeMapper.search(graphName);
        return buildGraph(graphName, nodes, edges);
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
    private ResponseVO<GexfVO> buildGraph(String graphName, List<CSVNode> nodes, List<CSVEdge> edges) {
        List<GexfUtil.GexfNode> gexfNodes = nodes.stream().map(n -> GexfUtil.GexfNode.getInstance(n.getName(), JsonUtil.unmarshall(n.getAttributes()))).collect(Collectors.toList());
        List<GexfUtil.GexfEdge> gexfEdges = edges.stream().map(e -> GexfUtil.GexfEdge.getInstance(e.getSource(), e.getTarget(), JsonUtil.unmarshall(e.getAttributes()))).collect(Collectors.toList());
        GexfUtil.createGexf(gexfNodes, gexfEdges, FileUtil.getPath(graphName));
        // TODO: 编写一种组织图像名称的方法，替换掉现在的File.getPath方法，比如：filename=md5(graph=?&strategy=?&filter=?)
        // TODO: 这里暂时写死为DefaultStrategy，可能之后通过工厂模式构造 graphStrategy
        String destPath = FileUtil.getPath(graphName, "test");
        GraphStrategy strategy = new DefaultStrategy();
        strategy.getGraph(graphName, destPath);
        return new ResponseVO<>(HttpStatus.COMMON_OK, new GexfVO(graphName, destPath));
    }

    @Override
    public ResponseVO<List<GexfVO>> look() {
        List<String> graphs = graphMapper.looksAll();
        List<GexfVO> gexfs = new ArrayList<>();
        for (String graph : graphs) {
            // TODO: 同上，替换FileUtil.getPath
            gexfs.add(new GexfVO(graph, FileUtil.getPath(graph, "test")));
        }
        return new ResponseVO<>(HttpStatus.COMMON_OK, gexfs);
    }
}
