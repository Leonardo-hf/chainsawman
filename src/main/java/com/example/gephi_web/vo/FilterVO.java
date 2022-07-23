package com.example.gephi_web.vo;

import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
public class FilterVO {
    String graphName;

    List<String> nodeNameList;

    int distance = 1;
}
