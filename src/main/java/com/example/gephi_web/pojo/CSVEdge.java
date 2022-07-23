package com.example.gephi_web.pojo;

import lombok.*;

@Data
@Getter
@Setter
@ToString
@EqualsAndHashCode
public class CSVEdge {
    Integer id;
    Integer source;
    Integer target;
    String attributes;
}
