package com.example.gephi_web.vo;

import lombok.Getter;

@Getter
public enum HttpStatus {

    COMMON_OK(4000, "ok");

    HttpStatus(int code, String message) {
        this.code = code;
        this.message = message;
    }

    private final int code;

    private final String message;
}

