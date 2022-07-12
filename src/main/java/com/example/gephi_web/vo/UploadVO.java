package com.example.gephi_web.vo;

import lombok.Getter;
import org.springframework.web.multipart.MultipartFile;

@Getter
public class UploadVO {

    String graphName;

    MultipartFile nodeFile;

    MultipartFile edgeFile;
}
