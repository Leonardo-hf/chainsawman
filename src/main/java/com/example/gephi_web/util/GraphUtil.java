package com.example.gephi_web.util;

import org.gephi.appearance.api.*;
import org.gephi.appearance.plugin.PartitionElementColorTransformer;
import org.gephi.appearance.plugin.RankingNodeSizeTransformer;
import org.gephi.appearance.plugin.palette.Palette;
import org.gephi.appearance.plugin.palette.PaletteManager;
import org.gephi.filters.api.FilterController;
import org.gephi.filters.api.Query;
import org.gephi.filters.api.Range;
import org.gephi.filters.plugin.graph.DegreeRangeBuilder;
import org.gephi.graph.api.*;
import org.gephi.io.exporter.api.ExportController;
import org.gephi.io.importer.api.Container;
import org.gephi.io.importer.api.ImportController;
import org.gephi.io.processor.plugin.DefaultProcessor;
import org.gephi.layout.plugin.force.StepDisplacement;
import org.gephi.layout.plugin.force.yifanHu.YifanHuLayout;
import org.gephi.preview.api.PreviewController;
import org.gephi.preview.api.PreviewModel;
import org.gephi.preview.api.PreviewProperty;
import org.gephi.project.api.ProjectController;
import org.gephi.project.api.Workspace;
import org.gephi.statistics.plugin.Modularity;
import org.openide.util.Lookup;
import uk.ac.ox.oii.sigmaexporter.SigmaExporter;
import uk.ac.ox.oii.sigmaexporter.model.ConfigFile;

import java.io.*;

