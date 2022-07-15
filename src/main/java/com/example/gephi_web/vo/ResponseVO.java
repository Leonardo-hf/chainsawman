package com.example.gephi_web.vo;


import lombok.Getter;
import lombok.NonNull;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@ToString
public class ResponseVO<T> {

    private Integer code;

    private String msg;

    private T data;

    public ResponseVO(HttpStatus httpStatus) {
        this.code = httpStatus.getCode();
        this.msg = httpStatus.getMessage();
    }

    public ResponseVO(HttpStatus httpStatus, @NonNull T vo) {
        this.code = httpStatus.getCode();
        this.msg = httpStatus.getMessage();
        this.data = vo;
    }

}
