package com.example.gephi_web.config;


import com.example.gephi_web.util.Const;
import org.springframework.context.annotation.Configuration;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import java.io.File;

@Component
@Configuration
public class FileConfig implements WebMvcConfigurer {

    /**
     * 将开放路径映射到准备的文件夹
     * @param registry
     */
    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        registry.addResourceHandler("/**").addResourceLocations("file:" + Const.Resource);
    }
}
