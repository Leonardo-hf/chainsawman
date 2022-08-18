package com.example.gephi_web.vo;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.util.List;

@Getter
@Setter
@ToString
@AllArgsConstructor
public class FilterVO {
    String graphName;

    List<String> nodeNameList;

    public FilterVO(){}

    int distance = 1;

    public FilterVO(String graphName, List<String> nodeNameList) {
        this.graphName=graphName;
        this.nodeNameList=nodeNameList;
    }
}
