package com.example.gephi_web.vo;

import lombok.Getter;

@Getter
public enum HttpStatus {

    COMMON_OK(4000, "ok"),
    FILE_ALREADY_EXISTS(6000, "不应有重复的表格名"),
    FILE_NOT_EXIST(6001, "文件不存在");

    HttpStatus(int code, String message) {
        this.code = code;
        this.message = message;
    }

    private final int code;

    private final String message;
}

