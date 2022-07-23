package com.example.gephi_web.util;

import org.springframework.util.DigestUtils;
import org.springframework.util.FileCopyUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.FileOutputStream;
import java.util.Properties;

public class FileUtil {

    static Properties properties = System.getProperties();

    public static String getMD5Name(String... args) {
        StringBuilder origin = new StringBuilder();
        for (String arg : args) {
            origin.append(arg);
        }
        return DigestUtils.md5DigestAsHex(origin.toString().getBytes());
    }

    public static String getRoot(){
        return properties.getProperty("user.dir")+File.separator;
    }

    public static File save(MultipartFile file, String path) {
        if (file == null) {
            return null;
        }
        File save = new File(path);
        try {
            FileOutputStream fileOutputStream = new FileOutputStream(save);
            FileCopyUtils.copy(file.getInputStream(), fileOutputStream);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return save;
    }

    public static File save(MultipartFile file) {
        return save(file, file.getOriginalFilename());
    }
}
