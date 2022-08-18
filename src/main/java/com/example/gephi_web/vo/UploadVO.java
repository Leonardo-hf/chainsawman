package com.example.gephi_web.vo;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.ToString;
import org.springframework.web.multipart.MultipartFile;

@Getter
@ToString
@AllArgsConstructor
public class UploadVO {

    public UploadVO(){}

    String graphName;

    MultipartFile nodeFile;

    MultipartFile edgeFile;
}
