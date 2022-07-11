package com.example.gephi_web.pojo;

import lombok.Data;

@Data
public class CSVEdge {
    Integer id;
    String source;
    String target;
    String attributes;
}
