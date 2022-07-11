package com.example.gephi_web.util;

import it.uniroma1.dis.wsngroup.gexf4j.core.*;
import it.uniroma1.dis.wsngroup.gexf4j.core.data.*;
import it.uniroma1.dis.wsngroup.gexf4j.core.impl.GexfImpl;

import it.uniroma1.dis.wsngroup.gexf4j.core.impl.StaxGraphWriter;
import it.uniroma1.dis.wsngroup.gexf4j.core.impl.data.AttributeListImpl;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GexfUtil {

    public static class GexfNode {
        String name;
        Map<String, Object> attributes;

        public static GexfNode getInstance(String name, Map<String, Object> attributes) {
            GexfNode gexfNode = new GexfNode();
            gexfNode.name = name;
            gexfNode.attributes = attributes;
            return gexfNode;
        }
    }

    public static class GexfEdge {
        String source;
        String target;
        Map<String, Object> attributes;

        public static GexfEdge getInstance(String source, String target, Map<String, Object> attributes) {
            GexfEdge gexfEdge = new GexfEdge();
            gexfEdge.source = source;
            gexfEdge.target = target;
            gexfEdge.attributes = attributes;
            return gexfEdge;
        }
    }

    public static void createGexf(List<GexfNode> nodes, List<GexfEdge> edges, String path) {
        if (nodes.size() == 0) {
            return;
        }
        // 定义gexf文件
        Gexf gexf = new GexfImpl();
        Graph graph = gexf.getGraph();
        AttributeList attrList = new AttributeListImpl(AttributeClass.NODE);
        graph.getAttributeLists().add(attrList);
        // 设置点属性列表
        List<Attribute> attributes = new ArrayList<>();
        int cnt = 0;
        for (String attrName : nodes.get(0).attributes.keySet()) {
            attributes.add(attrList.createAttribute(String.valueOf(cnt++), AttributeType.STRING, attrName));
        }
        // 创建点
        Map<String, Node> nodeMap = new HashMap<>();
        Node node;
        int pointLength = nodes.size();
        for (int i = 0; i < pointLength; i++) {
            GexfNode gexfNode = nodes.get(i);
            node = graph.createNode(String.valueOf(i))
                    .setLabel(gexfNode.name);
            AttributeValueList attributeValueList = node.getAttributeValues();
            for (Attribute attr : attributes) {
                attributeValueList.addValue(attr, String.valueOf(gexfNode.attributes.get(attr.getTitle())));
            }
            nodeMap.put(gexfNode.name, node);
        }
        if (edges.size() != 0) {
            // 设置边属性列表
            attributes.clear();
            cnt = 0;
            for (String attrName : edges.get(0).attributes.keySet()) {
                attributes.add(attrList.createAttribute(String.valueOf(cnt++), AttributeType.STRING, attrName));
            }
            // 创建边
            for (GexfEdge gexfEdge : edges) {
                try {
                    String sourceNodeKey = gexfEdge.source;
                    String targetNodeKey = gexfEdge.target;
                    Node sourceNode = nodeMap.get(sourceNodeKey);
                    Node targetNode = nodeMap.get(targetNodeKey);
                    Edge edge = sourceNode.connectTo(targetNode);
                    AttributeValueList attributeValueList = edge.getAttributeValues();
                    for (Attribute attr : attributes) {
                        attributeValueList.addValue(attr, String.valueOf(gexfEdge.attributes.get(attr.getTitle())));
                    }
                } catch (Exception ignored) {
                }
            }
        }
        // 生成gexf文件
        StaxGraphWriter graphWriter = new StaxGraphWriter();
        File f = FileUtil.getFile(path);
        Writer out;
        try {
            out = new FileWriter(f, false);
            graphWriter.writeToStream(gexf, out, "UTF-8");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}
