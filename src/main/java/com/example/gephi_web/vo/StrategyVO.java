package com.example.gephi_web.vo;

import lombok.Getter;
import lombok.Setter;

import java.util.Map;

@Getter
@Setter
public class StrategyVO {

    String name;

    Map<String, Object> params;
}