public class GraphUtil {
    public static void getGraph(String srcPath, String destPath) {
        // 准备环境
        // 初始化一个项目
        ProjectController pc = Lookup.getDefault().lookup(ProjectController.class);
        pc.newProject();
        Workspace workspace = pc.getCurrentWorkspace();
        // 获取控制器和模型
        ImportController importController = Lookup.getDefault().lookup(ImportController.class);
        GraphModel graphModel = Lookup.getDefault().lookup(GraphController.class).getGraphModel();
        AppearanceController appearanceController = Lookup.getDefault().lookup(AppearanceController.class);
        AppearanceModel appearanceModel = appearanceController.getModel();

        // 获得原始数据
        // 导入文件
        Container container = null;
        try {
            File file = new File(srcPath);
            container = importController.importFile(file);
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        // 过滤
        // 将导入数据附加到 Graph API，并移除 degree <= 1 的孤立点
        importController.process(container, new DefaultProcessor(), workspace);
        FilterController filterController = Lookup.getDefault().lookup(FilterController.class);
        DegreeRangeBuilder.DegreeRangeFilter degreeFilter = new DegreeRangeBuilder.DegreeRangeFilter();
        degreeFilter.init(graphModel.getGraph());
        degreeFilter.setRange(new Range(1, Integer.MAX_VALUE));
        Query query = filterController.createQuery(degreeFilter);
        GraphView view = filterController.filter(query);
        graphModel.setVisibleView(view);
        // 获得图
        UndirectedGraph filterNewGraph = graphModel.getUndirectedGraphVisible();
        int pointCount = filterNewGraph.getNodeCount(); // 过滤后点数量
        int edgeCount = filterNewGraph.getEdgeCount();  // 过滤后线数量
        System.out.println("Nodes: " + pointCount);
        System.out.println("Edges: " + edgeCount);

        // 模块化
        // 运行社区模块化算法
        Modularity modularity = new Modularity();
        modularity.setUseWeight(true); // 是否使用权重：使用边的权重
        modularity.setRandom(true);  // 是否随机：产生更好的分解，但是会增加计算时间
        modularity.setResolution(1.0); // 解析度:1.0是标准的解析度，数字越小，社区越多
        modularity.execute(graphModel);
        // 由模块化算法创建分类并设置颜色
        Column modColumn = graphModel.getNodeTable().getColumn(Modularity.MODULARITY_CLASS);

        Function func = appearanceModel.getNodeFunction(modColumn, PartitionElementColorTransformer.class);
        Partition partition = ((PartitionFunction) func).getPartition();
        int divisionCount = partition.size(filterNewGraph); // 社区划分数量
        System.out.println(divisionCount + " partitions found");
        Palette palette = PaletteManager.getInstance().randomPalette(divisionCount);
        partition.setColors(filterNewGraph, palette.getColors());
        appearanceController.transform(func);
        // 分配节点大小：按中心性排序大小
        Column centralityColumn = graphModel.getNodeTable().getColumn(Modularity.MODULARITY_CLASS);
        // TODO?没有传入图真得没事吗
        Function centralityRanking = appearanceModel.getNodeFunction(centralityColumn, RankingNodeSizeTransformer.class);
        RankingNodeSizeTransformer centralityTransformer = centralityRanking.getTransformer();
        centralityTransformer.setMinSize(1);
        centralityTransformer.setMaxSize(5);
        appearanceController.transform(centralityRanking);


        // 布局
        // YifanHuLayout布局算法
        YifanHuLayout layout = new YifanHuLayout(null, new StepDisplacement(1f));
        layout.setGraphModel(graphModel);
        layout.resetPropertiesValues();
        layout.setOptimalDistance(100.0f);  // 设置最佳距离(弹簧自然长度，更大的值意味着节点将相距较远)
        layout.setStep(10.0f);  // 设置初始步长（在整合阶段的初始步长。将此值设置为一个有意义的大小，相比于最佳距离，10%是一个很好的起点）
        layout.setStepRatio(0.95f); // 设置步比率（该比率用于更新各次迭代的步长）
        layout.setAdaptiveCooling(true); // 设置自适应冷却（控制自适应冷却的使用。它是用来帮助布局算法以避免能量局部极小）
        layout.initAlgo();
        for (int i = 0; i < 10000 && layout.canAlgo(); i++) {
            layout.goAlgo();
        }

        // 展示并导出文件
        // 设置“显示标签”选项，并禁用节点大小对文本大小的影响
        PreviewModel previewModel = Lookup.getDefault().lookup(PreviewController.class).getModel();
        previewModel.getProperties().putValue(PreviewProperty.SHOW_NODE_LABELS, Boolean.FALSE);  // 是否展示node_label
        previewModel.getProperties().putValue(PreviewProperty.NODE_LABEL_PROPORTIONAL_SIZE, Boolean.FALSE); // 禁用节点大小对文本大小的影响
        previewModel.getProperties().putValue(PreviewProperty.EDGE_CURVED, Boolean.TRUE);  // 是否弯曲
        // 导出gexf（仅导出可见部分
        ExportController ec = Lookup.getDefault().lookup(ExportController.class);
//        GraphExporter exporter = (GraphExporter) ec.getExporter("gexf");
//        exporter.setExportVisible(true);
//        exporter.setWorkspace(workspace);
        try {
            ec.exportFile(new File(FileUtil.getRoot()+"test.pdf"));
//            ec.exportFile(FileUtil.getFile(destPath), exporter);
        } catch (IOException ex) {
            ex.printStackTrace();
        }
        SigmaExporter se = new SigmaExporter();
        se.setWorkspace(workspace);
        ConfigFile cf = new ConfigFile();
        cf.setDefaults();
        se.setConfigFile(cf, FileUtil.getRoot()+Const.TempDir, false);
        se.execute();
        try (InputStream dataset = new FileInputStream(FileUtil.getRoot()+Const.TempDataSet); InputStream config = new FileInputStream(FileUtil.getRoot()+Const.TempConfig); OutputStream sigmaJSON = new FileOutputStream(destPath)) {
            byte[] buf = new byte[1024];
            int bytesRead;
            while ((bytesRead = dataset.read(buf)) > 0) {
                sigmaJSON.write(buf, 0, bytesRead);
            }
            sigmaJSON.write(System.lineSeparator().getBytes());
            while ((bytesRead = config.read(buf)) > 0) {
                sigmaJSON.write(buf, 0, bytesRead);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

    }
}
