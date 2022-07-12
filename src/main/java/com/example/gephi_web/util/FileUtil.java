package com.example.gephi_web.util;

import org.springframework.util.DigestUtils;
import org.springframework.util.FileCopyUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.FileOutputStream;
import java.util.Properties;

public class FileUtil {

    static Properties properties = System.getProperties();

    public static String getPath(String... args) {
        StringBuilder origin = new StringBuilder();
        for (String arg : args) {
            origin.append(arg);
        }
        return DigestUtils.md5DigestAsHex(origin.toString().getBytes());
    }

    public static File getFile(String path) {
        return new File(properties.getProperty("user.dir") + File.separator + path);
    }

    public static File save(MultipartFile file, String name) {
        if (file == null) {
            return null;
        }
        File save = getFile(name);
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
