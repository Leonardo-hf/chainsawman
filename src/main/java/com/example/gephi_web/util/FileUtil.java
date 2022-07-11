package com.example.gephi_web.util;

import org.springframework.core.io.ClassPathResource;

import java.io.File;
import java.io.IOException;

public class FileUtil {

    public static File getFile(String path) {
        //TODO: 错误的，文件不存在时路径不正确
        File file = null;
        try {
            ClassPathResource resource = new ClassPathResource(path);
            if (resource.exists()) {
                return resource.getFile();
            } else {
                assert resource.getPath() != null;
                file = new File(resource.getPath());
                return file;
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        return null;
    }

    public static File saveFile(String path) {
        File file = null;
        try {
            file = new ClassPathResource(path).getFile();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return file;
    }

}
