package com.example.gephi_web.vo;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.util.List;

@Getter
@Setter
@ToString
public class FilterVO {
    String graphName;

    List<String> nodeNameList;

    int distance = 1;
}
