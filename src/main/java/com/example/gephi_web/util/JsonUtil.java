package com.example.gephi_web.util;

import com.alibaba.fastjson.JSON;

import java.util.Map;

public class JsonUtil {

    public static String marshall(Map<String, Object> src) {
        return JSON.toJSONString(src);
    }

    public static Map<String, Object> unmarshall(String src) {
        return JSON.parseObject(src);
    }
}
